package nm_connection_manager

type NmCommonSettings struct {
	Autoconnect string   `json:"autoconnect" mapstructure:"autoconnect"`
	Id          string   `json:"id" mapstructure:"id"`
	Timestamp   uint64   `json:"timestamp" mapstructure:"timestamp"`
	Type        string   `json:"type" mapstructure:"type"`
	UUID        string   `json:"uuid" mapstructure:"uuid"`
	Permissions []string `json:"permissions" mapstructure:"permissions"`
}
