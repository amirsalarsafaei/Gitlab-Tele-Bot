package merge

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/pkg/quotes"
)

type Handler struct {
	brokers []clients.MessageBroker

	config MergeConfig
}

func NewHandler(brokers []clients.MessageBroker, config MergeConfig) Handler {
	return Handler{
		brokers: brokers,
		config:  config,
	}
}

func (h Handler) Handle(ctx context.Context, event Event) error {
	if event.EventType != "merge_request" || event.Attributes.State != "merged" || event.Attributes.Action != "merge" {
		return nil
	}

	var diff string

	var totalRemoved, totalAdded int
	if h.config.Diff.Enabled {
		diffs, err := h.getMergeRequestDiff(ctx, event.Project.Id, event.Attributes.Iid)
		if err != nil {
			log.WithError(err).Error("could not get diffs")
		}
		for i, diffI := range diffs {

			removed, added := diffI.GetDiffChanges()
			totalRemoved += removed
			totalAdded += added
			switch {
			case diffI.DeletedFile:
				diff += fmt.Sprintf("❌removed %s, %d lines", diffI.NewPath, removed)
			case diffI.NewFile:
				diff += fmt.Sprintf("🆕new file %s, %d lines", diffI.NewPath, added)
			case diffI.RenamedFile:
				diff += fmt.Sprintf("📛renamed from %s and ", diffI.OldPath)
			default:
				diff += fmt.Sprintf("✍️edited file %s, %d line removed, %d line added", diffI.NewPath, removed, added)
			}

			if i != len(diffs)-1 {
				diff += "\n\n"
			}
		}
	}

	quote := quotes.GetQuote()
	message := fmt.Sprintf(
		messageTemplate,
		event.Attributes.Title,
		event.Attributes.LastCommit.Author.Name,
		getReviewersText(event),
		totalRemoved, totalAdded,
		sanitizeDescription(event.Attributes.Description),
		quote.QuoteText, quote.QuoteAuthor,
		diff,
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
