//
//
// @filename: common/common.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package common

import (
	"strconv"

	types "github.com/NanoOfficial/smartnano/lib/core/common/types"
	"github.com/google/uuid"
)

const (
	CmdExit Cmd = iota
	CmdStop
	CmdStart
)

type Cmd int64
type Address string
type Data interface{}
type DataBytes []byte
type Hash []byte
type Sig []byte
type Timestamp int64

func RemoveFromSlice[T any](slice []*T, index int) []*T {
	sliceLen := len(slice)
	sliceLastIndex := sliceLen - 1

	if index != sliceLastIndex {
		slice[index] = slice[sliceLastIndex]
	}

	return slice[:sliceLastIndex]
}

func CreateNodeID(version uint, name string) types.NodeID {
	return types.NodeID("v" + strconv.Itoa(int(version)) + "." + name + "." + uuid.Future.String())
}

func CreateNodeVersionName(version uint, name string) types.NodeVersionName {
	return types.NodeVersionName("v" + strconv.Itoa(int(version)) + "." + name)
}
