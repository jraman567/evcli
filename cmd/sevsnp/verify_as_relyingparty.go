// Copyright 2025 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package sevsnp

import (
	"fmt"
	"strings"

	"github.com/jraman567/evcli/v2/common"
	"github.com/jraman567/tokens"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/veraison/apiclient/verification"
)

// ToDo: Must be generated randomly. But RATSd will supersede this tool, so don't bother fixing
const sessionNonce = "BwYFBAMCAQAPDg0MCwoJCBcWFRQTEhEQHx4dHBsaGRg="

var (
	rpEvidenceFile *string
	rpServerURL    *string
)

var (
	rpVeraisonClient common.IVeraisonClient = &verification.ChallengeResponseConfig{}
	rpCmd                                   = NewRelyingPartyCmd(common.Fs, rpVeraisonClient)
)

func NewRelyingPartyCmd(fs afero.Fs, veraisonClient common.IVeraisonClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relying-party",
		Short: "Emulate a Relying Party",
		Long: `This command implements the "relying party mode" of a
challenge-response interaction.

    evcli sev-snp verify-as relying-party \
        --api-server=https://localhost:8080/challenge-response/v1/newSession \
        --token=cmd/sevsnp/sample/SNP-EAT.json

	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var tsm tokens.TSMReport

			tokenBuf, err := afero.ReadFile(fs, *rpEvidenceFile)
			if err != nil {
				return err
			}

			err = tsm.FromJSON(tokenBuf)
			if err != nil {
				return err
			}

			cborToken, err := tsm.ToCBOR()
			if err != nil {
				return err
			}

			err = veraisonClient.SetNonce([]byte(sessionNonce))
			if err != nil {
				return err
			}

			err = veraisonClient.SetSessionURI(*rpServerURL)
			if err != nil {
				return err
			}

			eb := relyingPartyEvidenceBuilder{Token: cborToken}
			if err = veraisonClient.SetEvidenceBuilder(eb); err != nil {
				return err
			}

			veraisonClient.SetDeleteSession(true)

			attestationResults, err := veraisonClient.Run()
			if err != nil {
				return err
			}

			fmt.Println(string(attestationResults))

			return nil
		},
	}

	rpEvidenceFile = cmd.Flags().StringP(
		"token", "t", "", "file containing a SEV attestation token",
	)

	rpServerURL = cmd.Flags().StringP(
		"api-server", "s", "", "URL of the Veraison verification API",
	)

	return cmd
}

type relyingPartyEvidenceBuilder struct {
	Token []byte
}

func (eb relyingPartyEvidenceBuilder) BuildEvidence(_ []byte, accept []string) ([]byte, string, error) {
	for _, ct := range accept {
		if ct == SEVSNPTokenMediaType {
			return eb.Token, SEVSNPTokenMediaType, nil
		}
	}

	return nil, "", fmt.Errorf("expecting media type %s, got %s", SEVSNPTokenMediaType, strings.Join(accept, ", "))
}

func init() {
	if err := rpCmd.MarkFlagRequired("token"); err != nil {
		panic(err)
	}
	if err := rpCmd.MarkFlagRequired("api-server"); err != nil {
		panic(err)
	}
}
