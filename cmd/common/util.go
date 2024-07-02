package common

import (
	"fmt"
	"os/exec"
)

func Bash(cmd string) error {
	fmt.Println("Exec command: ", cmd)
	output, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println("Failed to execute command:", err)
		return err
	}
	fmt.Println(string(output))
	return nil
}

func Restart() error {
	return Bash("systemctl restart gost.service")
}
