// SPDX-License-Identifier: MIT
//
// Copyright © 2018 Gennady Trafimenkov <gennady.trafimenkov@gmail.com>

package cmd

import (
	"github.com/gtrafimenkov/go-hetzner"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "switch",
		Short: "Route traffic to another server",
		Long:  "This commands routes trafic coming to FAILOVER_IP to the server with NEW_ACTIVE_IP",
	}
	requireExactArgs(cmd, "FAILOVER_IP NEW_ACTIVE_IP")
	failoverCmd.AddCommand(cmd)
	cmd.Run = func(cmd *cobra.Command, args []string) {
		failoverIP, newActiveIP := args[0], args[1]

		username, password := getCredentials()
		client := hetzner.NewClient(username, password)
		_, _, err := client.Failover.Switch(failoverIP, newActiveIP)
		fatal(err)
	}
}
