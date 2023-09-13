package clients

import "context"

type MessageBroker interface {
	SendMessage(ctx context.Context, text string) error
}
