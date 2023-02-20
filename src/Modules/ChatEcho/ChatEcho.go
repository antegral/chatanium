package ChatEcho

import (
	"antegral.net/chatanium/src/Runtime/Interface"
	"antegral.net/chatanium/src/Runtime/Log"
	"github.com/bwmarrin/discordgo"
)

const (
	Name        string = "ChatEcho"
	Description string = "Returns the chat entered"
	Version     string = "1.0.0"
)

var (
	Commands []string = nil
	Tags     []string = []string{"Echo"}
)

type Module struct {
	Name        string
	Description string
	Version     string
	Tags        []string
	Commands    []string
	discord     *discordgo.Session
}

func (t *Module) OnInit() error {
	Log.Info.Printf("ChatEcho: Init")

	t.Name = Name
	t.Description = Description
	t.Version = Version
	t.Tags = Tags
	t.Commands = Commands

	return nil
}

func (t *Module) OnStart() error {
	Log.Info.Printf("ChatEcho: Started")
	return nil
}

func (t *Module) GetInfo() *Interface.ModuleInfo {
	return &Interface.ModuleInfo{
		Name:        t.Name,
		Description: t.Description,
		Version:     t.Version,
		Tags:        t.Tags,
		Commands:    t.Commands,
	}
}

func (t *Module) OnMessage(Request string) Interface.ModuleResponse {
	return nil
}
