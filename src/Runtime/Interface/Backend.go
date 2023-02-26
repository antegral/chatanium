package IChatanium

type Backend interface {
	Init(Info ModuleInfo) error
	SetCredentials(Credentials ...string) error
	Connect() error
}
