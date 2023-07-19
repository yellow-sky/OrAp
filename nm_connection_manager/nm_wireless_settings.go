package nm_connection_manager

import (
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
