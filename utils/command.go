package utils

import (
	"fmt"
	"os/exec"
)

func handleOutput(cmd *exec.Cmd) error {
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func RunCommand(commands []string) error {
	cmdToRun := exec.Command(commands[0], commands[1:]...)
	return handleOutput(cmdToRun)
}

func GoRunDot(commands []string) error {
	cmdToRun := exec.Command("go", "run", ".")
	cmdToRun.Args = append(cmdToRun.Args, commands...)
	return handleOutput(cmdToRun)
}
