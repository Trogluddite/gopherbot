package main

import "strings"

type UdPlugin struct{}

func (p UdPlugin) Name() string {
	return "ud"
}

func (p UdPlugin) Execute(command []string) string {
	if len(command) < 1 {
		return "herp"
	}
	query := strings.Join(command, "%20")
	return "http://www.urbandictionary.com/define.php?term=" + query
}
