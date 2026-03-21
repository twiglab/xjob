/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/it512/xxl-job-exec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/twiglab/xjob/hdp"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func run() error {
	exec := xxl.NewExecutor(
		xxl.ServerAddr(viper.GetString("xxl.url")),
		xxl.AccessToken(viper.GetString("xxl.token")),
		xxl.RegistryKey(viper.GetString("xxl.key")),
		xxl.ExecutorURL(viper.GetString("xxl.local.url")),
	)
	exec.Start()
	defer func() { _ = exec.Stop() }()

	dbx, err := hdp.NewDBx(
		viper.GetString("hdp.db.name"),
		viper.GetString("hdp.db.dsn"),
	)
	if err != nil {
		log.Fatal(err)
	}
	summary := &hdp.Summary{
		DBx: dbx,
		Tpl: hdp.SummaryTpl(),
	}
	defer dbx.Close()
	exec.RegTask(summary.Name(), task(summary))

	if err := http.ListenAndServe(":10008", exec.Handle("/hdp")); err != nil {
		log.Fatal(err)
	}

	return nil
}

type job interface {
	Run(ctx context.Context, task *xxl.Task) error
}

func task(j job) xxl.TaskFunc {
	return func(ctx context.Context, task *xxl.Task) error {
		return j.Run(ctx, task)
	}
}
