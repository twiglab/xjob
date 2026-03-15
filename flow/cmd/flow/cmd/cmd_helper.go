package cmd

import (
	"cmp"
	"log"

	"github.com/it512/xxl-job-exec"
	"github.com/spf13/viper"
	"github.com/twiglab/xjob/flow"
	"github.com/twiglab/xjob/flow/orm"
	"github.com/twiglab/xjob/flow/orm/ent"
)

func xxlexec() *xxl.Executor {
	exec := xxl.NewExecutor(
		xxl.ServerAddr(viper.GetString("xxl.url")),
		xxl.AccessToken(viper.GetString("xxl.token")),
		xxl.RegistryKey(viper.GetString("xxl.key")),
		xxl.ExecutorURL(viper.GetString("xxl.local.url")),
	)

	return exec
}

func webaddr() string {
	addr := viper.GetString("flow.web.addr")
	return cmp.Or(addr, ":10009")
}

func handlepath() string {
	return viper.GetString("xxl.local.path")
}

func entcli() *ent.Client {
	name := viper.GetString("flow.db.name")
	dsn := viper.GetString("flow.db.dsn")

	//cli, err := orm.OpenEntClient(name, dsn, ent.Debug())
	cli, err := orm.OpenEntClient(name, dsn)
	if err != nil {
		log.Fatal(err)
	}
	return cli
}

func dbx() *flow.DBx {
	return &flow.DBx{Client: entcli()}
}
