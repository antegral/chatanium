package main

import (
	IChatanium "antegral.net/chatanium/src/Runtime/Interface"
	"antegral.net/chatanium/src/Runtime/Log"
	"github.com/bwmarrin/discordgo"
)

// Modules can be built together in a way that is built into the source code.
// It also supports dynamic module insertion by compiling the module's source code into a DLL or SO.
// but, module must build with a "plugin" supported by Golang.
// For a detailed module insertion implementation, see `src/Runtime/Module/Local.go`.
//
// An example, this is a module that returns the chat entered.

// The runtime finds the `ChataniumModule` variable and type-casting it to a module it can use, so you can't rename the variable.
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

// This function is required when initializing the module.
//
// The runtime also runs this function in advance, before running OnStart().
//
// An example of what you might want to implement here would be getting environment variables.
func (t *Module) OnInit() error {
	Log.Info.Printf("ChatEcho: Init")
	return nil
}

// This function is required when starting the module.
//
// The runtime receives the channel from this function and processes it by sending incoming messages.
//
// An example of what you might want to implement here would be getting environment variables.
func (t *Module) OnStart() error {
	Log.Info.Printf("ChatEcho: Started")
	return nil
}

// A function that is called at runtime when information about the module is needed.
func (t *Module) GetInfo() *IChatanium.ModuleInfo {
	return &IChatanium.ModuleInfo{
		Name:        t.Name,
		Description: t.Description,
		Version:     t.Version,
		Tags:        t.Tags,
		Commands:    t.Commands,
	}
}

// Receive the chat backend modules supported by the runtime.
//
// Different chat platforms support different features and message types,
// so you can choose the chat platform that meets your needs.
func (t *Module) GetBackend(Backend IChatanium.Backend) error {
}

func (t *Module) OnMessage(Message any, IsFinished chan bool) {
}
