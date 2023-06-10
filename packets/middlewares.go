//
//
// @filename: packets/middlewares.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package packets

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
)

type PacketMiddleware interface {
	HandleMsg(ctx sdk.Context, msg sdk.Msg, ps PacketSender) (sdk.Context, PacketSender, error)
	HandlePacket(ctx sdk.Context, packet IncomingPacket, ps PacketSender, as ACKSender) (sdk.Context, PacketSender, ACKSender, error)
	HandleACK(ctx sdk.Context, packet IncomingPacket, ack IncomingPacketAcknowledgement, ps PacketSender) (sdk.Context, PacketSender, error)
}

type nopPacketMiddleware struct{}

var _ PacketMiddleware = (*nopPacketMiddleware)(nil)

func NewNOPPacketMiddleware() PacketMiddleware {
	return nopPacketMiddleware{}
}

func (m nopPacketMiddleware) HandleMsg(ctx sdk.Context, msg sdk.Msg, ps PacketSender) (sdk.Context, PacketSender, error) {
	return ctx, ps, nil
}

func (m nopPacketMiddleware) HandlePacket(ctx sdk.Context, packet IncomingPacket, ps PacketSender, as ACKSender) (sdk.Context, PacketSender, ACKSender, error) {
	return ctx, ps, as, nil
}

func (m nopPacketMiddleware) HandleACK(ctx sdk.Context, packet IncomingPacket, ack IncomingPacketAcknowledgement, ps PacketSender) (sdk.Context, PacketSender, error) {
	return ctx, ps, nil
}

type PacketSender interface {
	SendPacket(
		ctx sdk.Context,
		channelCap *capabilitytypes.Capability,
		packet OutgoingPacket,
	) error
}

type basicPacketSender struct {
	channelKeeper ChannelKeeper
}

func NewBasicPacketSender(channelKeeper ChannelKeeper) PacketSender {
	return basicPacketSender{channelKeeper: channelKeeper}
}

func (s basicPacketSender) SendPacket(
	ctx sdk.Context,
	channelCap *capabilitytypes.Capability,
	packet OutgoingPacket,
) error {
	return s.channelKeeper.SendPacket(ctx, channelCap, packet)
}

type ACKSender interface {
	SendACK(
		ctx sdk.Context,
		ack OutgoingPacketAcknowledgement,
	) error
}

type basicACKSender struct{}

var _ ACKSender = (*basicACKSender)(nil)

func NewBasicACKSender() ACKSender {
	return &basicACKSender{}
}

func (s *basicACKSender) SendACK(ctx sdk.Context, ack OutgoingPacketAcknowledgement) error {
	return nil
}
