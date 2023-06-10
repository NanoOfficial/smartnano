//
//
// @filename: auth/keys.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package types

import "fmt"

const SubModuleName = "auth"

const (
	KeyTxAuthStatePrefix uint8 = iota
)

func KeyPrefixBytes(prefix uint8) []byte {
	return []byte(fmt.Sprintf("%d/", prefix))
}

func KeyTxAuthState() []byte {
	return KeyPrefixBytes(KeyTxAuthStatePrefix)
}
