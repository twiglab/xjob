/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"net/http"

	"github.com/it512/xxl-job-exec"
	"github.com/spf13/cobra"
	"github.com/twiglab/xjob/flow/hik"
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
	exec := xxlexec()
	exec.Start()
	defer func() { _ = exec.Stop() }()

	j := &hik.HikJob{
		DBx: dbx(),
	}

	var push hik.CfasPushBot
	exec.RegTask(j.Name(), task(j))
	exec.RegTask(push.Name(), task(push))

	return http.ListenAndServe(webaddr(), exec.Handle(handlepath()))
}

type job interface {
	Run(ctx context.Context, task *xxl.Task) error
}

func task(j job) xxl.TaskFunc {
	return func(ctx context.Context, task *xxl.Task) error {
		return j.Run(ctx, task)
	}
}
