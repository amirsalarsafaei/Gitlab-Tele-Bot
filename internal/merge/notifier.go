package merge

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (h Handler) Notifier(resp http.ResponseWriter, req *http.Request) {
	var event Event
	err := json.NewDecoder(req.Body).Decode(&event)
	if err != nil {
		log.WithError(err).Error("merge notifier failed")
		resp.WriteHeader(400)
		return
	}

	if event.EventType != "merge_request" || event.Attributes.State != "merged" || event.Attributes.Action != "merge" {
		return
	}

	err = h.Handle(req.Context(), event.Project.Id, event.Attributes.Iid, event.Attributes.Title,
		event.Attributes.Description, event.Attributes.Url, event.Attributes.LastCommit.Author, event.Reviewers)
	if err != nil {
		log.WithError(err).Error("failed to send message")
		resp.WriteHeader(400)
		return
	}
}
