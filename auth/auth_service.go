package auth

import (
	"context"
	"fmt"
	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/jwt"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"
	"github.com/sirupsen/logrus"
	"github.com/yellow-sky/orap/conf"
	"net/http"
	"time"
)

var log = logrus.WithField("module", "auth")

type AuthService struct {
	jwtSecretsKeeper jwt.SecretsKeeper
	cache            auth.Cache
	strategy         union.Union
	adminCredentials conf.UserCredentials
}

func NewAuthService(config conf.AuthConfig) AuthService {
	authService := AuthService{
		adminCredentials: config.Admin,
	}

	// Init keeper
	authService.jwtSecretsKeeper = jwt.StaticSecret{
		ID:        "orap-id",
		Secret:    []byte(config.JwtSecret),
		Algorithm: jwt.HS256,
	}

	// Init cache
	cache := libcache.FIFO.New(500)
	cache.SetTTL(time.Minute * 5)

	// Init strategy
	basicStrategy := basic.NewCached(authService.authenticateUser, cache)
	jwtStrategy := jwt.New(cache, authService.jwtSecretsKeeper)

	authService.strategy = union.New(jwtStrategy, basicStrategy)

	return authService
}

func (a AuthService) CreateToken(user auth.Info) (string, error) {
	if user == nil {
		return "", fmt.Errorf("non authorised user")
	}
	return jwt.IssueAccessToken(user, a.jwtSecretsKeeper)
}

func (a AuthService) AuthenticateRequest(r *http.Request) (auth.Info, error) {
	_, user, err := a.strategy.AuthenticateRequest(r)
	return user, err
}

func (a AuthService) GetUserFromRequest(r *http.Request) auth.Info {
	return auth.User(r)
}

func (a AuthService) RequestWithUser(user auth.Info, r *http.Request) *http.Request {
	return auth.RequestWithUser(user, r)
}

func (a AuthService) authenticateUser(ctx context.Context, r *http.Request, userName, password string) (auth.Info, error) {
	// Simple static admin user auth
	if userName == a.adminCredentials.Username && password == a.adminCredentials.Password {
		return &User{
			DefaultUser: auth.DefaultUser{
				Name:       a.adminCredentials.Username,
				ID:         "0",
				Groups:     nil,
				Extensions: nil,
			},
		}, nil
	}
	return nil, fmt.Errorf("invalid credentials")
}
