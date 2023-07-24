//go:build swagger_enabled

package swagger

import (
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/yellow-sky/orap/server_app"
)

var log = logrus.WithField("module", "swagger")

func InitSwaggerService(server *server_app.Server) {
	server.GetRouter().PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
