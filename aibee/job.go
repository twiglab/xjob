package aibee

import (
	"context"
	"errors"
	"time"

	"github.com/imroc/req/v3"
	"github.com/it512/xxl-job-exec"
)

type TrafficSummary struct {
	MallID     string `json:"mall_id" url:"mall_id"`
	Token      string `json:"token" url:"token"`
	EntityType int    `json:"entity_type" url:"entityType"`
	StartTime  string `json:"start_time" url:"startTime"`
	EndTime    string `json:"end_time" url:"endTime"`
	Interval   string `json:"interval" url:"interval"`
	JobParam
}

type Entity struct {
	MallID     string `json:"maillId"`
	EntityName string `json:"entityName"`
	Day        string `json:"day"`
	TrafficIn  int    `json:"trafficIn"`
	TrafficOut int    `json:"trafficOut"`
	Visitors   int    `json:"visitors"`
}
type Data struct {
	List []Entity
}
type Result struct {
	RequestID string `json:"request_id"`
	Data      Data   `json:"data"`
}

type JobParam struct {
	StoreCode string `json:"store_code"`
	StoreName string `json:"store_name"`
}

type Job struct {
	req *req.Client
	q   *Queries
}

func New(aibeeURL string, q *Queries) *Job {
	req := req.C().SetBaseURL("https://face-event-api.aibee.cn")
	return &Job{
		req: req,
		q:   q,
	}
}

func (b *Job) Name() string {
	return "abg"
}

func (b *Job) Run(ctx context.Context, task *xxl.Task) error {
	yestodat := Yestoday(time.Now())
	ys := DateOnly(yestodat)
	param := TrafficSummary{
		EntityType: 70,
		Interval:   "D",
		EndTime:    ys,
		StartTime:  ys,

		JobParam: JobParam{
			StoreName: "N/A",
		},
	}
	if err := xxl.TaskJsonParam(task, &param); err != nil {
		return err
	}

	if param.StoreCode == "" {
		return errors.New("storecode is nil")
	}

	var r Result
	_, err := b.req.R().
		SetQueryParamsFromStruct(&param).
		SetSuccessResult(&r).
		Get("/traffic_summary")

	if err != nil {
		return err
	}

	var in int
	for _, ent := range r.Data.List {
		in = in + ent.TrafficIn
	}

	/*
	if in == 0 {
		return nil
	}
	*/

	err = b.q.CreateGmEntry(ctx, CreateGmEntryParams{
		StoreCode: param.StoreCode,
		StoreName: param.StoreName,
		DT:        DT(yestodat),
		InTotal:   in,
	})

	return err
}

func Yestoday(now time.Time) time.Time {
	return now.Add(-1 * 24 * time.Hour)
}

func DateOnly(t time.Time) string {
	return t.Format(time.DateOnly)
}

func DT(t time.Time) string {
	return t.Format("20060102")
}
