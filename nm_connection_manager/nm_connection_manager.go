package nm_connection_manager

import (
	"fmt"
	nm "github.com/Wifx/gonetworkmanager/v2"
	"github.com/yellow-sky/orap/common"
)

const NM_CONNECTION_TYPE_WIFI = "802-11-wireless"

type NmConnectionManager struct {
	sett nm.Settings
}

func NewNmConnectionManager() (NmConnectionManager, error) {
	settings, err := nm.NewSettings()
	return NmConnectionManager{sett: settings}, err
}

func (n NmConnectionManager) GetAllConnections() ([]NmConnection, error) {
	rawConnections, err := n.sett.ListConnections()
	if err != nil {
		return nil, fmt.Errorf("Error on get connections: %w", err)
	}
	var cons []NmConnection
	for _, rawConnection := range rawConnections {
		conn, err := NewNmConnection(rawConnection)
		if err != nil {
			return nil, fmt.Errorf("error on decode connections: %w", err)
		}
		cons = append(cons, *conn)
	}
	return cons, nil
}

func (n NmConnectionManager) GetWifiConnections() ([]NmConnection, error) {
	allConnections, err := n.GetAllConnections()
	if err != nil {
		return nil, err
	}
	wifiConnections := common.Filter(allConnections, func(conn NmConnection) bool {
		return conn.CommonSettings.Type == NM_CONNECTION_TYPE_WIFI
	})
	return wifiConnections, nil
}
