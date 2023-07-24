//go:build !swagger_enabled

package swagger

import (
	"github.com/yellow-sky/orap/server_app"
)

func InitSwaggerService(server *server_app.Server) {
	//TODO: add empty page with disabled feature message?
}
