package IChatanium

type ModuleInfo struct {
	Name        string
	Description string
	Version     string
	Tags        []string
	Commands    []string
}

type Module interface {
	OnInit() error
	OnStart() error
	GetInfo() *ModuleInfo
	GetBackend(Backend Backend) error
	OnMessage(Request string) ModuleChannel
}

type ModuleChannel chan string

var ModuleFuncMap = map[string]Module{}
