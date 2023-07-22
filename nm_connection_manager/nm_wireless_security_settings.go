package nm_connection_manager

import (
	nm "github.com/Wifx/gonetworkmanager/v2"
	"github.com/mitchellh/mapstructure"
)

type NmWirelessSecuritySettings struct {
	KeyMgmt string `json:"key-mgmt" mapstructure:"key-mgmt"`
}

func NewNmWirelessSecuritySettings(raw nm.ConnectionSettings) (*NmWirelessSecuritySettings, error) {
	section, exists := raw["802-11-wireless-security"]
	if exists {
		connSettings := NmWirelessSecuritySettings{}
		err := mapstructure.WeakDecode(section, &connSettings)
		return &connSettings, err
	} else {
		return nil, nil
	}
}
