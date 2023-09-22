// Copyright 2022 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package sevsnp

import (
	"errors"
	"fmt"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/jraman567/evcli/v2/common"
)

var (
	createClaimsFile   *string
	createTokenFile    *string
	createTokenProfile *string
	allowInvalidClaims *bool
)

var createCmd = NewCreateCmd(common.Fs)

func NewCreateCmd(fs afero.Fs) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a SEVSNP attestation token from the supplied claims",
		Long: `Create a SEVSNP attestation token from the JSON-encoded claims, optionally specifying the wanted profile

Create a SEVSNP attestation token from claims contained in claims.json, and save the result to my.cbor:

	evcli sevsnp create --claims=claims.json --token=my.cbor

Or, equivalently:

	evcli psa sevsnp -c claims.json -k es256.jwk -t my.cbor

Create a SEVSNP attestation token from claims contained in te-profile1.json
(using PSA_IOT_PROFILE_1), sign with es256.jwk and save the result to
te-profile1.cbor:

	evcli psa create -c te-profile1.json -k es256.jwk -p PSA_IOT_PROFILE_1

Note that the default profile is http://arm.com/psa/2.0.0.
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			validate := !*allowInvalidClaims

			if err := checkProfile(createTokenProfile); err != nil {
				return err
			}

			evidence, err := loadClaimsFromFile(fs, *createClaimsFile, validate)
			if err != nil {
				return err
			}

			var cwt []byte
			cwt, err = em.Marshal(&evidence)
			if err != nil {
				return fmt.Errorf("failed to convert evidence to CBOR: %w", err)
			}

			fn := tokenFileName()

			err = afero.WriteFile(fs, fn, cwt, 0644)
			if err != nil {
				return fmt.Errorf("error saving SEVSNP token to file %s: %w", fn, err)
			}

			fmt.Printf(">> %q successfully created\n", fn)

			return nil
		},
	}

	createClaimsFile = cmd.Flags().StringP(
		"claims", "c", "", "JSON file containing the PSA attestation claims to be signed",
	)

	createTokenFile = cmd.Flags().StringP(
		"token", "t", "", "name of the file where the produced PSA attestation token will be stored",
	)

	createTokenProfile = cmd.Flags().StringP(
		"profile", "p", "http://amd.com/sev-snp", "name of the SEVSNP profile to use",
	)

	allowInvalidClaims = cmd.Flags().BoolP(
		"allow-invalid", "I", false,
		"Do not validate provided claims, allowing invalid tokens to be generated. "+
			"This is intended for testing.",
	)

	return cmd
}

func checkProfile(profile *string) error {
	if profile == nil {
		return errors.New("nil profile")
	}

	switch *profile {
	case "http://amd.com/sev-snp":
		return nil
	}

	return fmt.Errorf(
		"wrong profile %s: allowed profile http://amd.com/sev-snp",
		*profile,
	)
}

func tokenFileName() string {
	if createTokenFile == nil || *createTokenFile == "" {
		return common.MakeFileName(".", *createClaimsFile, ".cbor")
	}

	return *createTokenFile
}

func init() {
	if err := createCmd.MarkFlagRequired("claims"); err != nil {
		panic(err)
	}
	if err := createCmd.MarkFlagRequired("key"); err != nil {
		panic(err)
	}
}
