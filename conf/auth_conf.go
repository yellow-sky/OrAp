package conf

import (
	"github.com/hashicorp/go-uuid"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthConfig struct {
	Admin     UserCredentials `json:"admin"`
	JwtSecret string          `json:"jwt_secret"`
}

const AuthDefaultConfKey = "auth"

func AuthGetConfigDefaults() AuthConfig {
	newUuid, _ := uuid.GenerateUUID()
	defaultConfig := AuthConfig{
		Admin: UserCredentials{
			Username: "admin",
			Password: "orap",
		},
		JwtSecret: newUuid,
	}
	return defaultConfig
}
