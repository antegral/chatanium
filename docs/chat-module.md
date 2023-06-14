> 현재 우리는 다양한 언어를 지원하기 위해 여러 노력을 하고 있지만,
> 런타임이 어느정도 완성되지 않았기에 Golang만 지원하고 있습니다.
## Get backend modules
Backend Module을 얻기 위해선, 그저 Module의 정보에 필요한 모듈을 적으세요.
`ChataniumModule.RequireBackends`에 Backend Module의 이름을 넣기만 하면 됩니다!
그렇게 하면, 런타임에서 자동으로 Backend Module을 얻고, `GetBackend()` 함수에 전달합니다.

크로스-플랫폼 연동을 위해 여러 개의 모듈이 필요하다구요? 괜찮습니다.
그저 더 많은 모듈을 원한다면, 더 많은 Backend Module의 이름들을 적으세요.
배열에 들어간 순서대로 모듈들이 전달 될 것입니다.

너무 복잡한가요? 여기 순서대로 한줄씩 정리했습니다!
1. `ChataniumModule.RequireBackends` 에 `string[]` 형태로 필요한만큼 모듈의 이름을 적는다.
2. 런타임은 `ChataniumModule.GetBackend(BackendModule[])` 로 Backend Module을 전달한다.
## Usage for Backend module
얻은 Backend Module은 어떻게 사용할까요?
답은 간단합니다. `ChataniumModule.GetBackend(BackendModule[])`에서 받아온 Backend Module은
포함되어 있는 서드파티 Wrapper 라이브러리의 모든 기능을 제약없이 사용할 수 있습니다.
당장 `BackendModule[0].Api` 에 접근해보세요!
## For Safety
모듈 사용자에 대한 채팅 API의 Credential 유출을 막기 위해 기본적으로 런타임에서는 Credential을  Chatbot Module에 전달 하지 않습니다.

따라서 이점을 염두하고 사용해야 할 것입니다.
## Receiving Message
서드파티 Wrapper 라이브러리에서 받아오는 메세지는 런타임에 먼저 전달됩니다.
이후로, `ChataniumModule.OnMessage(Message, BackendModuleName)` 함수를 호출하면서 메세지를 전달하게 됩니다.
### What of Message type?
해당되는 Backend Module에 대한 서드파티 Wrapper 라이브러리의 메세지 인터페이스 형태 그대로 받아오게 됩니다!

Chat Module은 `BackendModuleName`을 통해 무슨 Backend Module인지 알 수 있고 그에 맞게 행동하도록 코드를 만들 수 있습니다.

일반적으로 서드파티 Wrapper 라이브러리는 채팅 수신을 위해 그 메신저 만의 독특한 메시지 형태를 전부 구현하는데, 무려 그 인터페이스 그대로 받아옵니다!

따라서, 일반적으로 서드파티 Wrapper 라이브러리에 있는 인터페이스 그대로 모듈에 넣어도 문제가 발생하지 않습니다. (런타임에서 빌트인된 모든 Backend Modules)

하지만, 타입이 맞지 않아 오류를 발생시키거나, 어떤 타입인지 모르겠다면, Backend Module의 코드를 가져와 인터페이스만 따오는 것도 방법이 될 수 있습니다. (서드파티 Backend Module을 사용한다면 이 방법을 강력히 권장합니다!)
## Structure
이제 간단히 모듈의 구조를 하나씩 자세히 살펴보겠습니다.
```
package main

import (
	IChatanium "antegral.net/chatanium/src/Runtime/Interface"
	"antegral.net/chatanium/src/Runtime/Log"
	"github.com/bwmarrin/discordgo"
)

var ChataniumModule = Module{
	Name:        "ChatEcho",
	Description: "Returns the chat entered",
	Version:     "1.0.0",
	Tags:        nil,
	Commands:    nil,
}

type Module struct {
	Name        string
	Description string
	Version     string
	Tags        []string
	Commands    []string
	Discord     *discordgo.Session
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
		Name:        t.Name,
		Description: t.Description,
		Version:     t.Version,
		Tags:        t.Tags,
		Commands:    t.Commands,
	}
}

func (t *Module) GetBackend(Backend IChatanium.Backend) error {
}

func (t *Module) OnMessage(Message any, IsFinished chan bool) {
}
```

```
var ChataniumModule = Module{
	Name:        "ChatEcho",
	Description: "Returns the chat entered",
	Version:     "1.0.0",
	Tags:        nil,
	Commands:    nil,
}
```
이 코드는 모듈의 선언부입니다. 모듈에 꼭 필요한 정보들이 들어가 있습니다.
사실 이렇게 Struct를 initialize 할때 직접 적지 않아도 됩니다. 이건 나중에 `GetInfo()`에서 다시 설명하도록 하겠습니다.

Chatanium 런타임은 `ChataniumModule` 라는 변수를 읽어 모듈을 로딩합니다.
만약 이름이 다르거나, 일치하지 않으면 런타임은 인식하지 않습니다.
```
func (t *Module) OnInit() error {
	Log.Info.Printf("ChatEcho: Init")
	return nil
}
```
이 코드는 모듈이 로딩될 때 실행되는 함수입니다.

여기에서 모듈이 로딩될 때 필요한 작업들(실행 환경 확인)을 하고, 에러를 반환 할 수 있습니다.
이 때 반환되는 에러는 런타임을 중단시킵니다.
```
Runtime > ImportModule (Failed)
	Module initialize failed: Catched on Module "ChatEcho"
	ChatEcho: Invaild Environment. Terminating...

<Stack trace more...>
```

ChatEcho 모듈에서 "Invaild Environment. Terminating..." 라는 에러를 반환했군요!

```
func (t *Module) OnStart() error {
	Log.Info.Printf("ChatEcho: Started")
	return nil
}
```

이 코드는 모듈이 시작될 때 호출되는 함수입니다.
모든 모듈들에 대해 initialize가 끝나고, 런타임에서 모든 준비가 끝난 상태에서 `OnStart()`를 호출하여 런타임이 이제 정상적으로 가동 될 것임을 알립니다.

```
func (t *Module) GetInfo() *IChatanium.ModuleInfo {
	 return &IChatanium.ModuleInfo{
		Name:        t.Name,
		Description: t.Description,
		Version:     t.Version,
		Tags:        t.Tags,
		Commands:    t.Commands,
	}
}
```

이 코드는 런타임이 모듈을 식별하기 위해 호출하는 함수입니다.
이 함수는 런타임에서 모듈에 대해 정보가 필요할 때마다 호출하게 됩니다.
이는 런타임에서 기록을 위해 로그를 남기거나, 명령어에 대한 충돌이 발생하지 않도록 미리 확인하는 등의 충돌 방지를 위해 호출 될 수 있으니 잘 적는것이 중요합니다.