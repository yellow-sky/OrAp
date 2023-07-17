package api

import (
	nm "github.com/Wifx/gonetworkmanager/v2"
	"golang.org/x/exp/slices"
	"net/http"
	"path"
	"strings"
)

type DeviceShortInfo struct {
	ID        string `json:"id"`
	Interface string `json:"interface"`
	Type      string `json:"type"`
	Driver    string `json:"driver"`
	State     string `json:"state"`
}

// handleDevicesList godoc
// @Summary List of network devices
// @Description Get list of devices with short info: id, name, description and connection state.
// @Tags devices_info
// @Accept  json
// @Produce  json
// @Success 200 {object} CommonResponse{data=[]api.DeviceShortInfo} "Common JSON response with list of devices"
// @Failure 401 {object} CommonResponse{} "Unauthorized error"
// @Failure 500 {object} CommonResponse{} "Unhandled server error"
// @Security BasicAuth
// @Router /devices [get]
func (s ApiService) handleDevicesList(nmgr nm.NetworkManager) http.HandlerFunc {
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
		devices, err := nmgr.GetPropertyAllDevices()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get devices: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}

		var devicesList []DeviceShortInfo
		for _, device := range devices {
			deviceId := path.Base(string(device.GetPath()))
			deviceType, err := device.GetPropertyDeviceType()
			if err != nil {
				resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
				s.writeCommonJsonResponse(w, resp)
				return
			}
			if !slices.Contains(deviceTypeFilter, deviceType) {
				continue
			}
			deviceInterface, err := device.GetPropertyInterface()
			if err != nil {
				resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
				s.writeCommonJsonResponse(w, resp)
				return
			}
			deviceDriver, err := device.GetPropertyDriver()
			if err != nil {
				resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
				s.writeCommonJsonResponse(w, resp)
				return
			}
			state, err := device.GetPropertyState()
			if err != nil {
				resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
				s.writeCommonJsonResponse(w, resp)
				return
			}
			deviceInfo := DeviceShortInfo{
				ID:        deviceId,
				Interface: deviceInterface,
				Type:      strings.Replace(deviceType.String(), "NmDeviceType", "", 1),
				Driver:    deviceDriver,
				State:     strings.Replace(state.String(), "NmDeviceState", "", 1),
			}
			devicesList = append(devicesList, deviceInfo)
		}

		resp := CommonResponse{Data: devicesList, Status: http.StatusOK}
		s.writeCommonJsonResponse(w, resp)
	}
}
