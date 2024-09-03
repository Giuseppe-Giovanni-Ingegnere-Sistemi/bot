package main

import (
	"log"

	"project-name/bot"
	"project-name/config"
	"project-name/database"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Conectar a la base de datos
	db, err := database.Connect(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Error conectando a MongoDB: %v", err)
	}
	defer db.Disconnect()

	// Inicializar bots
	if cfg.UseSlack {
		slackBot := bot.NewSlackBot(cfg.SlackToken, db)
		go slackBot.Start()
	}

	if cfg.UseDiscord {
		discordBot := bot.NewDiscordBot(cfg.DiscordToken, db)
		go discordBot.Start()
	}

	// Mantener el bot activo
	select {}
}
