package analysis

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const startOfObjdump = "080001d0 <__do_global_dtors_aux>:\n 80001d0:\tb510      \tpush\t{r4, lr}"

type ObjdumpConsumerMock func(objdump string) error

func (o ObjdumpConsumerMock) ConsumeObjdump(objdump string) error {
	return o(objdump)
}

func TestObjdumpAnalyzer(t *testing.T) {
	t.Run("generating cleaned objdump succeeds", func(t *testing.T) {
		mock := func(objdump string) error {
			assert.True(t, strings.HasPrefix(objdump, startOfObjdump))
			return nil
		}

		objdumpAnalyzer := ObjdumpAnalyzerProvider(ObjdumpConsumerMock(mock))
		err := objdumpAnalyzer(binariesPath, executablePath)
		assert.Nil(t, err)
	})
}
