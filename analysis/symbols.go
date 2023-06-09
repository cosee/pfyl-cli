package analysis

import (
	"gitlab.cosee.biz/pfyl/pfyl-cli/cmd"
	"strings"
)

const (
	nmBinary = "arm-none-eabi-nm"
)

type SymbolTableConsumer interface {
	ConsumeSymbolTable(symbolTable []SymbolTableEntry) error
}

type SymbolTableEntry struct {
	Address    string
	Type       string
	SymbolName string
}

func SymbolsAnalyzerProvider(consumer SymbolTableConsumer) cmd.Analyzer {
	return func(toolchainPath string, binaryPath string) error {
		output, err := execute(toolchainPath, nmBinary, binaryPath, "-S", "--demangle", "--size-sort", "--defined-only", "--numeric-sort")
		if err != nil {
			return err
		}

		symbolTable := buildSymbolTable(output)
		return consumer.ConsumeSymbolTable(symbolTable)
	}
}

func buildSymbolTable(result string) []SymbolTableEntry {
	var entries []SymbolTableEntry
	lines := strings.Split(result, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		entries = append(entries, buildSymbolTableEntry(line))
	}

	return entries
}

func buildSymbolTableEntry(line string) SymbolTableEntry {
	segments := strings.Split(line, " ")
	return SymbolTableEntry{
		Address:    segments[0],
		Type:       segments[1],
		SymbolName: segments[2],
	}
}
