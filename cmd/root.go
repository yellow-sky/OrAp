package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yellow-sky/orap/conf"
	"os"
)

var (
	cfgFile       string
	cfgConsulAddr string
	cfgConsulKey  string

	appConfigManager *conf.ConfigManager
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "orap",
	Short:   "OrangePi Access Point",
	Long:    `Simple service for organize access point with OrangePi board`,
	Version: "0.0.1",
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initCommon)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "local config file (default is 'config.yaml')")
}

func initCommon() {
}

func initConfigManager() {
	appConfigManager = conf.NewConfigManager()
	appConfigManager.SetMappedDefault(conf.LoggerDefaultConfKey, conf.LoggerGetConfigDefaults())
	appConfigManager.SetMappedDefault(conf.ApiDefaultConfKey, conf.ApiGetConfigDefaults())
	appConfigManager.SetMappedDefault(conf.AuthDefaultConfKey, conf.AuthGetConfigDefaults())
}

func loadConfigs() {
	if err := appConfigManager.LoadLocalConfig(cfgFile); err != nil {
		os.Exit(3)
	}
	appConfigManager.Fix()
}

func initLogger() {
	logCfg := conf.LogConfig{}
	if err := appConfigManager.UnmarshalKey(conf.LoggerDefaultConfKey, &logCfg); err != nil {
		logrus.Errorln("Error on load log config. ", err)
		return
	} else {
		logrus.SetLevel(logCfg.GetLevelNum())
		if logCfg.Format == conf.JsonOutputFormat {
			logrus.SetFormatter(&logrus.JSONFormatter{})
		}
	}
}
