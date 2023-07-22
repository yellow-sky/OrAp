package nm_connection_manager

import (
	nm "github.com/Wifx/gonetworkmanager/v2"
	"github.com/mitchellh/mapstructure"
)

type NmCommonSettings struct {
	Autoconnect string   `json:"autoconnect" mapstructure:"autoconnect"`
	Id          string   `json:"id" mapstructure:"id"`
	Timestamp   uint64   `json:"timestamp" mapstructure:"timestamp"`
	Type        string   `json:"type" mapstructure:"type"`
	UUID        string   `json:"uuid" mapstructure:"uuid"`
	Permissions []string `json:"permissions" mapstructure:"permissions"`
}

func NewNmCommonSettings(raw nm.ConnectionSettings) (*NmCommonSettings, error) {
	section, exists := raw["connection"]
	if exists {
		connSettings := NmCommonSettings{}
		err := mapstructure.WeakDecode(section, &connSettings)
		return &connSettings, err
	} else {
		return nil, nil
	}
}
