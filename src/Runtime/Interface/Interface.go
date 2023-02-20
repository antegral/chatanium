package Interface

type ModuleInfo struct {
	Name        string
	Description string
	Version     string
	Tags        []string
	Commands    []string
}

type SenderInfo struct {
	Id   string
	Name string
}

type IncomingMessage struct {
	Command string
	Body    string
}

type ModuleRequest struct {
	Sender  SenderInfo
	Message IncomingMessage
}

type ModuleResponse error

type ChataniumBackend struct{}

func (t *ChataniumBackend) Init(Info ModuleInfo) error

func (t *ChataniumBackend) MakeSession()
