package nm_device_manager

import (
	nm "github.com/Wifx/gonetworkmanager/v2"
	"path"
)

func GetDeviceId(device nm.Device) string {
	return path.Base(string(device.GetPath()))
}
