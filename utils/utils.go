//
//
// @filename: utils/utils.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package utils

import "encoding/binary"

func Uint32ToBigEndian(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

func BigEndianToUint32(bz []byte) uint32 {
	if len(bz) == 0 {
		return 0
	}

	return binary.BigEndian.Uint32(bz)
}
