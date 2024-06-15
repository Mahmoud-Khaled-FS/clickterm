/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clickterm",
	Short: "Manage your ClickUp tasks efficiently from the command line with our versatile tool.",
	Long:  "Manage your ClickUp tasks efficiently from the command line with our versatile tool.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	rootCmd.AddCommand(keyCmd())
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
