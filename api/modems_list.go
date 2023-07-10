package api

import (
	"github.com/kr/pretty"
	"github.com/maltegrosse/go-modemmanager"
	"net/http"
	"path"
)

type ModemShortInfo struct {
	ID string `json:"id"`
	// TODO
}

// handleModemsList godoc
// @Summary List of modems
// @Description Get list of modems with short info: id, name, description and connection state.
// @Tags modems_info
// @Accept  json
// @Produce  json
// @Success 200 {object} CommonResponse{data=[]api.ModemShortInfo} "Common JSON response with list of modems"
// @Failure 401 {object} CommonResponse{} "Unauthorized error"
// @Failure 500 {object} CommonResponse{} "Unhandled server error"
// @Security BasicAuth
// @Router /modems [get]
func (s ApiService) handleModemsList(mmgr modemmanager.ModemManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
		//err = mmgr.ScanDevices()
		//if err != nil {
		//	resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on modem scan: " + err.Error()}
		//	s.writeCommonJsonResponse(w, resp)
		//	return
		//}

		sysModems, err := mmgr.GetModems()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get modems : " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}

		var modemsList []ModemShortInfo
		for _, sysModem := range sysModems {
			sysModemId := path.Base(string(sysModem.GetObjectPath()))
			sysModemInfo := ModemShortInfo{
				ID: sysModemId,
			}
			pretty.Println(sysModem.GetObjectPath())
			pretty.Println(sysModem.GetDeviceIdentifier())
			pretty.Println(sysModem.GetDevice())
			pretty.Println(sysModem.GetEquipmentIdentifier())
			pretty.Println(sysModem.GetModel())
			pretty.Println(sysModem.GetOwnNumbers())
			pretty.Println(sysModem.GetManufacturer())
			pretty.Println(sysModem.GetManufacturer())
			pretty.Println(sysModem.GetBearers())
			pretty.Println(sysModem.Get3gpp())

			modemsList = append(modemsList, sysModemInfo)
		}
		resp := CommonResponse{Data: modemsList, Status: http.StatusOK}
		s.writeCommonJsonResponse(w, resp)
	}
}
