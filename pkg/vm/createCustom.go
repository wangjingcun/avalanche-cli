// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package vm

import (
	"fmt"
	"os"

	"github.com/ava-labs/avalanche-cli/cmd/prompts"
	"github.com/ava-labs/avalanche-cli/ux"
	"github.com/ava-labs/avalanchego/utils/logging"
)

func CreateCustomGenesis(name string, log logging.Logger) ([]byte, error) {
	ux.PrintToUser(fmt.Sprintf("creating custom VM subnet %s", name), log)

	genesisPath, err := prompts.CaptureExistingFilepath("Enter path to custom genesis")
	if err != nil {
		return []byte{}, err
	}

	genesisBytes, err := os.ReadFile(genesisPath)
	return genesisBytes, err
}
