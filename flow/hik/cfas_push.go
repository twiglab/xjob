package hik

import (
	"context"
	"fmt"
	"time"

	"github.com/it512/xxl-job-exec"
	"github.com/twiglab/xjob/pfsdk"
	"github.com/twiglab/xjob/pfsdk/hik/cfas"
	"github.com/xen0n/go-workwx/v2"
)

type CfasPushBot struct {
}

type CfasPushBotPatam struct {
	CfasParam
	BotKey string `json:"bot_key"`
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
	in, out, keep, err := Collect(ctx, kt.OpenStart, now, cli, param.IDs)
	if err != nil {
		return err
	}

	return wc.SendTextMessage(fmt.Sprintf("%s %s 进%d，出%d，在场%d",
		param.StoreName,
		pfsdk.DateTime(now),
		in, out, keep), nil)
}
