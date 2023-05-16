//
//
// @filename: utils/any.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package utils

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
)

func PackAny(msg proto.Message) (*types.Any, error) {
	return types.NewAnyWithValue(msg)
}
