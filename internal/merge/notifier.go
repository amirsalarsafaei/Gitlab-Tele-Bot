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

	err = h.Handle(event)
	if err != nil {
		log.WithError(err).Error("merge notifier failed")
		resp.WriteHeader(400)
		return
	}
}
