package stdout

import (
	"context"
	"fmt"

	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients"
)

type PrintBot struct{}

func NewPrintBot() clients.MessageBroker {
	return &PrintBot{}
}

func (b *PrintBot) SendMessage(_ context.Context, text string) error {
	fmt.Println(text)

	return nil
}
