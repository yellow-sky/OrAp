package nm_connection_manager

import (
	"fmt"
	nm "github.com/Wifx/gonetworkmanager/v2"
)

const NM_CONNECTION_TYPE_WIFI = "802-11-wireless"
const NM_WIRELESS_MODE_AP = "ap"

type NmConnectionManager struct {
	sett nm.Settings
}

func NewNmConnectionManager() (NmConnectionManager, error) {
	settings, err := nm.NewSettings()
	return NmConnectionManager{sett: settings}, err
}

func (n NmConnectionManager) GetAllConnections() ([]NmConnectionSettings, error) {
	rawConnections, err := n.sett.ListConnections()
	if err != nil {
		return nil, fmt.Errorf("Error on get connections: %w", err)
	}
	var cons []NmConnectionSettings
	for _, rawConnection := range rawConnections {
		rawConnSettings, err := rawConnection.GetSettings()
		if err != nil {
			return nil, fmt.Errorf("Error on get connection settings: %w", err)
		}
		connSettings, err := NewNmConnectionSettings(rawConnSettings)
		if err != nil {
			return nil, fmt.Errorf("Error on decode connection settings: %w", err)
		}
		cons = append(cons, connSettings)
	}
	return cons, nil
}

func (n NmConnectionManager) GetWifiConnections() ([]NmWifiConnectionSettings, error) {
	allConnections, err := n.GetAllConnections()
	if err != nil {
		return nil, err
	}
	var wifiConnections []NmWifiConnectionSettings
	for _, connection := range allConnections {
		if connection.Common.Type == NM_CONNECTION_TYPE_WIFI {
			wifiConnection, err := NewWifiConnectionSettings(connection)
			if err != nil {
				return nil, err
			}
			wifiConnections = append(wifiConnections, wifiConnection)
		}
	}
	return wifiConnections, nil
}

func (n NmConnectionManager) GetApConnections() ([]NmWifiConnectionSettings, error) {
	wifiConnections, err := n.GetWifiConnections()
	if err != nil {
		return nil, err
	}
	var apCons []NmWifiConnectionSettings
	for _, connection := range wifiConnections {
		if connection.Wireless.Mode == NM_WIRELESS_MODE_AP {
			apCons = append(apCons, connection)
		}
	}
	return apCons, nil
}
