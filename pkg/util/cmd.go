package util

import (
	"errors"
	"os/exec"
)

// ExecCmd 快捷执行命令，返回标准输出流
func ExecCmd(name string, args []string) (output []byte, exitCode int, err error) {
	cmd := exec.Command(name, args...)
	output, err = cmd.Output()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			exitCode = exitError.ExitCode()
			return
		}
	}
	return
}
