package web

import (
	"embed"
	"github.com/koron/go-spafs"
	"github.com/sirupsen/logrus"
	"github.com/yellow-sky/orap/server_app"
	"io/fs"
	"net/http"
)

var log = logrus.WithField("module", "web")

//go:embed dist
var webDist embed.FS

func InitWebService(server *server_app.Server) {
	fsRoot, _ := fs.Sub(webDist, "dist")
	//server.GetRouter().PathPrefix("/web/").Handler(http.StripPrefix( "/web/",http.FileServer(http.FS(fsRoot))))
	server.GetRouter().PathPrefix("/web/").Handler(http.StripPrefix("/web/", spafs.FileServer(http.FS(fsRoot))))
}
