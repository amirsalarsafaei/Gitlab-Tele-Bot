package telegram

import (
	"context"

	tgbotapi "github.com/amirsalarsafaei/telegram-bot-api/v5"

	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients"
)

type Config struct {
	ChatID   int64 `mapstructure:"chat_id"`
	ThreadID int   `mapstructure:"thread_id"`

	Token string `mapstructure:"token"`
}

type notifierBot struct {
	*tgbotapi.BotAPI
	config *Config
}

func NewBot(config *Config) (clients.MessageBroker, error) {
	botAPI, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		return nil, err
	}

	return &notifierBot{
		BotAPI: botAPI,
		config: config,
	}, nil
}

func (b *notifierBot) SendMessage(_ context.Context, text string) error {
	message := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           b.config.ChatID,
			MessageThreadID:  b.config.ThreadID,
			ReplyToMessageID: 0,
		},
		Text:                  text,
		DisableWebPagePreview: false,
	}

	_, err := b.Send(message)
	return err
}
