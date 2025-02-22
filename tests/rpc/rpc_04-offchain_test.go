// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/ChainSafe/gossamer/tests/utils"
	"github.com/stretchr/testify/require"
)

func TestOffchainRPC(t *testing.T) {
	if utils.MODE != rpcSuite {
		t.Log("Going to skip RPC suite tests")
		return
	}

	testCases := []*testCase{
		{ //TODO
			description: "test offchain_localStorageSet",
			method:      "offchain_localStorageSet",
			skip:        true,
		},
		{ //TODO
			description: "test offchain_localStorageGet",
			method:      "offchain_localStorageGet",
			skip:        true,
		},
		{ //TODO
			description: "test offchain_localStorageGet",
			method:      "offchain_localStorageGet",
			skip:        true,
		},
	}

	t.Log("starting gossamer...")
	nodes, err := utils.InitializeAndStartNodes(t, 1, utils.GenesisDefault, utils.ConfigDefault)
	require.NoError(t, err)

	time.Sleep(time.Second) // give server a second to start

	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			ctx := context.Background()
			getResponseCtx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()
			_ = getResponse(getResponseCtx, t, test)
		})
	}

	t.Log("going to tear down gossamer...")
	errList := utils.TearDown(t, nodes)
	require.Len(t, errList, 0)
}
