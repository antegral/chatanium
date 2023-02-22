package IChatanium

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

type ModuleChannel chan string

type Module interface {
	OnInit() error
	OnStart() error
	GetInfo() *ModuleInfo
	GetBackend(Backend Backend) error
	OnMessage(Request string) ModuleChannel
}

type Backend interface {
	Init(Info ModuleInfo) error
	SetCredentials(Credentials ...string) error
	Connect() error
}
