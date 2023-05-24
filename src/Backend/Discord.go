package Backend

import (
	IChatanium "antegral.net/chatanium/src/Runtime/Interface"
	"github.com/bwmarrin/discordgo"
)

// A chat backend adapter for Discord.
//
// An implementation of the IChatanium.Backend interface.
type Discord struct {
	Name string
	Api  *discordgo.Session
	key  string
}

func (t *Discord) Init(Info IChatanium.ModuleInfo) error {
	t.Name = "Discord"
	return nil
}

func (t *Discord) SetCredentials(Credentials ...string) error {
	for i, v := range Credentials {
		switch i {
		case 0:
			t.key = v
			break
		default:
			break
		}
	}
	return nil
}

func (t *Discord) Connect() error {
	api, err := discordgo.New("Bot " + t.key)
	if err != nil {
		return err
	}

	t.Api = api
	return nil
}

type Requirements struct {
	BACKEND_KEY string
}
