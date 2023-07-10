package api

import (
	"net/http"
)

type ModemDetailedInfo struct {
	// TODO
}

// handleModemInfo godoc
// @Summary Modem detailed information
// @Description Get info about modem: common info, state, etc.
// @Tags modems_info
// @Accept  json
// @Produce  json
// @Param modem_id path string true "Modem id from list of modems"
// @Success 200 {object} CommonResponse{data=string} "Common JSON response with modem detailed info as data"
// @Failure 401 {object} CommonResponse{} "Unauthorized error"
// @Failure 404 {object} CommonResponse{} "Modem not found error"
// @Failure 500 {object} CommonResponse{} "Unhandled server error"
// @Security BasicAuth
// @Router /modems/{modem_id} [get]
func (s ApiService) handleModemInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
		//modem := r.Context().Value("modem").(modemmanager.Modem)
		detailedInfo := ModemDetailedInfo{}
		resp := CommonResponse{Data: detailedInfo, Status: http.StatusOK}
		s.writeCommonJsonResponse(w, resp)
	}
}
