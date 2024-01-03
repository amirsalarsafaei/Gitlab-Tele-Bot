package merge

import (
	"context"
	"errors"
	"fmt"

	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/pkg/quotes"
)

type Handler struct {
	Brokers []clients.MessageBroker
}

func (h Handler) Handle(ctx context.Context, event Event) error {
	if event.EventType != "merge_request" || event.Attributes.State != "merged" || event.Attributes.Action != "merge" {
		return nil
	}

	quote := quotes.GetQuote()
	message := fmt.Sprintf(
		messageTemplate,
		event.Attributes.Title,
		event.Attributes.LastCommit.Author.Name,
		getReviewersText(event),
		sanitizeDescription(event.Attributes.Description),
		quote.QuoteText, quote.QuoteAuthor,
		event.Attributes.Url,
	)

	var errs []error
	for _, broker := range h.Brokers {
		err := broker.SendMessage(ctx, message)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}
