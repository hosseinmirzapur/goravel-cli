package utils

import (
	"fmt"
	"os/exec"
)

func HandleOutput(cmd *exec.Cmd) error {
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}
