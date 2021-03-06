package main

import (
	"strings"

	"github.com/daph/goslack"
)

type Plugin interface {
	Name() string
	Execute(args []string) string
}

func handleMessage(msg goslack.Event, client *goslack.Client) {
	// If the message isn't of type message or is sent my the bot user or whas not sent by @<user>
	// then return
	if msg.Type != "message" || msg.User == client.Self.Id || len(strings.Split(msg.Text, "")) <= 0 || strings.Split(msg.Text, "")[0] != "!" {
		return
	}

	plugins := []Plugin{GiphyPlugin{}, RandomgifPlugin{}, FigletPlugin{}, UdPlugin{}}

	temp := strings.Join(strings.Split(msg.Text, "")[1:], "")
	command := strings.Split(temp, " ")
	if len(command) <= 0 {
		client.PushMessage(msg.Channel, "herp")
		return
	}

	for _, v := range plugins {
		if command[0] == v.Name() {
			client.PushMessage(msg.Channel, v.Execute(command[1:]))
			return
		}
	}

	client.PushMessage(msg.Channel, "derp")
}
