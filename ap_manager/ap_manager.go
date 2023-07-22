package ap_manager

import (
	"fmt"
	nm "github.com/Wifx/gonetworkmanager/v2"
	"github.com/yellow-sky/orap/common"
	nm_conn_man "github.com/yellow-sky/orap/nm_connection_manager"
	nm_dev_man "github.com/yellow-sky/orap/nm_device_manager"
)

const NM_WIRELESS_MODE_AP = "ap"

type ApManager struct {
	devManager  nm_dev_man.NmDeviceManager
	connManager nm_conn_man.NmConnectionManager
}

func NewApManager() (*ApManager, error) {
	devManager, err := nm_dev_man.NewNmDeviceManager()
	if err != nil {
		return nil, fmt.Errorf("error on init device manager: %w", err)
	}
	connManager, err := nm_conn_man.NewNmConnectionManager()
	if err != nil {
		return nil, fmt.Errorf("error on init connection manager: %w", err)
	}
	return &ApManager{
		devManager:  devManager,
		connManager: connManager,
	}, nil
}

func (a ApManager) GetCompatibleDevices() ([]nm.Device, error) {
	// TODO: add AP capabilities for filtered devices (is it possible with nm?)
	deviceTypeFilter := []nm.NmDeviceType{
		nm.NmDeviceTypeWifi,
		//nm.NmDeviceTypeWifiP2p,
	}
	devices, err := a.devManager.GetFilteredDevices(deviceTypeFilter)
	if err != nil {
		return nil, fmt.Errorf("error on get devices: %w", err)
	}
	return devices, nil
}

func (a ApManager) GetApConnections() ([]nm_conn_man.NmConnection, error) {
	wifiConnections, err := a.connManager.GetWifiConnections()
	if err != nil {
		return nil, err
	}
	apCons := common.Filter(wifiConnections, func(conn nm_conn_man.NmConnection) bool {
		return conn.WirelessSettings.Mode == NM_WIRELESS_MODE_AP
	})
	return apCons, nil
}

func (a ApManager) GetActiveAp() ([]ActiveAp, error) {
	devices, err := a.GetCompatibleDevices()
	if err != nil {
		return nil, fmt.Errorf("error on get compatible devices: %w", err)
	}
	var activeAps []ActiveAp
	for _, device := range devices {
		actConn, err := device.GetPropertyActiveConnection()
		if err != nil {
			return nil, fmt.Errorf("error on get device active connection: %w", err)
		}
		if actConn == nil {
			continue
		}
		rawConn, err := actConn.GetPropertyConnection()
		if err != nil {
			return nil, fmt.Errorf("error on get connection: %w", err)
		}
		if rawConn == nil {
			continue
		}

		conn, err := nm_conn_man.NewNmConnection(rawConn)
		if err != nil {
			return nil, fmt.Errorf("error on get connection: %w", err)
		}

		if conn.WirelessSettings.Mode == NM_WIRELESS_MODE_AP {
			activeAps = append(activeAps, NewActiveAp(rawConn, device))
		}
	}
	return activeAps, nil
}
