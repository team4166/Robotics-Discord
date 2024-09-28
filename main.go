package main

import (
	"os"

	"github.com/joho/godotenv"
	bot "jadepy.dev/robotics-discord/bot"
)

func main() {
	godotenv.Load()
	bot.BotToken = os.Getenv("BOT_TOKEN")
	bot.Run()
}
