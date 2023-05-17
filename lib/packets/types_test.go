//
//
// @filename: packets/types_test.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package packets

var _ PacketDataPayload = (*TestPacketDataPayload)(nil)

func (TestPacketDataPayload) ValidateBasic() error {
	return nil
}
