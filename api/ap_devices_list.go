package api

import (
	"github.com/yellow-sky/orap/ap_manager"
	nm_dev_man "github.com/yellow-sky/orap/nm_device_manager"
	"net/http"
)

// handleApCompatibleDevicesList godoc
// @Summary List of network devices for AP
// @Description Get list of devices with short info: id, name, description and connection state.
// @Tags ap_info
// @Accept  json
// @Produce  json
// @Success 200 {object} CommonResponse{data=[]nm_device_manager.DeviceShortInfo} "Common JSON response with list of devices"
// @Failure 401 {object} CommonResponse{} "Unauthorized error"
// @Failure 500 {object} CommonResponse{} "Unhandled server error"
// @Security BasicAuth
// @Router /ap/compatible_devices [get]
func (s ApiService) handleApCompatibleDevicesList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apManager, err := ap_manager.NewApManager()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on init AP manager: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}

		devices, err := apManager.GetCompatibleDevices()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get compatible devices: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}
		// TODO: make util func []Device -> []DeviceShortInfo???
		var devicesList []nm_dev_man.DeviceShortInfo
		for _, device := range devices {
			deviceInfo, err := nm_dev_man.NewDeviceShortInfo(device)
			if err != nil {
				resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
				s.writeCommonJsonResponse(w, resp)
				return
			}
			devicesList = append(devicesList, *deviceInfo)
		}
		resp := CommonResponse{Data: devicesList, Status: http.StatusOK}
		s.writeCommonJsonResponse(w, resp)
	}
}
