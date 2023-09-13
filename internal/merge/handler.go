package merge

import (
	"fmt"

	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/pkg/quotes"
)

type Handler struct {
	Brokers []clients.MessageBroker
}

func (h Handler) Handle(event Event) error {
	quote := quotes.GetQuote()
	message := fmt.Sprintf(
		messageTemplate,
		event.Attributes.Title,
		event.Attributes.LastCommit.Author.Name,
		event.Attributes.Url,
		quote.QuoteText, quote.QuoteAuthor,
	)

	fmt.Println(message)
	return nil
}
