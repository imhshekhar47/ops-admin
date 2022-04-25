/*
Copyright © 2022 Himanshu Shekhar <himanshu.kiit@gmail.com>
Code ownership is with Himanshu Shekhar. Use without modifications.
*/
package cmd

import (
	"os"

	"github.com/imhshekhar47/ops-admin/config"
	"github.com/imhshekhar47/ops-admin/server"
	"github.com/imhshekhar47/ops-admin/service"
	"github.com/imhshekhar47/ops-admin/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	argStartGrpcPort uint16
	argStartRestPort uint16

	adminConfiguration *config.AdminConfiguration
	adminService       *service.AdminService
	adminServer        *server.AdminServer
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ops-admin",
	Short: "Ops Admin server",
	Long:  `Ops Admin application keeps track of all the agents and manages them`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ops-admin.yaml)")
}

func initConfig() {
	util.Logger.WithField("origin", "cmd::root").Traceln("entry: initConfig()")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.SetConfigName(".ops-admin")
		viper.SetConfigType("yaml")

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		util.Logger.WithField("origin", "cmd::root").Debugln("Using config file:", viper.ConfigFileUsed())
	}

	loadServerConfig()

	util.Logger.WithField("origin", "cmd::root").Traceln("exit: initConfig()")
}

func loadServerConfig() {
	util.Logger.WithField("origin", "cmd::root").Traceln("entry: loadServerConfig()")
	hostname := viper.GetString("server.hostname")
	if hostname == "" {
		hostname = util.GetHostname()
	}

	coreConiguration := config.CoreConfiguration{
		Version: viper.GetString("core.version"),
	}

	adminConfiguration = &config.AdminConfiguration{
		Core:     coreConiguration,
		Hostname: hostname,
		Uuid:     util.Encode(hostname),
	}

	util.Logger.WithField("origin", "cmd::root").Debugln("agent_configuration", adminConfiguration)
	util.Logger.WithField("origin", "cmd::root").Traceln("exit: loadServerConfig()")
}
