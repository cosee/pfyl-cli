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

	symbolsAnalyzer := analysis.SymbolsAnalyzerProvider(client)
	objdumpAnalyzer := analysis.ObjdumpAnalyzerProvider(client)

	configureCmd := cmd.NewConfigure(config)
	rootCmd := cmd.NewRoot(config, symbolsAnalyzer, objdumpAnalyzer)
	rootCmd.AddCommands(configureCmd)
	rootCmd.Execute()
}
