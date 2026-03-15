package hik

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/it512/xxl-job-exec"
	"github.com/twiglab/xjob/flow"
	"github.com/twiglab/xjob/pfsdk"
	"github.com/twiglab/xjob/pfsdk/hik/cfas"
)

type CfasParam struct {
	flow.Param
	cfas.Config
	IDs string `json:"ids"`
}

type HikJob struct {
	DBx *flow.DBx
}

func (b *HikJob) Name() string {
	return "hik-cfas"
}

func (b *HikJob) Run(ctx context.Context, task *xxl.Task) error {
	var param CfasParam
	if err := xxl.TaskJsonParam(task, &param); err != nil {
		return err
	}

	if param.StoreCode == "" {
		return errors.New("storecode is nil")
	}

	fc := cfas.New(param.Config)
	yestoday := pfsdk.Yestoday(time.Now())
	kt := pfsdk.MakeKeyTime(yestoday)
	in, _, _, err := Collect(ctx, kt.OpenStart, kt.OpenEnd, fc, param.IDs)
	if err != nil {
		return err
	}

	fmt.Println(in)
	return b.DBx.Save(ctx, yestoday, param.Param, in)
}

func Collect(ctx context.Context, start, end time.Time, c *cfas.Client, ids string) (in int, out int, keep int, err error) {
	pf := cfas.PassengerFlowIn{IDs: ids, Granularity: "minutely", StartTime: start, EndTime: end}
	var pfr cfas.PassengerFlowRtn
	pfr, err = c.PassengerFlow(ctx, pf)
	if err != nil {
		return
	}

	in, out, keep = cfas.Collect(pfr.Data.List)
	return
}
