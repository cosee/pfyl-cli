package main

import (
	"gitlab.cosee.biz/pfyl/pfyl-cli/analysis"
	"gitlab.cosee.biz/pfyl/pfyl-cli/cmd"
	"gitlab.cosee.biz/pfyl/pfyl-cli/external"
)

func main() {
	client := external.NewClient()
	symbols := analysis.SymbolsProvider(client)
	rootCmd := cmd.NewRoot(cmd.Analyzer(symbols))
	rootCmd.Execute()
}
