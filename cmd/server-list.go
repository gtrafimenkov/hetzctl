// SPDX-License-Identifier: MIT
//
// Copyright © 2018 Gennady Trafimenkov <gennady.trafimenkov@gmail.com>

package cmd

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/gtrafimenkov/go-hetzner"
)

func init() {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List servers",
	}
	serverCmd.AddCommand(cmd)
	cmd.Run = func(cmd *cobra.Command, args []string) {
		username, password := getCredentials()
		client := hetzner.NewClient(username, password)
		servers, _, err := client.Server.ListServers()
		fatal(err)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"Server Name", "Number", "IP",
			"Product", "DC", "Traffic", "Flatrate", "Status",
			"Throttled", "Cancelled", "Paid Until"})

		for _, srv := range servers {
			table.Append([]string{
				srv.ServerName, strconv.Itoa(srv.ServerNumber), srv.ServerIP,
				srv.Product, srv.Dc, srv.Traffic, strconv.FormatBool(srv.Flatrate), srv.Status,
				strconv.FormatBool(srv.Throttled), strconv.FormatBool(srv.Cancelled), srv.PaidUntil,
			})
		}
		table.Render()
	}
}
