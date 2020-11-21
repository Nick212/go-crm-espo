package cmd

import (
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Execute the application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Short: "ESPO CRM Ecosystem",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logrus.Info("Version " + viper.GetString("version"))
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./configs/default.yaml)")
}

var configFile string

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("default")
		viper.AddConfigPath("./configs")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Fatal("Unable to read config from file")
	}
}
