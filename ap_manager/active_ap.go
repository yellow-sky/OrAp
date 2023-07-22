package ap_manager

import (
	nm "github.com/Wifx/gonetworkmanager/v2"
	nm_con_man "github.com/yellow-sky/orap/nm_connection_manager"
	nm_dev_man "github.com/yellow-sky/orap/nm_device_manager"
)

type ActiveAp struct {
	ConnectionID string `json:"connection_id"`
	DeviceID     string `json:"device_id"`
}

func NewActiveAp(connection nm.Connection, device nm.Device) ActiveAp {
	return ActiveAp{
		ConnectionID: nm_con_man.GetConnectionsId(connection),
		DeviceID:     nm_dev_man.GetDeviceId(device),
	}
}
