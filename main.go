package main

import (
	"github.com/spf13/afero"
	"gitlab.cosee.biz/pfyl/pfyl-cli/analysis"
	"gitlab.cosee.biz/pfyl/pfyl-cli/cmd"
	"gitlab.cosee.biz/pfyl/pfyl-cli/configuration"
	"gitlab.cosee.biz/pfyl/pfyl-cli/external"
	"log"
)

func main() {
	config := configuration.New(afero.NewOsFs())
	client := external.NewClient(config)

	symbolsAnalyzer := analysis.SymbolsAnalyzerProvider(client)
	objdumpAnalyzer := analysis.ObjdumpAnalyzerProvider(client)

	rootCmd := cmd.NewRoot(config, symbolsAnalyzer, objdumpAnalyzer)
	rootCmd.AddCommand(cmd.NewConfigure(config))
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
