package telegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/cmd/config"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients"
)

type notifierBot struct {
	*tgbotapi.BotAPI
	chatID   string
	threadID string
}

func NewBot(config *config.Config) clients.MessageBroker {
	return &notifierBot{}
}

func (b *notifierBot) SendMessage(ctx context.Context, text string) error {
	return nil
}
