package api

import (
	"github.com/gorilla/handlers"
	"github.com/maltegrosse/go-modemmanager"
	"github.com/sirupsen/logrus"
	"github.com/yellow-sky/orap/auth"
	"github.com/yellow-sky/orap/server_app"
	"net/http"
)

// @title OrAP Service API
// @version 0.1
// @description API for network management - AP, ethernet, modems.
// @BasePath /api
// @securityDefinitions.basic BasicAuth
// @tag.name auth
// @tag.description Auth operations
// @tag.name common
// @tag.description Common operations
// @tag.name modems_info
// @tag.description Modems info operations

// TODO:
// // @securityDefinitions.apikey ApiKeyAuth
// // @name Authorization: Bearer
// // @in header

type ApiService struct{}

var log = logrus.WithField("module", "api")

func InitApiService(server *server_app.Server, authService *auth.AuthService, mmgr modemmanager.ModemManager) *ApiService {
	// Init ApiService
	apiService := ApiService{}

	// Create cors middleware
	// TODO: Move settings to config?
	corsMiddleware := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type", "content-length", "accept-encoding", "authorization", "x-csrf-token"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodOptions}),
		handlers.AllowCredentials(),
	)

	authMiddleware := apiService.createAuthMiddleware(authService)

	apiSubrouter := server.GetRouter().PathPrefix("/api").Subrouter()
	apiSubrouter.Use(corsMiddleware)
	apiSubrouter.Use(authMiddleware)

	// auth
	apiSubrouter.HandleFunc("/auth/token", apiService.handleAuthToken(authService)).Methods(http.MethodPost)

	// common
	//apiSubrouter.HandleFunc("/metadata", apiService.handleMetadata(webRtcSettings.IceServer))

	// modems
	apiSubrouter.HandleFunc("/modems/", apiService.handleModemsList(mmgr)).Methods(http.MethodGet)
	modemsSubrouter := apiSubrouter.PathPrefix("/modems/{modem_id}").Subrouter()

	modemsSubrouter.Use(apiService.createModemsMiddleware(mmgr))
	modemsSubrouter.HandleFunc("/", apiService.handleModemInfo()).Methods(http.MethodGet)

	return &apiService
}
