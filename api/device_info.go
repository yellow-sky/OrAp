package api

import (
	nm "github.com/Wifx/gonetworkmanager/v2"
	nm_dev_man "github.com/yellow-sky/orap/nm_device_manager"
	"net/http"
)

// handleDeviceInfo godoc
// @Summary Device detailed information
// @Description Get info about device: common info, state, etc.
// @Tags devices_info
// @Accept  json
// @Produce  json
// @Param device_id path string true "Device id from list of devices"
// @Success 200 {object} CommonResponse{data=nm_device_manager.DeviceDetailedInfo} "Common JSON response with modem detailed info as data"
// @Failure 401 {object} CommonResponse{} "Unauthorized error"
// @Failure 404 {object} CommonResponse{} "Device not found error"
// @Failure 500 {object} CommonResponse{} "Unhandled server error"
// @Security BasicAuth
// @Router /devices/{device_id} [get]
func (s ApiService) handleDeviceInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
		device := r.Context().Value("device").(nm.Device)

		deviceInfo, err := nm_dev_man.NewDeviceDetailedInfo(device)
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}

		resp := CommonResponse{Data: deviceInfo, Status: http.StatusOK}
		s.writeCommonJsonResponse(w, resp)
	}
}
