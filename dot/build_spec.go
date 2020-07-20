// Copyright 2020 ChainSafe Systems (ON) Corp.
// This file is part of gossamer.
//
// The gossamer library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gossamer library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gossamer library. If not, see <http://www.gnu.org/licenses/>.
package dot

import (
	"fmt"
	"github.com/ChainSafe/gossamer/dot/state"
	"github.com/ChainSafe/gossamer/lib/trie"
	log "github.com/ChainSafe/log15"
)

func BuildSpec(basepath string) error {

	stateSrvc := state.NewService(basepath, log.LvlDebug)

	// start state service (initialize state database)
	err := stateSrvc.Start()
	if err != nil {
		return err
	}

	//// load most recent state from database
	//latestState, err := state.LoadLatestStorageHash(stateSrvc.DB())
	//if err != nil {
	//	logger.Error("failed to load latest state root hash", "error", err)
	//}
	//
	//// load most recent state from database
	//err = stateSrvc.Storage.LoadFromDB(latestState)
	//if err != nil {
	//	logger.Error("failed to load latest state from database", "error", err)
	//}

	ent := stateSrvc.Storage.Entries()
	for k, v := range ent {
		fmt.Printf("key %x vL %v\n", k, len(v))
	}

	gh := stateSrvc.Block.GenesisHash()
	fmt.Printf("Genesis Hash %s\n", gh)

	rootTrie := trie.NewEmptyTrie()
	err = state.LoadTrie(stateSrvc.DB(), rootTrie, gh)
	if err != nil {
		logger.Error("failed te load genesis state trie", "error", err)
	}

	fmt.Printf("root trie %v\n", rootTrie)

	return nil
}
