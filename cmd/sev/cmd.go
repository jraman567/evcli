// Copyright 2023 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package sev

import (
	"os"

	"github.com/spf13/cobra"
)

var cmdValidArgs = []string{"verify-as", "create"}

var Cmd = &cobra.Command{
	Use:   "sev",
	Short: "SEV token manipulation",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help() // nolint: errcheck
			os.Exit(0)
		}
	},
	ValidArgs: cmdValidArgs,
}

func init() {
	Cmd.AddCommand(createCmd)
	Cmd.AddCommand(verifyAsCmd)
}
