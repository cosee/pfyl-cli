package analysis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	binariesPath   = "../test/binaries"
	executablePath = "../test/executables/f7-device"
)

type SymbolTableConsumerMock func(symbolTable []SymbolTableEntry) error

func (s SymbolTableConsumerMock) ConsumeSymbolTable(symbolTable []SymbolTableEntry) error {
	return s(symbolTable)
}

func TestSymbolsAnalyzer(t *testing.T) {
	assertSymbolTableEntry := func(entry SymbolTableEntry, address string, symbolType string, symbolName string) {
		assert.Equal(t, address, entry.Address)
		assert.Equal(t, symbolType, entry.Type)
		assert.Equal(t, symbolName, entry.SymbolName)
	}

	t.Run("building symbol table succeeds", func(t *testing.T) {
		mock := func(symbolTable []SymbolTableEntry) error {
			assertSymbolTableEntry(symbolTable[0], "00000000", "A", "_Min_Heap_Size")
			assertSymbolTableEntry(symbolTable[1], "00000200", "A", "_Min_Stack_Size")
			assertSymbolTableEntry(symbolTable[2], "08000000", "R", "g_pfnVectors")
			return nil
		}

		symbolsAnalyzer := SymbolsAnalyzerProvider(SymbolTableConsumerMock(mock))
		err := symbolsAnalyzer(binariesPath, executablePath)
		assert.Nil(t, err)
	})
}
