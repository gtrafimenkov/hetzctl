// SPDX-License-Identifier: MIT
//
// Copyright © 2018 Gennady Trafimenkov <gennady.trafimenkov@gmail.com>

package cmd

import (
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: `Operations on the server, e.g. "server list"`,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
