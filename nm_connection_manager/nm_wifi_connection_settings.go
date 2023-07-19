package nm_connection_manager

import "github.com/mitchellh/mapstructure"

type NmWifiConnectionSettings struct {
	NmConnectionSettings
	Wireless         NmWirelessSettings         `json:"802-11-wireless" mapstructure:"802-11-wireless"`
	WirelessSecurity NmWirelessSecuritySettings `json:"802-11-wireless-security,omitempty" mapstructure:"802-11-wireless-security,omitempty"`
}

func NewWifiConnectionSettings(connSet NmConnectionSettings) (NmWifiConnectionSettings, error) {
	conn := NmWifiConnectionSettings{NmConnectionSettings: connSet}
	err := mapstructure.WeakDecode(conn.raw, &conn)
	conn.Wireless.FillJsonFields()
	return conn, err
}
