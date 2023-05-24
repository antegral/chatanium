package IChatanium

// Fill in the information for the module. Make it recognizable to users.
//
// This is the interface that all modules must implement.
type ModuleInfo struct {
	Name        string
	Description string
	Version     string
	Tags        []string
	Commands    []string
}

// runtime executes the functions of the interface in context.
//
// This is the interface that all modules must implement.
type Module interface {
	OnInit() error
	OnStart() error
	GetInfo() *ModuleInfo
	GetBackend(Backend Backend) error
	OnMessage(Request string) ModuleChannel
}

type ModuleChannel chan string

var ModuleFuncMap = map[string]Module{}
