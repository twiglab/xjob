package be

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/twiglab/xjob/pkg/xe"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Handle struct {
	client *clientv3.Client
}

func (h *Handle) Reg(ctx context.Context, rp xe.RegistryParam) error {
	kv := clientv3.NewKV(h.client)
	l := clientv3.NewLease(h.client)
	lr, err := l.Grant(ctx, 30)
	if err != nil {
		return err
	}

	_, err = kv.Put(ctx, execKey(rp.RegistryGroup, rp.RegistryKey), rp.RegistryValue, clientv3.WithLease(lr.ID))

	return err
}

func (h *Handle) Ta(ctx context.Context, tp xe.TriggerParam) error {
	kv := clientv3.NewKV(h.client)

	var sb strings.Builder
	enc := json.NewEncoder(&sb)
	if err := enc.Encode(&tp); err != nil {
		return err
	}
	kv.Put(ctx, "", sb.String())
	return nil
}
