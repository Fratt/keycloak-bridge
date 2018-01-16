// automatically generated by the FlatBuffers compiler, do not modify

package events

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AuthDetails struct {
	_tab flatbuffers.Table
}

func GetRootAsAuthDetails(buf []byte, offset flatbuffers.UOffsetT) *AuthDetails {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AuthDetails{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *AuthDetails) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AuthDetails) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AuthDetails) RealmId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AuthDetails) ClientId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AuthDetails) UserId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AuthDetails) IpAddress() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func AuthDetailsStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func AuthDetailsAddRealmId(builder *flatbuffers.Builder, realmId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(realmId), 0)
}
func AuthDetailsAddClientId(builder *flatbuffers.Builder, clientId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(clientId), 0)
}
func AuthDetailsAddUserId(builder *flatbuffers.Builder, userId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(userId), 0)
}
func AuthDetailsAddIpAddress(builder *flatbuffers.Builder, ipAddress flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(ipAddress), 0)
}
func AuthDetailsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}