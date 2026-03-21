package hdp

import (
	"cmp"
	"context"
	"log/slog"
	"maps"
	"slices"
	"strings"
	"text/template"
	"time"

	"github.com/it512/xxl-job-exec"
	"github.com/xen0n/go-workwx/v2"
)

type SummaryParam struct {
	StoreCode string `json:"store_code"`
	StoreName string `json:"store_name"`
	BotKey    string `json:"bot_key"`
}

type SummaryOutline struct {
	Yestoday time.Time

	Param SummaryParam

	Fee  FeeRecord
	Pay  PaymentRecord
	Sale SaleRecord
	Gm   GmRecord

	Gr GatherRate

	Holiday Holiday
}

type GatherRate struct {
	Gr map[string]*GatherItem
}

func (g GatherRate) Table() []*GatherItem {
	l := maps.Values(g.Gr)
	return slices.SortedFunc(l,
		func(a, b *GatherItem) int { return cmp.Compare(a.Ym, b.Ym) },
	)
}

func g(grs []GatherRecord) GatherRate {
	gm := make(map[string]*GatherItem)

	for _, gr := range grs {
		if gi, ok := gm[gr.Ym]; ok {
			if gr.IsPayment == 0 {
				gi.Fee0 = gr.Fee
			} else {
				gi.Fee1 = gr.Fee
			}
		} else {
			gi := &GatherItem{Ym: gr.Ym}
			if gr.IsPayment == 0 {
				gi.Fee0 = gr.Fee
			} else {
				gi.Fee1 = gr.Fee
			}
			gm[gr.Ym] = gi
		}
	}

	return GatherRate{Gr: gm}
}

type Summary struct {
	DBx    *DBx
	Tpl    *template.Template
	Logger *slog.Logger
}

func (b *Summary) Name() string {
	return "project-summary-per-day"
}

func (b *Summary) Run(ctx context.Context, task *xxl.Task) error {
	var param SummaryParam
	if err := xxl.TaskJsonParam(task, &param); err != nil {
		return err
	}

	outline, _ := b.DoRun(ctx, param)

	return b.Push(ctx, outline)
}

func (b *Summary) DoRun(ctx context.Context, param SummaryParam) (SummaryOutline, error) {

	yestoday := Yestoday(time.Now())
	dt := DT(yestoday)

	var outline SummaryOutline
	outline.Yestoday = yestoday
	outline.Fee, _ = b.DBx.FeeAgg(param.StoreCode, dt)
	outline.Pay, _ = b.DBx.PaymentAgg(param.StoreCode, dt)
	outline.Sale, _ = b.DBx.SaleAgg(param.StoreCode, dt)
	outline.Gm, _ = b.DBx.GmEntry(param.StoreCode, dt)
	gg, _ := b.DBx.GatherAgg(param.StoreCode)
	outline.Gr = g(gg)

	outline.Param = param

	return outline, nil
}

func (b *Summary) Push(ctx context.Context, outline SummaryOutline) error {
	var sb strings.Builder
	sb.Grow(2048)
	if err := b.Tpl.Execute(&sb, outline); err != nil {
		return err
	}

	wc := workwx.NewWebhookClient(outline.Param.BotKey)
	return wc.SendMarkdownV2Message(sb.String())
}
