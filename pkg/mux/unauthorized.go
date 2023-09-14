package mux

import "net/http"

func UnAuthorized(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, 401, "UnAuthorized")
}
