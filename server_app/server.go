package server_app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/yellow-sky/orap/conf"
	"net/http"
)

var log = logrus.WithField("module", "http_server")

type Server struct {
	router       *mux.Router
	httpServer   http.Server
	useTls       bool
	certFilePath string
	keyFilePath  string
}

func (s *Server) Run() {
	var err error
	if s.useTls {
		err = s.httpServer.ListenAndServeTLS(s.certFilePath, s.keyFilePath)
	} else {
		err = s.httpServer.ListenAndServe()
	}
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

func (s *Server) Shutdown() {
	err := s.httpServer.Shutdown(context.Background())
	if err != nil {
		log.Errorln(err)
	}
}

func (s *Server) GetRouter() *mux.Router {
	return s.router
}

func NewServer(config conf.ApiConfig) *Server {
	server := Server{
		router: mux.NewRouter(),
		httpServer: http.Server{
			Addr: fmt.Sprintf(":%d", config.Port),
		},
	}

	if config.TlsKeyPath != "" && config.TlsCertPath != "" {
		log.Infof("Init https endpoint on port %d with keys: %s %s \n",
			config.Port,
			config.TlsCertPath,
			config.TlsKeyPath,
		)
		server.useTls = true
		server.certFilePath = config.TlsCertPath
		server.keyFilePath = config.TlsKeyPath
	} else {
		log.Infof("Init http endpoint on port: %d", config.Port)
		server.useTls = false
	}

	server.router.StrictSlash(true)
	//server.router.Use(handlers.CompressHandler)
	server.httpServer.Handler = server.router
	return &server
}
