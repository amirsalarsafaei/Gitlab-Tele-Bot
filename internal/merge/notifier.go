package merge

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (h Handler) Notifier(resp http.ResponseWriter, req *http.Request) {
	var reqJson map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&reqJson)
	if err != nil {
		log.WithError(err).Error("merge notifier failed")
	}

	fmt.Println(reqJson)
}
