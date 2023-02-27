package main

import (
	IChatanium "antegral.net/chatanium/src/Runtime/Interface"
	"antegral.net/chatanium/src/Runtime/Log"
	"github.com/bwmarrin/discordgo"
)

var ChataniumModule = Module{
	Name:        "ChatEcho",
	Description: "Returns the chat entered",
	Version:     "1.0.0",
	Tags:        nil,
	Commands:    nil,
}

type Module struct {
	Name        string
	Description string
	Version     string
	Tags        []string
	Commands    []string
	Discord     *discordgo.Session
}

func (t *Module) OnInit() error {
	Log.Info.Printf("ChatEcho: Init")
	return nil
}

func (t *Module) OnStart() error {
	Log.Info.Printf("ChatEcho: Started")
	return nil
}

func (t *Module) GetInfo() *IChatanium.ModuleInfo {
	return &IChatanium.ModuleInfo{
		Name:        t.Name,
		Description: t.Description,
		Version:     t.Version,
		Tags:        t.Tags,
		Commands:    t.Commands,
	}
}

func (t *Module) GetBackend(Backend IChatanium.Backend) error {
}

func (t *Module) OnMessage(Message any, IsFinished chan bool) {
}
