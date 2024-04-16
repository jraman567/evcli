// Copyright 2024 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package sevsnp

import (
	"fmt"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/veraison/apiclient/verification"
	"github.com/jraman567/evcli/v2/common"
	"github.com/jraman567/tokens"
)

var (
	relyingPartyTokenFile *string
	relyingPartyAPIURL    *string
)

var (
	relyingPartyVeraisonClient common.IVeraisonClient = &verification.ChallengeResponseConfig{}
	relyingPartyCmd                                   = NewRelyingPartyCmd(common.Fs, relyingPartyVeraisonClient)
)

func NewRelyingPartyCmd(fs afero.Fs, veraisonClient common.IVeraisonClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relying-party",
		Short: "Emulate a Relying Party",
		Long: `This command implements the "relying party mode" of a
challenge-response interaction.

	evcli sevsnp verify-as relying-party \
	              --api-server=https://veraison.example/challenge-response/v1/newSession \
	              --token=sevsnp-token.cbor

	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			tokenBuf, err := afero.ReadFile(fs, *relyingPartyTokenFile)
			if err != nil {
				return err
			}

			token, err := tokens.GetToken(tokenBuf)
			if err != nil {
				return err
			}

			err = veraisonClient.SetNonce(token.GetNonce())
			if err != nil {
                                return err
                        }

			err = veraisonClient.SetSessionURI(*relyingPartyAPIURL)
			if err != nil {
				return err
			}

			eb := relyingPartyEvidenceBuilder{Token: token.GetBytes()}
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

	relyingPartyTokenFile = cmd.Flags().StringP(
		"token", "t", "", "file containing a SEV attestation token",
	)

	relyingPartyAPIURL = cmd.Flags().StringP(
		"api-server", "s", "", "URL of the Veraison verification API",
	)

	return cmd
}

type relyingPartyEvidenceBuilder struct {
	Token []byte
}

func (eb relyingPartyEvidenceBuilder) BuildEvidence(nonce []byte, accept []string) ([]byte, string, error) {
	for _, ct := range accept {
		if ct == SEVSNPTokenMediaType {
			return eb.Token, SEVSNPTokenMediaType, nil
		}
	}

	return nil, "", fmt.Errorf("expecting media type %s, got %s", SEVSNPTokenMediaType, strings.Join(accept, ", "))
}

func init() {
	if err := relyingPartyCmd.MarkFlagRequired("token"); err != nil {
		panic(err)
	}
	if err := relyingPartyCmd.MarkFlagRequired("api-server"); err != nil {
		panic(err)
	}
}
