package cmd

import (
	"context"
	"net/http"

	"github.com/it512/xxl-job-exec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/twiglab/xjob/aibee"
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

	q, err := aibee.Open(
		viper.GetString("aibee.db.name"),
		viper.GetString("aibee.db.dsn"),
	)
	if err != nil {
		return err
	}

	defer q.Close()

	j := aibee.New("", q)
	exec.RegTask(j.Name(), task(j))

	if err := http.ListenAndServe(":10009", exec.Handle("/aibee")); err != nil {
		return err
	}

	return nil

}

func task(app *aibee.Job) xxl.TaskFunc {
	return func(ctx context.Context, task *xxl.Task) error {
		return app.Run(ctx, task)
	}
}
