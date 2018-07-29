// SPDX-License-Identifier: MIT
//
// Copyright © 2018 Gennady Trafimenkov <gennady.trafimenkov@gmail.com>

package cmd

import (
	"github.com/spf13/cobra"
)

var failoverCmd = &cobra.Command{
	Use:   "failover",
	Short: "Operations on faiolver ip addresses",
}

func init() {
	rootCmd.AddCommand(failoverCmd)
}
