// SPDX-License-Identifier: MIT
//
// Copyright © 2018 Gennady Trafimenkov <gennady.trafimenkov@gmail.com>

package cmd

import (
	"fmt"
	"os"

	hetzner "github.com/gtrafimenkov/go-hetzner"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get list of failover ip addresses",
		Run:   failoverListRun,
	}
	failoverCmd.AddCommand(cmd)
}

func failoverListRun(cmd *cobra.Command, args []string) {
	username, password := getCredentials()
	client := hetzner.NewClient(username, password)
	ips, _, err := client.Failover.List()
	fatal(err)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"IP",
		"Netmask",
		"ServerIP",
		"ServerNumber",
		"ActiveServerIP",
	})

	for _, ip := range ips {
		table.Append([]string{
			ip.IP,
			ip.Netmask,
			ip.ServerIP,
			fmt.Sprintf("%d", ip.ServerNumber),
			ip.ActiveServerIP,
		})
	}
	table.Render()
}
