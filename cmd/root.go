package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.cosee.biz/pfyl/pfyl-cli/configuration"
	"log"
)

type Analyzer func(toolchainPath string, binaryPath string) error


func NewRoot(config configuration.Configuration,analyzers ...Analyzer) *cobra.Command {
	return &cobra.Command{
		Use:   "pfyl-cli",
		Short: "pfyl-cli analyzes binaries and sends results to a pfyl-server",
		Run: func(cmd *cobra.Command, args []string) {
			for _, analyzer := range analyzers {
				err := analyzer(config.ToolChainPath, "test/executables/f7-device")
				if err != nil {
					log.Print(err)
				}
			}
		},
	}
}
