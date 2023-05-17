//
//
// @filename: xcc/types.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package xcc

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type XCC interface {
	proto.Message
	Type() string
	Equal(XCC) bool
	String() string
}

var _ XCC = (*ChannelInfo)(nil)

func (ci ChannelInfo) Type() string {
	return "channelinfo"
}

func (ci *ChannelInfo) Equal(other XCC) bool {
	oi, ok := other.(*ChannelInfo)
	if !ok {
		return false
	}
	return ci.Port == oi.Port && ci.Channel == oi.Channel
}

type XCCResolver interface {
	ResolveCrossChainChannel(ctx sdk.Context, xcc XCC) (*ChannelInfo, error)
	ResolveChannel(ctx sdk.Context, channel *ChannelInfo) (XCC, error)
	ConvertCrossChainChannel(ctx sdk.Context, calleeXCC XCC, callerXCC XCC) (calleeXCCOnCaller XCC, err error)
	GetSelfCrossChainChannel(ctx sdk.Context) XCC
	IsSelfCrossChainChannel(ctx sdk.Context, xcc XCC) bool
	Capabilities() XCCResolverCapabilities
}

type CommitProtocol = uint32

type XCCResolverCapabilities interface {
	CrossChainCalls(ctx sdk.Context) bool
}

type crossChainChannelResolverCapabilities struct{}

func (c crossChainChannelResolverCapabilities) CrossChainCalls(ctx sdk.Context) bool {
	return true
}

type ChannelInfoResolver struct {
	channelKeeper ChannelKeeper
}

func NewChannelInfoResolver(channelKeeper ChannelKeeper) ChannelInfoResolver {
	return ChannelInfoResolver{
		channelKeeper: channelKeeper,
	}
}

var _ XCCResolver = (*ChannelInfoResolver)(nil)

func (r ChannelInfoResolver) ResolveCrossChainChannel(ctx sdk.Context, xcc XCC) (*ChannelInfo, error) {
	ci, ok := xcc.(*ChannelInfo)
	if !ok {
		return nil, fmt.Errorf("cannot resolve '%v'", xcc)
	}
	return ci, nil
}

func (r ChannelInfoResolver) ResolveChannel(ctx sdk.Context, channel *ChannelInfo) (XCC, error) {
	_, found := r.channelKeeper.GetChannel(ctx, channel.Port, channel.Channel)
	if !found {
		return nil, fmt.Errorf("channel '%v' not found", channel.String())
	}
	return channel, nil
}

func (r ChannelInfoResolver) ConvertCrossChainChannel(ctx sdk.Context, calleeXCC XCC, callerXCC XCC) (XCC, error) {
	isLocalCallee := r.IsSelfCrossChainChannel(ctx, calleeXCC)
	isLocalCaller := r.IsSelfCrossChainChannel(ctx, callerXCC)

	if !isLocalCallee && !isLocalCaller {
		return nil, fmt.Errorf("either callee or caller must be self xcc")
	} else if isLocalCaller {
		return calleeXCC, nil
	} else if isLocalCallee {
		callerChannelInfo, err := r.ResolveCrossChainChannel(ctx, callerXCC)
		if err != nil {
			return nil, err
		}
		callerChannel, found := r.channelKeeper.GetChannel(ctx, callerChannelInfo.Port, callerChannelInfo.Channel)
		if !found {
			return nil, fmt.Errorf("channel '%v' not found", callerChannelInfo.String())
		}
		return &ChannelInfo{Port: callerChannel.GetCounterparty().GetPortID(), Channel: callerChannel.GetCounterparty().GetChannelID()}, nil
	} else {
		panic("unreachable")
	}
}

func (ChannelInfoResolver) GetSelfCrossChainChannel(ctx sdk.Context) XCC {
	return &ChannelInfo{}
}

func (r ChannelInfoResolver) IsSelfCrossChainChannel(ctx sdk.Context, xcc XCC) bool {
	return xcc.Equal(r.GetSelfCrossChainChannel(ctx))
}

func (r ChannelInfoResolver) Capabilities() XCCResolverCapabilities {
	return crossChainChannelResolverCapabilities{}
}
