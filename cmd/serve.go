package cmd

import (
	"github.com/maltegrosse/go-modemmanager"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yellow-sky/orap/api"
	"github.com/yellow-sky/orap/auth"
	"github.com/yellow-sky/orap/conf"
	_ "github.com/yellow-sky/orap/docs"
	"github.com/yellow-sky/orap/server_app"
	"github.com/yellow-sky/orap/swagger"
	"github.com/yellow-sky/orap/web"
	"os"
	"os/signal"
	"syscall"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start serving",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		initConfigManager()
		loadConfigs()
		initLogger()
		startServe()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

// TODO: Need FULL REFACTORING !!!!
// TODO: separate configs to app and core

var apiCfg = conf.ApiConfig{}
var authCfg = conf.AuthConfig{}

func startServe() {
	//Configure logger
	log := logrus.WithField("module", "main")

	// Read config
	if err := appConfigManager.UnmarshalKey(conf.ApiDefaultConfKey, &apiCfg); err != nil {
		log.Errorf("Error on load '%s' section of config: %s", conf.ApiDefaultConfKey, err.Error())
		os.Exit(4)
	}
	if err := appConfigManager.UnmarshalKey(conf.AuthDefaultConfKey, &authCfg); err != nil {
		log.Errorf("Error on load '%s' section of config: %s", conf.AuthDefaultConfKey, err.Error())
		os.Exit(4)
	}

	// Init services
	mmgr, err := modemmanager.NewModemManager()
	if err != nil {
		log.Errorf("Error on create modem manager: %s", err.Error())
		os.Exit(5)
		return
	}
	authService := auth.NewAuthService(authCfg)
	server := server_app.NewServer(apiCfg)
	web.InitWebService(server)
	api.InitApiService(server, &authService, mmgr)
	swagger.InitSwaggerService(server)

	// Start service
	go server.Run()

	// Wait stop signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Info("OS Signal: ", sig)
		done <- true
	}()
	log.Info("Server Started...")
	<-done
	server.Shutdown()
	log.Info("Exiting")
}
