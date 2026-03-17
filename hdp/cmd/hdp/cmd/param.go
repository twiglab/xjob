/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	"github.com/twiglab/xjob/hdp"
)

// paramCmd represents the param command
var paramCmd = &cobra.Command{
	Use:   "param",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var param hdp.AppParam
		param.Tags = []string{"1"}
		enc := json.NewEncoder(os.Stdout)
		enc.Encode(param)
	},
}

func init() {
	rootCmd.AddCommand(paramCmd)
}
