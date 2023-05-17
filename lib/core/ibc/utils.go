//
//
// @filename: ibc/utils.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package ibc

import "github.com/NanoOfficial/smartnano/lib/core/packets"

type capturePacketSender struct {
	inner   packets.PacketSender
	packets []packets.OutgoingPacket
}

var _ packets.PacketSender = (*capturePacketSender)(nil)

func newCapturePacketSender(ps packets.PacketSender) *capturePacketSender {
	return &capturePacketSender{inner: ps}
}

func (ps *capturePacketSender) Packets() []packets.OutgoingPacket {
	return ps.packets
}

func (ps *capturePacketSender) Clear() {
	ps.packets = nil
}
