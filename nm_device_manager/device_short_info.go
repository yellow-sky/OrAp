package nm_device_manager

import (
	nm "github.com/Wifx/gonetworkmanager/v2"
	"strings"
)

type DeviceShortInfo struct {
	ID        string `json:"id"`
	Interface string `json:"interface"`
	Type      string `json:"type"`
	Driver    string `json:"driver"`
	State     string `json:"state"`
}

func NewDeviceShortInfo(device nm.Device) (*DeviceShortInfo, error) {
	deviceId := GetDeviceId(device)
	deviceType, err := device.GetPropertyDeviceType()
	if err != nil {
		return nil, err
	}
	deviceInterface, err := device.GetPropertyInterface()
	if err != nil {
		return nil, err
	}
	deviceDriver, err := device.GetPropertyDriver()
	if err != nil {
		return nil, err
	}
	state, err := device.GetPropertyState()
	if err != nil {
		return nil, err
	}
	deviceInfo := DeviceShortInfo{
		ID:        deviceId,
		Interface: deviceInterface,
		Type:      strings.Replace(deviceType.String(), "NmDeviceType", "", 1),
		Driver:    deviceDriver,
		State:     strings.Replace(state.String(), "NmDeviceState", "", 1),
	}
	return &deviceInfo, nil
}
