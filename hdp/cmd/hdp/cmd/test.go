package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/twiglab/xjob/hdp"
)

// runCmd represents the run command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return t()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}

func t() error {

	q, err := hdp.NewStore(
		viper.GetString("hdp.db.name"),
		viper.GetString("hdp.db.dsn"),
	)
	if err != nil {
		return err
	}
	defer q.Close()

	j := &hdp.App{
		Store: q,
	}

	o, err := j.GetOutLine()

	if err != nil {
		return err
	}

	fmt.Println(o)

	return err

}
