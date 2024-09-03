package config

import (
	"os"
)

type Config struct {
	SlackToken    string
	DiscordToken  string
	MongoURI      string
	UseSlack      bool
	UseDiscord    bool
}

func LoadConfig() *Config {
	return &Config{
		SlackToken:   os.Getenv("SLACK_TOKEN"),
		DiscordToken: os.Getenv("DISCORD_TOKEN"),
		MongoURI:     os.Getenv("MONGO_URI"),
		UseSlack:     os.Getenv("USE_SLACK") == "true",
		UseDiscord:   os.Getenv("USE_DISCORD") == "true",
	}
}
