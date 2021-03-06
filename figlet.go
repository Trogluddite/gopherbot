package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type FigletPlugin struct{}

func (p FigletPlugin) Name() string {
	return "figlet"
}
func (p FigletPlugin) Execute(command []string) string {
	if len(command) < 1 {
		return "herp"
	}
	figletCmd := exec.Command("figlet", strings.Join(command, " "))
	figletOut, err := figletCmd.Output()
	if err != nil {
		return fmt.Sprintf("Couldn't get figlet output. ERR %v", err)
	}
	return "```" + string(figletOut) + "```"
}
