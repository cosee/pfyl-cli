package analysis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type SectionTableConsumerMock func(table SectionsTable) error

func (s SectionTableConsumerMock) ConsumeSectionTable(table SectionsTable) error {
	return s(table)
}

func TestSectionsAnalyzer(t *testing.T) {
	assertSection := func(t *testing.T, section Section, index int64, name string, size int64, vma int64, lma int64, fileOff int64, alignment int, flags ...SectionFlag) {
		assert.Equal(t, index, section.Index)
		assert.Equal(t, name, section.Name)
		assert.Equal(t, size, section.Size)
		assert.Equal(t, vma, section.VMA)
		assert.Equal(t, lma, section.LMA)
		assert.Equal(t, fileOff, section.FileOff)
		assert.Equal(t, alignment, section.Alignment)

		for _, flag := range flags {
			assert.Contains(t, section.Flags, flag)
		}
	}

	t.Run("generating section table succeeds", func(t *testing.T) {
		mock := func(table SectionsTable) error {
			assertSection(t, table.Sections[0], 0, ".isr_vector", 456, 134217728, 134217728, 65536, 1, contents, alloc, load, readonly, data)
			assertSection(t, table.Sections[1], 1, ".text", 25116, 134218192, 134218192, 66000, 16, contents, alloc, load, readonly, code)
			assertSection(t, table.Sections[2], 2, ".rodata", 532, 134243308, 134243308, 91116, 4, contents, alloc, load, readonly, data)

			assert.Len(t, table.Sections, 18)
			return nil
		}

		sectionsAnalyzer := SectionsAnalyzerProvider(SectionTableConsumerMock(mock))
		err := sectionsAnalyzer(binariesPath, executablePath)
		assert.Nil(t, err)
	})
}
