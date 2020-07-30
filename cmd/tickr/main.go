package main

import (
	"github.com/spf13/cobra"
	"os"
)

func main() {
	cmd := NewCmdRoot()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func NewCmdRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "tickr",
		Short: "command line tool for tickr",
		Long: `It's a CLI tool for inspecting the tickr package,
which provides news categorization of listed companies in Bangladesh stock market.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.AddCommand(NewTickerCmd())
	rootCmd.AddCommand(NewSectorCmd())
	return rootCmd
}
