package actions

import (
	"github.com/miiniper/loges"
	bot2 "github.com/miiniper/tgmsg_bot/bot"
	"go.uber.org/zap"
)

func NewLifeBot(botName string, botToken string) *bot2.BotApi {
	bot, err := bot2.NewBotApi(botToken, botName)
	if err != nil {
		loges.Loges.Error("get bot error", zap.Error(err))
	}
	return bot
}
