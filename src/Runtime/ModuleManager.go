package runtime

import (
	Backends "antegral.net/chatanium/src/Backend"
	"antegral.net/chatanium/src/Runtime/Interface"
	"antegral.net/chatanium/src/Runtime/Log"
)

func CallModule(Module *ChataniumModule) {
	if err := Module.OnInit(); err != nil {
		Log.Error.Printf("Runtime > Some Module occurred Error during OnInit()")
		Log.Error.Fatalln(err)
	}

	if err := Module.OnStart(); err != nil {
		Info := Module.GetInfo()
		Log.Error.Printf(Info.Name, " > Error occurred during OnStart()")
		Log.Error.Fatalln(err)
	}

	// TODO: 모듈과 런타임간 통신을 위해 모듈에 대한 채널 반환
}

func GetBackend(BackendType string, Module Interface.ModuleInfo) *Interface.ChataniumBackend {
	Log.Info.Print(Module.Name, " > Access granted to backend")
	switch BackendType {
	case "discord":
		return &Backends.Discord{}
		break
	default:
		Log.Error.Fatal(Module.Name, " > Unknown backend: ", BackendType)
		return nil
		break
	}
}

type ChataniumModule struct {
	Info Interface.ModuleInfo
}

func (t *ChataniumModule) OnInit() error

func (t *ChataniumModule) OnStart() error

func (t *ChataniumModule) GetInfo() *Interface.ModuleInfo

func (t *ChataniumModule) GetBackend(Backend Interface.ChataniumBackend) error

func (t *ChataniumModule) OnMessage(Request string) Interface.ModuleResponse
