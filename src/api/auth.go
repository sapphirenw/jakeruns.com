package api

import (
	"net/http"

	"github.com/sapphirenw/jakeruns.com/src/api/response"
	"github.com/sapphirenw/jakeruns.com/src/lib"
	"github.com/sapphirenw/jakeruns.com/src/logger"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		if apiKey != lib.API_KEY {
			logger.Error.Printf("Invalid api key: '%s'\n", apiKey)
			response.WriteStr(w, http.StatusForbidden, "Not Allowed.")
			return
		}
		next.ServeHTTP(w, r)
	})
}
