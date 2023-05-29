//
//
// @filename: contract/types/store.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CommitMode = uint8

const (
	UnspecifiedMode CommitMode = iota + 1
	BasicMode
	AtomicMode
)

type CommitStoreI interface {
	Precommit(ctx sdk.Context, id []byte) error
	Abort(ctx sdk.Context, id []byte) error
	Commit(ctx sdk.Context, id []byte) error
	CommitImmediately(ctx sdk.Context)
}
