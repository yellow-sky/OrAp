package nm_device_manager

import (
	nm "github.com/Wifx/gonetworkmanager/v2"
	"golang.org/x/exp/slices"
)

type NmDeviceManager struct {
	nmgr nm.NetworkManager
}

func NewNmDeviceManager() (NmDeviceManager, error) {
	nmgr, err := nm.NewNetworkManager()
	return NmDeviceManager{nmgr: nmgr}, err
}

func (n NmDeviceManager) GetAllDevices() ([]nm.Device, error) {
	devices, err := n.nmgr.GetPropertyAllDevices()
	return devices, err
}

func (n NmDeviceManager) GetFilteredDevices(deviceTypeFilter []nm.NmDeviceType) ([]nm.Device, error) {
	devices, err := n.GetAllDevices()
	if err != nil {
		return nil, err
	}
	var filteredDevices []nm.Device
	for _, device := range devices {
		deviceType, err := device.GetPropertyDeviceType()
		if err != nil {
			return nil, err
		}
		if slices.Contains(deviceTypeFilter, deviceType) {
			filteredDevices = append(filteredDevices, device)
		}
	}
	return filteredDevices, err
}

func (n NmDeviceManager) GetDeviceById(id string) (nm.Device, error) {
	devices, err := n.GetAllDevices()
	if err != nil {
		return nil, err
	}
	var searchedDev nm.Device
	for _, device := range devices {
		if id == GetDeviceId(device) {
			searchedDev = device
			break
		}
	}
	return searchedDev, err
}
