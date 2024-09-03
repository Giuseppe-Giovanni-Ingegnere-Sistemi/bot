package bot

import (
	"log"
	"github.com/slack-go/slack"
	"project-name/database"
)

type SlackBot struct {
	api *slack.Client
	db  *database.Database
}

func NewSlackBot(token string, db *database.Database) *SlackBot {
	return &SlackBot{
		api: slack.New(token),
		db:  db,
	}
}

func (bot *SlackBot) Start() {
	rtm := bot.api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			// Aquí puedes manejar los comandos y tareas del bot
			log.Printf("Mensaje recibido: %v\n", ev.Text)
		}
	}
}
