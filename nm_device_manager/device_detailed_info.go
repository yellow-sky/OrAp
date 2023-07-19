package nm_device_manager

import (
	nm "github.com/Wifx/gonetworkmanager/v2"
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

func NewDeviceDetailedInfo(device nm.Device) (*DeviceDetailedInfo, error) {
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
	deviceDriverVersion, err := device.GetPropertyDriverVersion()
	if err != nil {
		return nil, err
	}
	fwVersion, err := device.GetPropertyFirmwareVersion()
	if err != nil {
		return nil, err
	}
	state, err := device.GetPropertyState()
	if err != nil {
		return nil, err
	}
	managed, err := device.GetPropertyManaged()
	if err != nil {
		return nil, err
	}
	autoConnect, err := device.GetPropertyAutoConnect()
	if err != nil {
		return nil, err
	}
	firmwareMissing, err := device.GetPropertyFirmwareMissing()
	if err != nil {
		return nil, err
	}
	nmPluginMissing, err := device.GetPropertyNmPluginMissing()
	if err != nil {
		return nil, err
	}
	isReal, err := device.GetPropertyReal()
	if err != nil {
		return nil, err
	}
	mtu, err := device.GetPropertyMtu()
	if err != nil {
		return nil, err
	}
	physicalPortId, err := device.GetPropertyPhysicalPortId()
	if err != nil {
		return nil, err
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
	return &deviceInfo, nil
}
