//
//
// @filename: ibc/utils.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package ibc

import (
	"github.com/NanoOfficial/smartnano/lib/core/packets"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
)

type capturePacketSender struct {
	inner   packets.PacketSender
	packets []packets.OutgoingPacket
}

var _ packets.PacketSender = (*capturePacketSender)(nil)

func NewCapturePacketSender(ps packets.PacketSender) *capturePacketSender {
	return &capturePacketSender{inner: ps}
}

func (ps *capturePacketSender) SendPacket(
	ctx sdk.Context,
	channelCap *capabilitytypes.Capability,
	packet packets.OutgoingPacket,
) error {
	if err := ps.inner.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}
	ps.packets = append(ps.packets, packet)
	return nil
}

func (ps *capturePacketSender) Packets() []packets.OutgoingPacket {
	return ps.packets
}

func (ps *capturePacketSender) Clear() {
	ps.packets = nil
}
