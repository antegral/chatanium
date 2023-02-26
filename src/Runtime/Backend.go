package runtime

import (
	"os"

	Backends "antegral.net/chatanium/src/Backend"
	IChatanium "antegral.net/chatanium/src/Runtime/Interface"
	"antegral.net/chatanium/src/Runtime/Log"
)

func GetBackend(BackendType string, Module IChatanium.ModuleInfo) IChatanium.Backend {
	Log.Info.Print(Module.Name, " > Access granted to backend")

	switch BackendType {
	case "discord":
		Backend := Backends.Discord{}
		Backend.Init(Module)
		// TODO: 모듈에서 Credential을 알 수 없게 privately하게 짜기
		// DotEnv는 모듈측에서 os.Getenv로 읽을 수 있으므로 사용 할 수 없음.
		// 1. 경로 방식으로 privately하게 env 파일을 가져오고 로컬에서만 쓰면 되지 않을까?
		// 1-1. 프로그램 실행 인수로 암호화된 env 파일의 key를 가져오게?
		Backend.SetCredentials(os.Getenv("CHATANIUM_TEST_DISCORD_TOKEN"))
		Backend.Connect()
		return &Backend
	default:
		break
	}

	Log.Error.Fatal(Module.Name, " > Unknown backend: ", BackendType)
	return nil
}
