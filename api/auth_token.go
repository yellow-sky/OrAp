package api

import (
	"github.com/yellow-sky/orap/auth"
	"net/http"
)

// handleAuthToken godoc
// @Summary User Login method
// @Description Get JWT token for sent username and password.
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} CommonResponse{data=string} "Common JSON response with JWT as data"
// @Failure 401 {object} CommonResponse{} "Unauthorized error"
// @Failure 500 {object} CommonResponse{} "Unhandled server error"
// @Security BasicAuth
// @Router /auth/token [post]
func (s ApiService) handleAuthToken(authService *auth.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := authService.GetUserFromRequest(r)
		token, err := authService.CreateToken(user)
		resp := CommonResponse{Data: token, Status: http.StatusOK}
		if err != nil {
			resp = CommonResponse{Error: err.Error(), Status: http.StatusInternalServerError}
		}
		s.writeCommonJsonResponse(w, resp)
	}
}
