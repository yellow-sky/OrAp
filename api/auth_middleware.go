package api

import (
	"github.com/yellow-sky/orap/auth"
	"net/http"
)

func (s ApiService) createAuthMiddleware(authService *auth.AuthService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				user, err := authService.AuthenticateRequest(r)
				if err != nil {
					log.Debugf("auth error: %s\n", err.Error())
					resp := CommonResponse{Status: http.StatusUnauthorized, Error: "Error on auth: " + err.Error()}
					s.writeCommonJsonResponse(w, resp)
					return
				}
				r = authService.RequestWithUser(user, r)
				next.ServeHTTP(w, r)
			})
	}
}
