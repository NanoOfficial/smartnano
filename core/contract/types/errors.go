//
//
// @filename: contract/types/errors.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package types

import "fmt"

type ErrContractCall struct {
	err error
}

func NewErrContractCall(err error) ErrContractCall {
	return ErrContractCall{err: err}
}

func (e ErrContractCall) Error() string {
	return fmt.Sprintf("failed to call the contract: %v", e.err)
}
