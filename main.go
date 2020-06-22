package main

import (
	"github.com/spf13/afero"
	"gitlab.cosee.biz/pfyl/pfyl-cli/analysis"
	"gitlab.cosee.biz/pfyl/pfyl-cli/cmd"
	"gitlab.cosee.biz/pfyl/pfyl-cli/configuration"
	"gitlab.cosee.biz/pfyl/pfyl-cli/external"
)

func main() {
	config := configuration.New(afero.NewOsFs())
	client := external.NewClient(config)
	symbols := analysis.SymbolsProvider(client)
	configureCmd := cmd.NewConfigure(config)
	rootCmd := cmd.NewRoot(config, cmd.Analyzer(symbols))
	rootCmd.AddCommands(configureCmd)
	rootCmd.Execute()
}
