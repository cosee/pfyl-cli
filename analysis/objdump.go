package analysis

import (
	"gitlab.cosee.biz/pfyl/pfyl-cli/cmd"
	"strings"
)

const objdumpBinary = "arm-none-eabi-objdump"

type ObjdumpConsumer interface {
	ConsumeObjdump(objdump string) error
}

func ObjdumpAnalyzerProvider(consumer ObjdumpConsumer) cmd.Analyzer {
	return func(toolchainPath string, binaryPath string) error {
		output, err := execute(toolchainPath, objdumpBinary, "-S", binaryPath)
		if err != nil {
			return err
		}

		strippedObjdump := stripObjdumpHeader(output)
		return consumer.ConsumeObjdump(strippedObjdump)
	}
}

func stripObjdumpHeader(output string) string {
	stripped := strings.Split(output, ".text:")[1]
	trimmedAndStrippedObjdump := strings.TrimSpace(stripped)
	return trimmedAndStrippedObjdump
}
