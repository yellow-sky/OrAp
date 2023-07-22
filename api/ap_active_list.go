package api

import (
	"github.com/yellow-sky/orap/ap_manager"
	"net/http"
)

// handleApActiveList godoc
// @Summary List of active AP
// @Description Get list of active AP with info about Connection-Device
// @Tags ap_info
// @Accept  json
// @Produce  json
// @Success 200 {object} CommonResponse{data=[]ap_manager.ActiveAp} "Common JSON response with list of active AP"
// @Failure 401 {object} CommonResponse{} "Unauthorized error"
// @Failure 500 {object} CommonResponse{} "Unhandled server error"
// @Security BasicAuth
// @Router /ap/active [get]
func (s ApiService) handleApActiveList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apManager, err := ap_manager.NewApManager()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on init AP manager: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}

		activeAp, err := apManager.GetActiveAp()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get active AP: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}

		resp := CommonResponse{Data: activeAp, Status: http.StatusOK}
		s.writeCommonJsonResponse(w, resp)
	}
}
