package configuration

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestConfiguration(t *testing.T) {
	t.Run("creates new file when no configuration exists", func(t *testing.T) {
		homeDir, err := os.UserHomeDir()
		assert.Nil(t, err)

		fs := afero.NewMemMapFs()
		New(fs)

		_, err = fs.Stat(path.Join(homeDir, "/.pfyl/config.json"))
		assert.True(t, !os.IsNotExist(err))
	})

	t.Run("loads configuration from existing file", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		config := New(fs)
		config.ToolChainPath = "toolchain"
		config.ServerAddress = "address"
		config.AccessToken = "token"
		err := config.Save()
		assert.Nil(t, err)

		configLoaded := New(fs)
		assert.Equal(t, config, configLoaded)
	})
}
