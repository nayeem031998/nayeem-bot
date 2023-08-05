package bot

import (
	"fmt"

	"nayeem_bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

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
	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "welcome nayeem eman bot")
	}
	if m.MentionEveryone {
		_, _ = s.ChannelMessageSend(m.ChannelID, "hii everyone")
	}
	if m.Content == "nayeem" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "hii nayeem how are you")
	}
	if m.Author.Email == "nayeemakhtar371@gmail.com" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "hii nayeem how are you")
	}
	
	
}