package configuration

import (
	"encoding/json"
	"github.com/spf13/afero"
	"os"
	"path"
)

const configurationDirectory = ".pfyl"
const configurationFile = "config.json"
const dirPermissions = 0755
const filePermissions = 0644

type Configuration struct {
	ToolChainPath string
	ServerAddress string
	AccessToken   string
	fs            afero.Fs
}

func New(fs afero.Fs) Configuration {
	var config Configuration
	config.fs = fs
	(&config).load()
	return config
}

func (c *Configuration) load() {
	pfylConfigurationDir := c.buildFullConfigurationDirectoryPath()
	_, err := c.fs.Stat(pfylConfigurationDir)
	if !os.IsNotExist(err) {
		c.processConfigurationDirectory(pfylConfigurationDir)
		return
	} else if os.IsNotExist(err) {
		err := c.fs.Mkdir(pfylConfigurationDir, dirPermissions)
		if err != nil {
			panic(err)
		}

		c.processConfigurationDirectory(pfylConfigurationDir)
		return
	}

	panic(err)
}

func (c *Configuration) buildFullConfigurationDirectoryPath() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return path.Join(userHomeDir, configurationDirectory)
}

func (c *Configuration) processConfigurationDirectory(directory string) {
	fullPath := path.Join(directory, configurationFile)
	_, err := c.fs.Stat(fullPath)
	if !os.IsNotExist(err) {
		c.readConfigurationFile(fullPath)
		return
	} else if os.IsNotExist(err) {
		c.createConfigurationFile(fullPath)
		return
	}

	panic(err)
}

func (c *Configuration) readConfigurationFile(fullPath string) {
	data, err := afero.ReadFile(c.fs, fullPath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
}

func (c *Configuration) createConfigurationFile(fullPath string) {
	data, err := json.Marshal(&c)
	if err != nil {
		panic(err)
	}

	err = afero.WriteFile(c.fs, fullPath, data, filePermissions)
	if err != nil {
		panic(err)
	}

}

func (c Configuration) Save() error {
	fullPath := path.Join(c.buildFullConfigurationDirectoryPath(), configurationFile)
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return afero.WriteFile(c.fs, fullPath, data, filePermissions)
}
