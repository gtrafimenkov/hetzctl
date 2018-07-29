// SPDX-License-Identifier: MIT
//
// Copyright © 2018 Gennady Trafimenkov <gennady.trafimenkov@gmail.com>

package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func requireExactArgs(cmd *cobra.Command, args string) {
	n := len(strings.Split(args, " "))
	cmd.Args = cobra.ExactArgs(n)
	cmd.Short = cmd.Short + " " + args
}

func getCredentials() (username string, password string) {
	username = viper.GetString("username")
	password = viper.GetString("password")

	if username == "" {
		log.Fatal("username is empty")
	}

	if password == "" {
		log.Fatal("password is empty")
	}

	return
}
