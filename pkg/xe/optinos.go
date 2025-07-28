package xe

import (
	"context"
	"log/slog"
)

type Options struct {
	ServerAddr  string `json:"server_addr"`  //调度中心地址
	AccessToken string `json:"access_token"` //请求令牌
	ExecutorURL string `json:"executor_url"` //本地(执行器)URL
	RegistryKey string `json:"registry_key"` //执行器名称

	log     *slog.Logger
	client  Doer
	rootCtx context.Context
}

type Option func(o *Options)

var (
	DefaultExecutorPort = "19999"
	DefaultRegistryKey  = "golang-jobs-plus"
	DefaultAccessToken  = "default_token"
)

// ServerAddr 设置调度中心地址
func ServerAddr(addr string) Option {
	return func(o *Options) {
		o.ServerAddr = addr
	}
}

// AccessToken 请求令牌
func AccessToken(token string) Option {
	return func(o *Options) {
		o.AccessToken = token
	}
}

// ExecutorURL 设置执行器URL
func ExecutorURL(url string) Option {
	return func(o *Options) {
		o.ExecutorURL = url
	}
}

// RegistryKey 设置执行器标识
func RegistryKey(registryKey string) Option {
	return func(o *Options) {
		o.RegistryKey = registryKey
	}
}

// SetLogger 设置日志处理器
func SetLogger(l *slog.Logger) Option {
	return func(o *Options) {
		o.log = l
	}
}

func SetContext(ctx context.Context) Option {
	return func(o *Options) {
		o.rootCtx = ctx
	}
}

func SetHttpClient(doer Doer) Option {
	return func(o *Options) {
		o.client = doer
	}
}
