package nm_connection_manager

import (
	nm "github.com/Wifx/gonetworkmanager/v2"
	"github.com/mitchellh/mapstructure"
)

type NmConnectionSettings struct {
	raw    nm.ConnectionSettings
	Common NmCommonSettings `json:"connection" mapstructure:"connection"`
}

func NewNmConnectionSettings(raw nm.ConnectionSettings) (NmConnectionSettings, error) {
	conn := NmConnectionSettings{raw: raw}
	err := mapstructure.WeakDecode(raw, &conn)
	return conn, err
}
