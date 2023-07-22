package nm_connection_manager

import (
	nm "github.com/Wifx/gonetworkmanager/v2"
	"github.com/mitchellh/mapstructure"
	"net"
	"strings"
)

type NmWirelessSettings struct {
	NmMacAddress        net.HardwareAddr `json:"-" mapstructure:"mac-address,omitempty"`
	MacAddress          string           `json:"mac-address,omitempty"`
	MacAddressBlacklist []string         `json:"mac-address-blacklist" mapstructure:"mac-address-blacklist"`
	Mode                string           `json:"mode" mapstructure:"mode"`
	SeenBSSIDs          []string         `json:"seen-bssids" mapstructure:"seen-bssids"`
	//SSID              []uint8  `mapstructure:"ssid"`
	SSID     string `json:"ssid" mapstructure:"ssid"`
	Security string `json:"security,omitempty" mapstructure:"security,omitempty"`
}

func (nm *NmWirelessSettings) FillJsonFields() {
	nm.MacAddress = strings.ToUpper(nm.NmMacAddress.String())
}

func NewNmWirelessSettings(raw nm.ConnectionSettings) (*NmWirelessSettings, error) {
	section, exists := raw["802-11-wireless"]
	if exists {
		connSettings := NmWirelessSettings{}
		err := mapstructure.WeakDecode(section, &connSettings)
		return &connSettings, err
	} else {
		return nil, nil
	}
}
