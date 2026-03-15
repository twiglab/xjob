package flow

import (
	"context"
	"time"

	"github.com/twiglab/xjob/flow/orm/ent"
	"github.com/twiglab/xjob/pfsdk"
)

func PK(code, dt string) string {
	return code + "-" + dt + "-" + "in"
}

type DBx struct {
	Client *ent.Client
}

func (d *DBx) Save(ctx context.Context, yestoday time.Time, p Param, inTotal int) error {
	dt := pfsdk.DT(yestoday)
	pk := PK(p.StoreCode, dt)

	cr := d.Client.Flow.Create()
	cr.SetDt(dt)
	cr.SetPk(pk)
	cr.SetStoreCode(p.StoreCode)
	cr.SetStoreName(p.StoreName)
	cr.SetInTotal(inTotal)
	return cr.Exec(ctx)
}
