package conf

type ApiConfig struct {
	Port        int
	TlsCertPath string `mapstructure:"tls_cert_path" json:"tls_cert_path"`
	TlsKeyPath  string `mapstructure:"tls_key_path" json:"tls_key_path"`
}

const ApiDefaultConfKey = "api"

func ApiGetConfigDefaults() ApiConfig {
	defaultConfig := ApiConfig{
		Port:        8080,
		TlsCertPath: "",
		TlsKeyPath:  "",
	}
	return defaultConfig
}
