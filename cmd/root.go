// SPDX-License-Identifier: MIT
//
// Copyright © 2018 Gennady Trafimenkov <gennady.trafimenkov@gmail.com>

package cmd

import (
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "hetzctl",
	Short: "hetzctl controls your servers in Hetzner",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	fatal(err)
}

const defaultConfigName = ".hetzctl.yaml"
const defaultConfigPath = "$HOME/" + defaultConfigName

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "viper-config", defaultConfigPath, "viper configuration file")
	rootCmd.PersistentFlags().String("username", "", "username for access Hetzner API")
	rootCmd.PersistentFlags().String("password", "", "password for access Hetzner API")
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != defaultConfigPath {
		viper.SetConfigFile(cfgFile)
		err := viper.ReadInConfig()
		fatal(err)
	} else {
		u, err := user.Current()
		fatal(err)
		configPath := path.Join(u.HomeDir, defaultConfigName)
		_, err = os.Stat(configPath)
		if err == nil {
			// config file exists
			viper.SetConfigFile(configPath)
			err := viper.ReadInConfig()
			fatal(err)
		}
	}

	viper.SetEnvPrefix("hetzctl")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()
}
