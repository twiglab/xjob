package be

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/tinylib/msgp/msgp"
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

func (h *Handle) SaveTarger(ctx context.Context, tp xe.TriggerParam) error {
	kv := clientv3.NewKV(h.client)

	var sb strings.Builder
	enc := json.NewEncoder(&sb)
	if err := enc.Encode(&tp); err != nil {
		return err
	}
	msgp.AppendBool
	_, err := kv.Put(ctx, "", sb.String())
	return err
}

func (h *Handle) ListTarger(ctx context.Context) ([]xe.TriggerParam, error) {
	kv := clientv3.NewKV(h.client)

	resp, err := kv.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	var xx []xe.TriggerParam
	for _, v := range resp.Kvs {
		var x xe.TriggerParam
		if err := json.Unmarshal(v.Value, &x); err == nil {
			xx = append(xx, x)
		}
	}
	return xx, nil
}

func (h *Handle) W(ctx context.Context) error {
	xxx := h.client.Watch(ctx, "")

	for x := range xxx {
		for _, e := range x.Events {
			e.Type = clientv3.EventTypeDelete
		}
	}
	return nil
}
