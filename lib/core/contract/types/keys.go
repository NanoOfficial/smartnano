//
//
// @filename: contract/types/keys.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package types

import (
	"fmt"

	"github.com/NanoOfficial/smartnano/lib/core/types"
	"github.com/NanoOfficial/smartnano/lib/utils"
)

const (
	ModuleName = "cross-contract"
)

const (
	KeyContractCallResultPrefix uint8 = iota
)

func KeyPrefixBytes(prefix uint8) []byte {
	return []byte(fmt.Sprintf("%d/", prefix))
}

func KeyContractCallResult(txID types.TxID, txIndex types.TxIndex) []byte {
	return append(
		append(
			KeyPrefixBytes(KeyContractCallResultPrefix),
			txID[:]...,
		),
		utils.Uint32ToBigEndian(txIndex)...,
	)
}
