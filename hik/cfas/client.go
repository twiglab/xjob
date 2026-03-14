package cfas

import (
	"context"
	"net/http"
	"time"

	"github.com/imroc/req/v3"
)

const (
	JsonContentType = "application/json;charset=utf-8"
	ContentType     = "Content-Type"
)

type Client struct {
	c    *req.Client
	conf Config
}

func New(conf Config) *Client {
	c := req.C().SetBaseURL(conf.BaseURL).OnBeforeRequest(aksk(conf))
	return &Client{c: c, conf: conf}
}

func (c *Client) Do(req *http.Request) (resp *http.Response, err error) {
	resp, err = c.c.Do(req)
	return
}

func (c *Client) CountGroup(ctx context.Context, in CountGroupIn) (out CountGroupOutRtn, err error) {
	_, err = c.c.R().
		SetContext(ctx).
		SetContentType(JsonContentType).
		SetSuccessResult(&out).
		SetBody(in).
		SetErrorResult(&out).
		Post("/artemis/api/cfas/v2/countGroup/groups/page")
	return
}

func (c *Client) PassengerFlow(ctx context.Context, in PassengerFlowIn) (out PassengerFlowRtn, err error) {
	_, err = c.c.R().
		SetContext(ctx).
		SetContentType(JsonContentType).
		SetSuccessResult(&out).
		SetBody(in).
		SetErrorResult(&out).
		Post("/artemis/api/cfas/v2/passengerFlow/groups")
	return
}

func (c Client) Collect(ctx context.Context, start, end time.Time) (in int, out int, keep int, err error) {
	pf := PassengerFlowIn{IDs: "1", Granularity: "minutely", StartTime: start, EndTime: end}
	var pfr PassengerFlowRtn
	pfr, err = c.PassengerFlow(ctx, pf)
	if err != nil {
		return
	}

	in, out, keep = collect(pfr.Data.List)
	return
}
