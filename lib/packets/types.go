//
//
// @filename: packets/types.go
// @author: Krisna Pranav, NanoBlocks Developers
//
//

package packets

import (
	"github.com/NanoOfficial/smartnano/lib/utils"
	sptypes "github.com/bluele/interchain-simple-packet/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/gogo/protobuf/proto"
)

type (
	Header                    = sptypes.Header
	PacketData                = sptypes.PacketData
	PacketAcknowledgementData = PacketData
)

func NewPacketData(h *Header, payload PacketDataPayload) PacketData {
	if h == nil {
		h = &Header{}
	}
	any, err := utils.PackAny(payload)
	if err != nil {
		panic(err)
	}
	return sptypes.NewSimplePacketData(*h, any)
}

type PacketDataPayload interface {
	proto.Message
	Type() string
	ValidateBasic() error
}

type PacketAcknowledgementPayload interface {
	proto.Message
	Type() string
	ValidateBasic() error
}

type IncomingPacket interface {
	exported.PacketI
	PacketData() PacketData
	Header() Header
	Payload() PacketDataPayload
}

func UnmarshalIncomingPacket(m codec.Codec, raw exported.PacketI) (IncomingPacket, error) {
	var pd PacketData
	var payload PacketDataPayload
	if err := proto.Unmarshal(raw.GetData(), &pd); err != nil {
		return nil, err
	}
	if err := m.UnpackAny(pd.GetPayload(), &payload); err != nil {
		return nil, err
	}
	return NewIncomingPacket(raw, pd, payload), nil
}

var _ IncomingPacket = (*incomingPacket)(nil)

type incomingPacket struct {
	exported.PacketI
	packetData PacketData
	payload    PacketDataPayload
}

func NewIncomingPacket(raw exported.PacketI, packetData PacketData, payload PacketDataPayload) IncomingPacket {
	return &incomingPacket{
		PacketI:    raw,
		packetData: packetData,
		payload:    payload,
	}
}

func (p incomingPacket) PacketData() PacketData {
	return p.packetData
}

func (p incomingPacket) Header() Header {
	return p.packetData.Header
}

func (p incomingPacket) Payload() PacketDataPayload {
	return p.payload
}

type OutgoingPacket interface {
	IncomingPacket
	SetPacketData(header Header, payload PacketDataPayload)
}

var _ OutgoingPacket = (*outgoingPacket)(nil)

type outgoingPacket struct {
	exported.PacketI
	packetData PacketData
	payload    PacketDataPayload
}

func NewOutgoingPacket(raw exported.PacketI, packetData PacketData, payload PacketDataPayload) OutgoingPacket {
	return &outgoingPacket{
		PacketI:    raw,
		packetData: packetData,
		payload:    payload,
	}
}

func (p outgoingPacket) PacketData() PacketData {
	return p.packetData
}

func (p outgoingPacket) Header() Header {
	return p.packetData.Header
}

func (p outgoingPacket) Payload() PacketDataPayload {
	return p.payload
}

func (p *outgoingPacket) SetPacketData(header Header, payload PacketDataPayload) {
	p.payload = payload
	p.packetData = NewPacketData(&header, payload)
}

func (p outgoingPacket) GetData() []byte {
	bz, err := proto.Marshal(&p.packetData)
	if err != nil {
		panic(err)
	}
	return bz
}

func NewPacketAcknowledgementData(h *Header, payload PacketAcknowledgementPayload) PacketAcknowledgementData {
	if h == nil {
		h = new(Header)
	}
	any, err := utils.PackAny(payload)
	if err != nil {
		panic(err)
	}
	return PacketAcknowledgementData{
		Header:  *h,
		Payload: any,
	}
}

type IncomingPacketAcknowledgement interface {
	Data() PacketAcknowledgementData
	Header() Header
	Payload() PacketAcknowledgementPayload
}

type incomingPacketAcknowledgement struct {
	data    PacketAcknowledgementData
	payload PacketAcknowledgementPayload
}

var _ IncomingPacketAcknowledgement = (*incomingPacketAcknowledgement)(nil)

func NewIncomingPacketAcknowledgement(h *Header, payload PacketAcknowledgementPayload) IncomingPacketAcknowledgement {
	return incomingPacketAcknowledgement{data: NewPacketAcknowledgementData(h, payload), payload: payload}
}

func (a incomingPacketAcknowledgement) Data() PacketAcknowledgementData {
	return a.data
}

func (a incomingPacketAcknowledgement) Header() Header {
	return a.data.Header
}

func (a incomingPacketAcknowledgement) Payload() PacketAcknowledgementPayload {
	return a.payload
}

func UnmarshalIncomingPacketAcknowledgement(m codec.Codec, ack []byte) (IncomingPacketAcknowledgement, error) {
	var pd PacketAcknowledgementData
	var payload PacketAcknowledgementPayload

	if err := proto.Unmarshal(ack, &pd); err != nil {
		return nil, err
	}
	if err := m.UnpackAny(pd.GetPayload(), &payload); err != nil {
		return nil, err
	}

	return NewIncomingPacketAcknowledgement(&pd.Header, payload), nil
}

type OutgoingPacketAcknowledgement interface {
	IncomingPacketAcknowledgement
	SetData(header Header, payload PacketAcknowledgementPayload)
}

type outgoingPacketAcknowledgement struct {
	data    PacketAcknowledgementData
	payload PacketAcknowledgementPayload
}

func NewOutgoingPacketAcknowledgement(h *Header, payload PacketAcknowledgementPayload) OutgoingPacketAcknowledgement {
	return &outgoingPacketAcknowledgement{
		data:    NewPacketAcknowledgementData(h, payload),
		payload: payload,
	}
}

var _ OutgoingPacketAcknowledgement = (*outgoingPacketAcknowledgement)(nil)

func (a outgoingPacketAcknowledgement) Data() PacketAcknowledgementData {
	return a.data
}

func (a outgoingPacketAcknowledgement) Header() Header {
	return a.data.Header
}

func (a outgoingPacketAcknowledgement) Payload() PacketAcknowledgementPayload {
	return a.payload
}

func (a outgoingPacketAcknowledgement) SetData(header Header, payload PacketAcknowledgementPayload) {
	a.data = NewPacketAcknowledgementData(&header, payload)
	a.payload = payload
}
