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

type Module struct {
	Info ModuleInfo
}

func (t *Module) OnInit() error

func (t *Module) OnStart() error

func (t *Module) GetInfo() *ModuleInfo

func (t *Module) GetBackend(Backend Backend) error

func (t *Module) OnMessage(Request string) ModuleChannel

type Backend struct{}

func (t *Backend) Init(Info ModuleInfo) error

func (t *Backend) SetCredentials(Credentials ...string) error

func (t *Backend) Connect() error
