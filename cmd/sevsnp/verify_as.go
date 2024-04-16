// Copyright 2024 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package sevsnp

import (
	"os"

	"github.com/spf13/cobra"
)

var verifyValidArgs = []string{"relying-party"}

const SEVSNPTokenMediaType = "application/vnd.veraison.tsm-report+cbor"

var verifyAsCmd = &cobra.Command{
	Use:   "verify-as",
	Short: "use Veraison REST API to verify SEV-SNP tokens",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help() // nolint: errcheck
			os.Exit(0)
		}
	},
	ValidArgs: verifyValidArgs,
}

func init() {
	verifyAsCmd.AddCommand(rpCmd)
}
