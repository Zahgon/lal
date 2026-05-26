// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package base

// ---------------------------------------------------------------------------------------------------------------------

type AvPacketPt int

const (
	AvPacketPtUnknown AvPacketPt = -1
	AvPacketPtG711U   AvPacketPt = 0   // g711u
	AvPacketPtG711A   AvPacketPt = 8   // g711a
	AvPacketPtMp2     AvPacketPt = 14  // mp2
	AvPacketPtAvc     AvPacketPt = 96  // h264
	AvPacketPtHevc    AvPacketPt = 98  // h265
	AvPacketPtAac     AvPacketPt = 97  // aac
	AvPacketPtOpus    AvPacketPt = 101 // opus
)

func (a AvPacketPt) ReadableString() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------

// AvPacket
//
// 不同场景使用时，字段含义可能不同。
// 使用 AvPacket 的地方，应注明各字段的含义。
type AvPacket struct {
	PayloadType AvPacketPt
	Timestamp   int64 // 如无特殊说明，此字段是Dts
	Pts         int64
	Payload     []byte
}

func (packet *AvPacket) IsAudio() bool { _ = "STUB: not implemented"; return false }

func (packet *AvPacket) IsVideo() bool { _ = "STUB: not implemented"; return false }

func (packet *AvPacket) DebugString() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------

type OnAvPacketFunc func(packet *AvPacket)
