package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	"{{GitServer}}/{{Organization}}/{{Name}}/implem/gin.server"
	"{{GitServer}}/{{Organization}}/{{Name}}/infra"
	"{{GitServer}}/{{Organization}}/{{Name}}/uc"
)

// Build number and versions injected at compile time, set yours
var (
	Version = "unknown"
	Build   = "unknown"
)

// the command to run the server
var rootCmd = &cobra.Command{
	Use:   "{{Name}}",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Build: %s\nVersion: %s\n", Build, Version)
	},
}

func main() {
	rootCmd.AddCommand(versionCmd)
	cobra.OnInitialize(infra.CobraInitialization)

	infra.LoggerConfig(rootCmd)
	infra.ServerConfig(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal()
	}
}

func run() {
	ginServer := infra.NewServer(
		viper.GetInt("server.port"),
		infra.DebugMode,
	)

	rH := server.NewRouter(
		uc.NewInteractor(),
	)
	rH.SetRoutes(ginServer.Router)

	ginServer.Start()
}
