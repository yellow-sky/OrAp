package nm_connection_manager

import (
	"fmt"
	nm "github.com/Wifx/gonetworkmanager/v2"
)

type NmConnection struct {
	raw                      nm.Connection
	ID                       string                      `json:"id"`
	CommonSettings           NmCommonSettings            `json:"common_settings"`
	WirelessSettings         *NmWirelessSettings         `json:"wireless_settings,omitempty"`
	WirelessSecuritySettings *NmWirelessSecuritySettings `json:"wireless_security_settings,omitempty"`
}

func NewNmConnection(rawConnection nm.Connection) (*NmConnection, error) {
	rawConnSettings, err := rawConnection.GetSettings()
	if err != nil {
		return nil, fmt.Errorf("error on get connection settings: %w", err)
	}
	commonSettings, err := NewNmCommonSettings(rawConnSettings)
	if err != nil {
		return nil, fmt.Errorf("error on decode connection settings: %w", err)
	}
	wirelessSettings, err := NewNmWirelessSettings(rawConnSettings)
	if err != nil {
		return nil, fmt.Errorf("error on decode wireless connection settings: %w", err)
	}
	wirelessSecSettings, err := NewNmWirelessSecuritySettings(rawConnSettings)
	if err != nil {
		return nil, fmt.Errorf("error on decode wireless security connection settings: %w", err)
	}
	conn := NmConnection{
		raw:                      rawConnection,
		ID:                       GetConnectionsId(rawConnection),
		CommonSettings:           *commonSettings,
		WirelessSettings:         wirelessSettings,
		WirelessSecuritySettings: wirelessSecSettings,
	}
	return &conn, err
}
