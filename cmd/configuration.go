package cmd

import (
	"github.com/spf13/cobra"
)

var outCfgFile string

var configCmd = &cobra.Command{
	Use:   "configuration",
	Short: "Configuration management",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) {},
}

var genConfigCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate default config file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		initConfigManager()
		if err := appConfigManager.SafeWriteConfigAs(outCfgFile); err != nil {
			cmd.PrintErrln("Error on write template configuration file. ", err)
		} else {
			cmd.Printf("Default configuration written to file '%s'\n", outCfgFile)
		}
	},
}

var checkConfigCmd = &cobra.Command{
	Use:   "check",
	Short: "Check current config file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		initConfigManager()
		loadConfigs()
		cmd.Printf("Current local config is correct\n")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(genConfigCmd)
	genConfigCmd.PersistentFlags().StringVar(&outCfgFile, "out-path", "config.yaml", "output path for config file")

	configCmd.AddCommand(checkConfigCmd)
}
