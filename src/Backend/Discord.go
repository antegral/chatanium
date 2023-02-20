package Backend

import (
	"antegral.net/chatanium/src/Runtime/Interface"
	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	Api *discordgo.Session
	key string
}

func (t *Discord) Init(Info Interface.ModuleInfo) error {
	api, err := discordgo.New("Bot " + "authentication token")
	if err != nil {
		return err
	}

	t.Api = api
	return nil
}

func (t *Discord) MakeSession() {
}
