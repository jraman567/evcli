// Copyright 2022 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package sevsnp

import (
	"os"

	"github.com/spf13/cobra"
)

var cmdValidArgs = []string{"verify-as", "create"}

var Cmd = &cobra.Command{
	Use:   "sevsnp",
	Short: "SEVSNP token manipulation",
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