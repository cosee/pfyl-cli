package analysis

import (
	"gitlab.cosee.biz/pfyl/pfyl-cli/cmd"
	"math"
	"strconv"
	"strings"
)

const (
	sectionsBinary             = "arm-none-eabi-objdump"
	contents       SectionFlag = "CONTENTS"
	alloc                      = "ALLOC"
	load                       = "LOAD"
	readonly                   = "READONLY"
	data                       = "DATA"
	code                       = "CODE"
	debugging                  = "DEBUGGING"
)

type SectionsTable struct {
	Sections []Section
}

type SectionFlag string

type Section struct {
	Index     int64
	Name      string
	Size      int64
	VMA       int64
	LMA       int64
	FileOff   int64
	Alignment int
	Flags     []SectionFlag
}

type SectionTableConsumer interface {
	ConsumeSectionTable(table SectionsTable) error
}

func SectionsAnalyzerProvider(consumer SectionTableConsumer) cmd.Analyzer {
	return func(toolchainPath string, binaryPath string) error {
		output, err := execute(toolchainPath, sectionsBinary, "-h", binaryPath)
		if err != nil {
			return err
		}

		stripped := stripSectionHeader(output)
		sectionTable, err := buildSectionTable(stripped)
		if err != nil {
			return err
		}

		return consumer.ConsumeSectionTable(sectionTable)
	}
}

func stripSectionHeader(output string) string {
	return strings.Split(output, "Algn\n")[1]
}

func buildSectionTable(stripped string) (SectionsTable, error) {
	lines := strings.Split(stripped, "\n")

	var sections []Section
	for index := 0; index < len(lines)-1; index += 2 {
		info := lines[index]
		flags := lines[index+1]

		section, err := buildSection(info, flags)
		if err != nil {
			return SectionsTable{}, err
		}

		sections = append(sections, section)
	}

	return SectionsTable{Sections: sections}, nil
}

func buildSection(info string, flags string) (Section, error) {
	sectionFlags := parseFlags(flags)
	infoElements := strings.Fields(info)

	index, err := strconv.ParseInt(infoElements[0], 10, 64)
	if err != nil {
		return Section{}, err
	}

	name := infoElements[1]

	size, err := strconv.ParseInt(infoElements[2], 16, 64)
	if err != nil {
		return Section{}, err
	}

	vma, err := strconv.ParseInt(infoElements[3], 16, 64)
	if err != nil {
		return Section{}, err
	}

	lma, err := strconv.ParseInt(infoElements[4], 16, 64)
	if err != nil {
		return Section{}, err
	}

	fileOff, err := strconv.ParseInt(infoElements[5], 16, 64)
	if err != nil {
		return Section{}, err
	}

	baseAndPow := strings.Split(infoElements[6], "**")
	base, err := strconv.ParseInt(baseAndPow[0], 10, 64)
	if err != nil {
		return Section{}, err
	}

	pow, err := strconv.ParseInt(baseAndPow[1], 10, 64)
	if err != nil {
		return Section{}, err
	}

	alignment := math.Pow(float64(base), float64(pow))

	return Section{
		Index:     index,
		Name:      name,
		Size:      size,
		VMA:       vma,
		LMA:       lma,
		FileOff:   fileOff,
		Alignment: int(alignment),
		Flags:     sectionFlags,
	}, nil
}

func parseFlags(flagsStr string) []SectionFlag {
	flags := strings.Split(flagsStr, ",")

	var sectionFlags []SectionFlag
	for _, flag := range flags {
		cleanedFlag := strings.TrimSpace(flag)
		sectionFlags = append(sectionFlags, SectionFlag(cleanedFlag))
	}

	return sectionFlags
}
