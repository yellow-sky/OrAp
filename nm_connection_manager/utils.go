package nm_connection_manager

import (
	nm "github.com/Wifx/gonetworkmanager/v2"
	"path"
)

func GetConnectionsId(connection nm.Connection) string {
	return path.Base(string(connection.GetPath()))
}
