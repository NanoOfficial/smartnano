//
//
// @filename: xcc/codec.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package xcc

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*XCC)(nil),
		&ChannelInfo{},
	)
}
