package analysis

import (
	"os/exec"
	"path"
)

func execute(toolchainPath string, binary string, args ...string) (string, error){
	binaryPath := path.Join(toolchainPath, binary)
	command := exec.Command(binaryPath, args...)
	output, err := command.CombinedOutput()
	return string(output), err
}
