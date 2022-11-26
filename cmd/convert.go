/*
Copyright © 2022 Marcus Venicius Nunes Biava
*/
package cmd

import (
	"github.com/marcusbiava/path-win-2-wsl/app"
	"github.com/spf13/cobra"
)

// convertCmd represents the pathWin2Wsl command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		app.Run(path)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
