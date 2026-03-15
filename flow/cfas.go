package flow

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/it512/xxl-job-exec"
	"github.com/twiglab/xjob/pfsdk"
	"github.com/twiglab/xjob/pfsdk/hik/cfas"
)

type Param struct {
	StoreCode string `json:"store_code"`
	StoreName string `json:"store_name"`
}

type HikParam struct {
	Param
	BaseURL   string `json:"url"`
	AppKey    string `json:"key"`
	AppSecret string `json:"secret"`
	GroupID   string `json:"ids"`
}

type HikJob struct {
	DBx *DBx
}

func (b *HikJob) Name() string {
	return "hik-cfas"
}

func (b *HikJob) Run(ctx context.Context, task *xxl.Task) error {
	var param HikParam
	if err := xxl.TaskJsonParam(task, &param); err != nil {
		return err
	}

	if param.StoreCode == "" {
		return errors.New("storecode is nil")
	}

	cfg := cfas.Config{
		BaseURL:   param.BaseURL,
		AppKey:    param.AppKey,
		AppSecret: param.AppSecret,
	}

	fc := cfas.New(cfg)
	yestoday := pfsdk.Yestoday(time.Now())
	start, end, _ := pfsdk.OpenTime(yestoday)
	in, _, _, err := b.Collect(ctx, start, end, fc, param.GroupID)
	if err != nil {
		return err
	}

	fmt.Println(in)
	return b.DBx.Save(ctx, yestoday, param.Param, in)
}

func (b *HikJob) Collect(ctx context.Context, start, end time.Time, c *cfas.Client, ids string) (in int, out int, keep int, err error) {
	pf := cfas.PassengerFlowIn{IDs: ids, Granularity: "minutely", StartTime: start, EndTime: end}
	var pfr cfas.PassengerFlowRtn
	pfr, err = c.PassengerFlow(ctx, pf)
	if err != nil {
		return
	}

	in, out, keep = cfas.Collect(pfr.Data.List)
	return
}
