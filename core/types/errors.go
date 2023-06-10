//
//
// @filename: errors.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidVersion         = sdkerrors.Register(ModuleName, 1100, "invalid nano module version")
	ErrInvalidAcknowledgement = sdkerrors.Register(ModuleName, 1101, "invalid acknowledgement")
	ErrAcknowledgementExists  = sdkerrors.Register(ModuleName, 1102, "acknowledgement for packet already exists")
)
