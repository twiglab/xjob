package aibee

import (
	"context"

	"github.com/imroc/req/v3"
)

const (
	BaseURL = "https://face-event-api.aibee.cn"
)

type TrafficSummaryIn struct {
	EntityType int    `json:"entity_type" url:"entityType"`
	StartTime  string `json:"start_time" url:"startTime"`
	EndTime    string `json:"end_time" url:"endTime"`
	Interval   string `json:"interval" url:"interval"`
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

type Auth struct {
	MallID string `json:"mall_id" url:"mall_id"`
	Token  string `json:"token" url:"token"`
}

type Client struct {
	req *req.Client

	auth Auth
}

func New(auth Auth) *Client {
	req := req.C().SetBaseURL(BaseURL)
	return &Client{
		req:  req,
		auth: auth,
	}
}

type param struct {
	Auth
	TrafficSummaryIn
}

func (b *Client) TrafficSummary(ctx context.Context, in TrafficSummaryIn) (out Result, err error) {
	_, err = b.req.R().
		SetQueryParamsFromStruct(&param{Auth: b.auth, TrafficSummaryIn: in}).
		SetSuccessResult(&out).
		Get("/traffic_summary")
	return
}
