package cmd

import (
	"log"

	"github.com/it512/xxl-job-exec"
	"github.com/spf13/viper"
	"github.com/twiglab/xjob/hdp"
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

func dbx() *hdp.DBx {
	dbx, err := hdp.NewDBx(
		viper.GetString("hdp.db.name"),
		viper.GetString("hdp.db.dsn"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return dbx
}

func holidays() hdp.Holidays {
	file := viper.GetString("hdp.holiday.file")
	hs, err := hdp.LoadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return hs
}
