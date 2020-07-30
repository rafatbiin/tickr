package main

import (
	"encoding/json"
	"fmt"
	"github.com/rafatbiin/tickr"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func NewTickerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ticker",
		Short: "extract ticker of the listed company(s) from given news data.",
		Example: strings.TrimLeft(`Command: tickr ticker national bank publishes 2nd quarterly report
Response: {"NBL":1}`, "\n"),
		Run: func(cmd *cobra.Command, args []string) {
			text := strings.Join(args, " ")
			extractTicker(text)
		},
	}
	return cmd
}

func extractTicker(text string) {
	t, err := tickr.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to parse news data, error:"+err.Error())
	}
	tickers := t.Get(text)
	ftickers, _ := json.Marshal(tickers)
	fmt.Fprintln(os.Stdout, string(ftickers))

}
