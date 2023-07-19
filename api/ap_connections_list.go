package api

import (
	nm_con_man "github.com/yellow-sky/orap/nm_connection_manager"
	"net/http"
)

// handleApConnectionsList godoc
// @Summary List of AP connections
// @Description Get list of AP connection settings with short info: id, name, description and connection state.
// @Tags ap_info
// @Accept  json
// @Produce  json
// @Success 200 {object} CommonResponse{data=[]nm_connection_manager.NmWifiConnectionSettings} "Common JSON response with list of AP connections"
// @Failure 401 {object} CommonResponse{} "Unauthorized error"
// @Failure 500 {object} CommonResponse{} "Unhandled server error"
// @Security BasicAuth
// @Router /ap/connections [get]
func (s ApiService) handleApConnectionsList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		connManager, err := nm_con_man.NewNmConnectionManager()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on init connection manager: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}
		apCons, err := connManager.GetApConnections()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get AP connections: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}

		resp := CommonResponse{Data: apCons, Status: http.StatusOK}
		s.writeCommonJsonResponse(w, resp)
	}
}
