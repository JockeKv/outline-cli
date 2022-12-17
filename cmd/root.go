/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"outline/pkg/config"

	"github.com/spf13/cobra"
)

var conf config.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "outline",
	Short: "CLI for Outline",
	Long:  `CLI for Outline`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	conf = config.Config{}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
