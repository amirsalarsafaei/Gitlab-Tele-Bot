package gitlab

import (
	"net/http"

	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/pkg/mux"
)

const SecretHeader = "X-Gitlab-Token"

func NewGitLabAuthMiddleWare(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if header := r.Header.Get(SecretHeader); header != secret {
				mux.UnAuthorized(w, r)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
