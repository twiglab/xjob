/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json/v2"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/twiglab/xjob/hdp"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return summary()
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}

func summary() error {

	dbx, err := hdp.NewDBx(
		viper.GetString("hdp.db.name"),
		viper.GetString("hdp.db.dsn"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer dbx.Close()

	j := &hdp.Summary{
		DBx: dbx,
		Tpl: hdp.SummaryTpl(),
	}
	param := hdp.SummaryParam{
		StoreCode: "1006",
		StoreName: "长乐金陵",
		BotKey:    "xxxxxx",
	}
	json.MarshalWrite(os.Stdout, param)
	fmt.Println()
	fmt.Println("-------------------------------")
	o, err := j.DoRun(context.Background(), param)

	if err != nil {
		log.Fatal(err)
	}

	j.Tpl.Execute(os.Stdout, o)

	return err

}
