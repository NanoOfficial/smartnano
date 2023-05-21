//
//
// @filename: common/common.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package common

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
