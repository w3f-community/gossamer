// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package genesis

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/ChainSafe/gossamer/dot/types"
	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/ChainSafe/gossamer/lib/trie"
	"github.com/ChainSafe/gossamer/lib/utils"
	"github.com/stretchr/testify/require"
)

const testProtocolID = "/gossamer/test/0"

var testBootnodes = []string{
	"/dns4/p2p.cc3-0.kusama.network/tcp/30100/p2p/QmeCit3Nif4VfNqrEJsdYHZGcKzRCnZvGxg6hha1iNj4mk",
	"/dns4/p2p.cc3-1.kusama.network/tcp/30100/p2p/QmchDJtEGiEWf7Ag58HNoTg9jSGzxkSZ23VgmF6xiLKKsZ",
}

var testEndpoints = []interface{}{}
var testEndpoint1 = []interface{}{"wss://telemetry.polkadot.io/submit/", float64(1)}
var testProperties = map[string]interface{}{
	"ss58Format":    float64(0),
	"tokenDecimals": float64(10),
	"tokenSymbol":   "DOT",
}

var testForkBlocks = []string{"fork1", "forkBlock2"}

var testBadBlocks = []string{"badBlock1", "badBlock2"}

// TestGenesis instance of Genesis struct for testing
var TestGenesis = &Genesis{
	Name:               "gossamer",
	ID:                 "gossamer",
	Bootnodes:          testBootnodes,
	TelemetryEndpoints: append(testEndpoints, testEndpoint1),
	ProtocolID:         testProtocolID,
	Properties:         testProperties,
	ForkBlocks:         testForkBlocks,
	BadBlocks:          testBadBlocks,
}

// TestFieldsHR instance of human-readable Fields struct for testing, use with TestGenesis
var TestFieldsHR = Fields{
	Raw: map[string]map[string]string{},
	Runtime: map[string]map[string]interface{}{
		"System": {
			"code": "mocktestcode",
		},
	},
}

// TestFieldsRaw instance of raw Fields struct for testing use with TestGenesis
var TestFieldsRaw = Fields{
	Raw: map[string]map[string]string{
		"top": {
			"0x3a636f6465": "mocktestcode",
			common.BytesToHex(common.UpgradedToDualRefKey): "0x01",
		},
	},
}

// CreateTestGenesisJSONFile utility to create mock test genesis JSON file
func CreateTestGenesisJSONFile(t *testing.T, asRaw bool) (filename string) {
	tGen := &Genesis{
		Name:       "test",
		ID:         "",
		Bootnodes:  nil,
		ProtocolID: "",
		Genesis:    Fields{},
	}

	if asRaw {
		tGen.Genesis = Fields{
			Raw: map[string]map[string]string{},
			Runtime: map[string]map[string]interface{}{
				"System": {
					"code": "mocktestcode",
				},
			},
		}
	} else {
		tGen.Genesis = TestFieldsHR
	}

	bz, err := json.Marshal(tGen)
	require.NoError(t, err)
	filename = filepath.Join(t.TempDir(), "genesis-test")
	err = os.WriteFile(filename, bz, os.ModePerm)
	require.NoError(t, err)
	return filename
}

// NewTestGenesisWithTrieAndHeader generates genesis, genesis trie and genesis header
func NewTestGenesisWithTrieAndHeader(t *testing.T) (*Genesis, *trie.Trie, *types.Header) {
	genesisPath := utils.GetGssmrGenesisRawPathTest(t)
	gen, err := NewGenesisFromJSONRaw(genesisPath)
	require.NoError(t, err)

	tr, h := newGenesisTrieAndHeader(t, gen)
	return gen, tr, h
}

// NewDevGenesisWithTrieAndHeader generates test dev genesis, genesis trie and genesis header
func NewDevGenesisWithTrieAndHeader(t *testing.T) (*Genesis, *trie.Trie, *types.Header) {
	genesisPath := utils.GetDevGenesisPath(t)

	gen, err := NewGenesisFromJSONRaw(genesisPath)
	require.NoError(t, err)

	tr, h := newGenesisTrieAndHeader(t, gen)
	return gen, tr, h
}

func newGenesisTrieAndHeader(t *testing.T, gen *Genesis) (*trie.Trie, *types.Header) {
	genTrie, err := NewTrieFromGenesis(gen)
	require.NoError(t, err)

	genesisHeader, err := types.NewHeader(common.NewHash([]byte{0}),
		genTrie.MustHash(), trie.EmptyHash, 0, types.NewDigest())
	require.NoError(t, err)

	return genTrie, genesisHeader
}
