package api

import (
	"github.com/Wifx/gonetworkmanager/v2"
	"net/http"
	"strings"
)

type DeviceDetailedInfo struct {
	ID              string `json:"id"`
	Interface       string `json:"interface"`
	Type            string `json:"type"`
	Driver          string `json:"driver"`
	DriverVersion   string `json:"driver_version"`
	FwVersion       string `json:"fw_version"`
	State           string `json:"state"`
	Managed         bool   `json:"managed"`
	AutoConnect     bool   `json:"auto_connect"`
	FirmwareMissing bool   `json:"firmware_missing"`
	NmPluginMissing bool   `json:"nm_plugin_missing"`
	Real            bool   `json:"real"`
	Mtu             uint32 `json:"mtu"`
	PhysicalPortId  string `json:"physical_port_id"`
}

// handleModemInfo godoc
// @Summary Device detailed information
// @Description Get info about device: common info, state, etc.
// @Tags devices_info
// @Accept  json
// @Produce  json
// @Param device_id path string true "Device id from list of devices"
// @Success 200 {object} CommonResponse{data=string} "Common JSON response with modem detailed info as data"
// @Failure 401 {object} CommonResponse{} "Unauthorized error"
// @Failure 404 {object} CommonResponse{} "Device not found error"
// @Failure 500 {object} CommonResponse{} "Unhandled server error"
// @Security BasicAuth
// @Router /devices/{device_id} [get]
func (s ApiService) handleDeviceInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
		device := r.Context().Value("device").(gonetworkmanager.Device)
		deviceId := r.Context().Value("device_id").(string)
		deviceType, err := device.GetPropertyDeviceType()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
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
		deviceDriverVersion, err := device.GetPropertyDriverVersion()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}
		fwVersion, err := device.GetPropertyFirmwareVersion()
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
		managed, err := device.GetPropertyManaged()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}
		autoConnect, err := device.GetPropertyAutoConnect()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}
		firmwareMissing, err := device.GetPropertyFirmwareMissing()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}
		nmPluginMissing, err := device.GetPropertyNmPluginMissing()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}
		isReal, err := device.GetPropertyReal()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}
		mtu, err := device.GetPropertyMtu()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}
		physicalPortId, err := device.GetPropertyPhysicalPortId()
		if err != nil {
			resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device info: " + err.Error()}
			s.writeCommonJsonResponse(w, resp)
			return
		}

		deviceInfo := DeviceDetailedInfo{
			ID:              deviceId,
			Interface:       deviceInterface,
			Type:            strings.Replace(deviceType.String(), "NmDeviceType", "", 1),
			Driver:          deviceDriver,
			DriverVersion:   deviceDriverVersion,
			FwVersion:       fwVersion,
			State:           strings.Replace(state.String(), "NmDeviceState", "", 1),
			Managed:         managed,
			AutoConnect:     autoConnect,
			FirmwareMissing: firmwareMissing,
			NmPluginMissing: nmPluginMissing,
			Real:            isReal,
			Mtu:             mtu,
			PhysicalPortId:  physicalPortId,
		}
		resp := CommonResponse{Data: deviceInfo, Status: http.StatusOK}
		s.writeCommonJsonResponse(w, resp)
	}
}
