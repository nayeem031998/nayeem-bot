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

