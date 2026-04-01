package wanf

import (
	"context"

	"github.com/imroc/req/v3"
)

const (
	JsonContentType = "application/json;charset=utf-8"
	ContentType     = "Content-Type"
)

type Config struct {
	Token   string
	BaseURL string
}

type Client struct {
	c    *req.Client
	conf Config
}

func New(conf Config) *Client {
	c := req.C().SetBaseURL(conf.BaseURL).OnBeforeRequest(auth(conf))
	return &Client{c: c, conf: conf}
}

func auth(conf Config) req.RequestMiddleware {
	return func(client *req.Client, req *req.Request) error {
		req.SetHeader("Authorization", "token "+conf.Token)
		return nil
	}
}

func (c *Client) PeopleCounting(ctx context.Context, in PeopleCountingIn) (out PeopleCountingOut, err error) {
	_, err = c.c.R().
		SetContext(ctx).
		SetContentType(JsonContentType).
		SetSuccessResult(&out).
		SetBody(in).
		SetErrorResult(&out).
		Post("/wanf-api/people-counting-data/")
	return
}
