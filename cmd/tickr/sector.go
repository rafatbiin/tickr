package main

import (
	"fmt"
	"github.com/rafatbiin/tickr"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func NewSectorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sector",
		Short: "get the sector of the listed company from the given ticker.",
		Example: strings.TrimLeft(`Command: tickr sector NBL
Response: bank`, "\n"),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 {
				extractSector(args[0])
			} else {
				fmt.Fprintln(os.Stderr, "please give no more than 1 ticker.")
				cmd.Help()
			}
		},
	}
	return cmd
}

func extractSector(ticker string) {
	t, err := tickr.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to extract ticker, error:"+err.Error())
	}
	sector := t.Sector(ticker)
	fmt.Fprintln(os.Stdout, sector)

}
