package hik

import (
	"context"
	"strings"
	"text/template"
	"time"

	"github.com/it512/xxl-job-exec"
	"github.com/twiglab/xjob/pfsdk"
	"github.com/twiglab/xjob/pfsdk/hik/cfas"
	"github.com/xen0n/go-workwx/v2"
)

const last = 21 * 60 * 45

type Outline struct {
	Now       time.Time
	StoreName string
	StoreCode string

	GroupBy *GroupBy
}

func (o Outline) IsLast() bool {
	h, m, _ := o.Now.Clock()
	return (h*60 + m) > last
}

const SummaryTpl = `
# {{ .StoreName }}（{{ .StoreCode }}）营业期间客流 {{ .Now.Format "2006.01.02 15:04" }}
{{- $item := .GroupBy.Get "1" }}
{{- if not .IsLast}}
> 全场**{{ $item.In }}** (入)，**{{ $item.Out }}** (出)，场内人数 **{{ $item.Keep }}**人
{{- else}}
> 全场**{{ $item.In }}** (入)
{{- end }}
{{- $item := .GroupBy.Get "3" }}
> 长乐路方向 **{{ $item.In }}** 人，武定门方向 **{{ $item.Out }}** 人
{{- $item := .GroupBy.Get "4" }}
> 夫子庙方向 **{{ $item.In }}** 人，老门东方向 **{{ $item.Out }}** 人
`

type CfasPushBotPatam struct {
	CfasParam
	BotKey string `json:"bot_key"`
}

type CfasPushBot struct {
	Tpl *template.Template
}

func NewCfasPushBot() CfasPushBot {
	tpl, _ := template.New("summary").Parse(SummaryTpl)
	return CfasPushBot{
		Tpl: tpl,
	}
}

func (b CfasPushBot) Name() string {
	return "hik-cfas-push-bot"
}

func (b CfasPushBot) Run(ctx context.Context, task *xxl.Task) error {
	var param CfasPushBotPatam

	if err := xxl.TaskJsonParam(task, &param); err != nil {
		return err
	}
	now := time.Now()
	kt := pfsdk.MakeKeyTime(now)
	wc := workwx.NewWebhookClient(param.BotKey)

	cli := cfas.New(param.Config)

	pf := cfas.PassengerFlowIn{IDs: param.IDs, Granularity: cfas.MINUTELY, StartTime: kt.OpenStart, EndTime: now}
	pfr, err := cli.PassengerFlow(ctx, pf)
	if err != nil {
		return err
	}

	groupBy := GroupBySum(pfr.Data.List)

	var outline = Outline{
		GroupBy:   groupBy,
		Now:       now,
		StoreName: param.StoreName,
		StoreCode: param.StoreCode,
	}

	var sb strings.Builder
	sb.Grow(2048)
	if err := b.Tpl.Execute(&sb, outline); err != nil {
		return err
	}

	return wc.SendMarkdownV2Message(sb.String())
}
