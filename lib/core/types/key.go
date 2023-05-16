//
//
// @filename: key.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName   = "cross"
	Version      = "cross-1"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	PortID       = "cross"
)

var (
	InitiatorKeyPrefix     = []byte("initiator")
	AuthKeyPrefix          = []byte("auth")
	AtomicKeyPrefix        = []byte("atomic")
	ContractManagerPrefix  = []byte("cmanager")
	ContractStoreKeyPrefix = []byte("cstore")
)

type PrefixStoreKey struct {
	StoreKey sdk.StoreKey
	Prefix   []byte
}

var _ sdk.StoreKey = (*PrefixStoreKey)(nil)

func NewPrefixStoreKey(storeKey sdk.StoreKey, prefix []byte) *PrefixStoreKey {
	return &PrefixStoreKey{
		StoreKey: storeKey,
		Prefix:   prefix,
	}
}

func (sk *PrefixStoreKey) Name() string {
	return sk.StoreKey.Name()
}

func (sk *PrefixStoreKey) String() string {
	return fmt.Sprintf("PrefixStoreKey{%p, %s, %v}", sk, sk.Name(), sk.Prefix)
}
