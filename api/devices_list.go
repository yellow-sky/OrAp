package api

import (
	nm "github.com/Wifx/gonetworkmanager/v2"
	nm_dev_man "github.com/yellow-sky/orap/nm_device_manager"
	"net/http"
)

// handleDevicesList godoc
// @Summary List of network devices
// @Description Get list of devices with short info: id, name, description and connection state.
// @Tags devices_info
// @Accept  json
// @Produce  json
// @Success 200 {object} CommonResponse{data=[]nm_device_manager.DeviceShortInfo} "Common JSON response with list of devices"
// @Failure 401 {object} CommonResponse{} "Unauthorized error"
// @Failure 500 {object} CommonResponse{} "Unhandled server error"
// @Security BasicAuth
// @Router /devices [get]
func (s ApiService) handleDevicesList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: change to const or use as request filter
		deviceTypeFilter := []nm.NmDeviceType{
			nm.NmDeviceTypeEthernet,
			nm.NmDeviceTypeEthernet,
			nm.NmDeviceTypeWifi,
			nm.NmDeviceTypeBt,
			nm.NmDeviceTypeWimax,
			nm.NmDeviceTypeModem,
			nm.NmDeviceTypeInfiniband,
			nm.NmDeviceTypeAdsl,
			nm.NmDeviceTypeWifiP2p,
		}

		devManager, err := nm_dev_man.NewNmDeviceManager()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on init device manager: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}

		devices, err := devManager.GetFilteredDevices(deviceTypeFilter)
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get devices: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}

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
