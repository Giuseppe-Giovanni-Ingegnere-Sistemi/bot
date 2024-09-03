package bot

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"project-name/database"
)

type DiscordBot struct {
	session *discordgo.Session
	db      *database.Database
}

func NewDiscordBot(token string, db *database.Database) *DiscordBot {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creando el bot de Discord: %v", err)
	}
	return &DiscordBot{
		session: dg,
		db:      db,
	}
}

func (bot *DiscordBot) Start() {
	bot.session.AddMessageCreateEventHandler(bot.messageCreate)
	err := bot.session.Open()
	if err != nil {
		log.Fatalf("Error abriendo la sesión de Discord: %v", err)
	}
}

func (bot *DiscordBot) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Manejar comandos personalizados
	log.Printf("Mensaje recibido: %v\n", m.Content)
}
