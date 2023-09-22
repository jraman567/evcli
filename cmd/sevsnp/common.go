// Copyright 2022 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package sevsnp

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/afero"
)

func loadClaimsFromFile(fs afero.Fs, fn string, validate bool) (Evidence, error) {
	buf, err := afero.ReadFile(fs, fn)

	return claimsFromJSON(buf, validate)
}

func claimsFromJSON(j []byte, validate bool) (Evidence, error) {
	var err	error
	var ev	Evidence

	err = json.Unmarshal(j, &ev)
	if err == nil {
		return ev, nil
	}

	return ev, fmt.Errorf("unable to load claims from JSON: (%v)", err)
}
