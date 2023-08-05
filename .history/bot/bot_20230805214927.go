package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"nayeem_bot/config"
)

var (
	BotID string
	goBot *discordgo.Session
)

// Start starts the bot
func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")
}

// messageHandler handles the messages
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
