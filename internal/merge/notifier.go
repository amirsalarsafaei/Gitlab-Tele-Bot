package merge

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (h Handler) Notifier(resp http.ResponseWriter, req *http.Request) {
	var reqMap map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&reqMap)
	if err != nil {
		log.WithError(err).Error("merge notifier failed")
	}

	reqJSON, err := json.MarshalIndent(reqMap, "", "\t")

	fmt.Println(string(reqJSON))
}
