//
//
// @filename: auth/types.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package types

import "fmt"

type ErrIDNotFound struct {
	id []byte
}

func NewErrIDNotFound(id []byte) ErrIDNotFound {
	return ErrIDNotFound{id: id}
}

func (e ErrIDNotFound) Error() string {
	return fmt.Sprintf("id '%x' not found", e.id)
}
