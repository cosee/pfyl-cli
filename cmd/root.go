package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.cosee.biz/pfyl/pfyl-cli/configuration"
	"log"
)

type Analyzer func(toolchainPath string, binaryPath string) error

func NewRoot(config configuration.Configuration, analyzers ...Analyzer) *cobra.Command {
	return &cobra.Command{
		Use:   "pfyl-cli [program]",
		Short: "pfyl-cli analyzes a program and sends the results to your pfyl-server",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			binaryPath := args[0]
			for _, analyzer := range analyzers {
				err := analyzer(config.ToolChainPath, binaryPath)
				if err != nil {
					log.Print(err)
				}
			}
		},
	}
}
