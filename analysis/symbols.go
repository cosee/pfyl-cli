package analysis

import (
	"os/exec"
	"path"
	"strings"
)

const (
	nmBinary = "arm-none-eabi-nm"
)

type Symbols func(toolchainPath string, binaryPath string) error

type SymbolTableConsumer interface {
	ConsumeSymbolTable(symbolTable []SymbolTableEntry) error
}

type SymbolTableEntry struct {
	Address    string
	Type       string
	SymbolName string
}

func SymbolsProvider(consumer SymbolTableConsumer) Symbols {
	return func(toolchainPath string, binaryPath string) error {
		nmBinaryPath := path.Join(toolchainPath, nmBinary)
		command := exec.Command(nmBinaryPath, binaryPath, "-S", "--demangle", "--size-sort", "--defined-only", "--numeric-sort")
		output, err := command.CombinedOutput()
		if err != nil {
			return err
		}

		symbolTable := buildSymbolTable(string(output))
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
