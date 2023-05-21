//
//
// @filename: common/types.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package common

import "github.com/NanoOfficial/smartnano/lib/errors"

type (
	NodeID          string
	NodeName        string
	NodeVersionName string
	BlockName       string
	ActionName      string
	DeliveryMethod  string
)

type TargetAction struct {
	Name    NodeName   `json:"name"`
	Version uint       `json:"version"`
	Block   BlockName  `json:"block"`
	Action  ActionName `json:"action"`
}

func (ta *TargetAction) Validate() error {
	if ta.Name.String() == "" || ta.Block.String() == "" || ta.Action.String() == "" {
		return errors.ErrInvalidTargetAction
	}

	return nil
}

func (nodeID NodeID) String() string {
	return string(nodeID)
}

func (name NodeName) String() string {
	return string(name)
}

func (name NodeVersionName) String() string {
	return string(name)
}

func (name BlockName) String() string {
	return string(name)
}

func (name ActionName) String() string {
	return string(name)
}
