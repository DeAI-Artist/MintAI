package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"

	"github.com/DeAI-Artist/MintAI/core/chains"
	"github.com/DeAI-Artist/MintAI/core/services/chainlink"
	"github.com/DeAI-Artist/MintAI/core/services/relay"
)

type chainBatcher struct {
	app chainlink.Application
}

func (b *chainBatcher) loadByIDs(_ context.Context, keys dataloader.Keys) []*dataloader.Result {
	// Create a map for remembering the order of keys passed in
	keyOrder := make(map[string]int, len(keys))
	// Collect the keys to search for
	var chainIDs []relay.ChainID
	for ix, key := range keys {
		chainIDs = append(chainIDs, key.String())
		keyOrder[key.String()] = ix
	}

	// Fetch the chains
	cs, _, err := b.app.EVMORM().Chains(chainIDs...)
	if err != nil {
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	// Construct the output array of dataloader results
	results := make([]*dataloader.Result, len(keys))
	for _, c := range cs {
		ix, ok := keyOrder[c.ID]
		// if found, remove from index lookup map, so we know elements were found
		if ok {
			results[ix] = &dataloader.Result{Data: c, Error: nil}
			delete(keyOrder, c.ID)
		}
	}

	// fill array positions without any nodes
	for _, ix := range keyOrder {
		results[ix] = &dataloader.Result{Data: nil, Error: chains.ErrNotFound}
	}

	return results
}
