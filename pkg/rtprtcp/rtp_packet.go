// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtprtcp

import (
	"github.com/q191201771/lal/pkg/base"
)

// -----------------------------------
// rfc3550 5.1 RTP Fixed Header Fields
// -----------------------------------
//
// 0                   1                   2                   3
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |V=2|P|X|  CC   |M|     PT      |       sequence number         |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                           timestamp                           |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |           synchronization source (SSRC) identifier            |
// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
// |            contributing source (CSRC) identifiers             |
// |                             ....                              |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

const (
	RtpFixedHeaderLength = 12

	DefaultRtpVersion = 2
)

const (
	PositionTypeSingle    uint8 = 1
	PositionTypeFuaStart  uint8 = 2
	PositionTypeFuaMiddle uint8 = 3
	PositionTypeFuaEnd    uint8 = 4
	PositionTypeStapa     uint8 = 5 // 1个rtp包包含多个帧，目前供h264的stapa使用
	PositionTypeAp        uint8 = 6 // 1个rtp包包含多个帧，目前供h265的ap使用
)

type RtpHeader struct {
	Version    uint8  // 2b  *
	Padding    uint8  // 1b
	Extension  uint8  // 1
	CsrcCount  uint8  // 4b
	Mark       uint8  // 1b  *
	PacketType uint8  // 7b
	Seq        uint16 // 16b **
	Timestamp  uint32 // 32b **** samples
	Ssrc       uint32 // 32b **** Synchronization source

	Csrc []uint32

	ExtensionProfile uint16

	// Extensions 包含了整个extension，引用的是包体的内存
	//
	// TODO(chef): [opt] 后续考虑解析extension中的单独个item存储至结构体中 202211
	Extensions []byte

	payloadOffset uint32 // body部分，真正数据部分的起始位置
	paddingLength int    // 末尾padding的长度
}

type RtpPacket struct {
	Header RtpHeader
	Raw    []byte // 包含header内存

	positionType uint8
}

func (h *RtpHeader) PackTo(out []byte) { _ = "STUB: not implemented"; return }

// TODO(chef): pack csrc 202210

func MakeDefaultRtpHeader() RtpHeader { _ = "STUB: not implemented"; return *new(RtpHeader) }

func MakeRtpPacket(h RtpHeader, payload []byte) (pkt RtpPacket) {
	_ = "STUB: not implemented"
	return *new(RtpPacket)
}

func ParseRtpHeader(b []byte) (h RtpHeader, err error) {
	_ = "STUB: not implemented"
	return *new(RtpHeader), nil
}

// TODO 按照RFC-5285-RTP-Header-Extensions去解析

// rfc3550#section-5.3.1

// ParseRtpPacket 函数调用结束后，不持有参数<b>的内存块
func ParseRtpPacket(b []byte) (pkt RtpPacket, err error) {
	_ = "STUB: not implemented"
	return *new(RtpPacket), nil
}

func (p *RtpPacket) Body() []byte { _ = "STUB: not implemented"; return nil }

func (p *RtpPacket) DebugString() string { _ = "STUB: not implemented"; return "" }

// IsAvcHevcBoundary
//
// @param pt: 取值范围为AvPacketPtAvc或AvPacketPtHevc，否则直接返回false
func IsAvcHevcBoundary(pkt RtpPacket, pt base.AvPacketPt) bool {
	_ = "STUB: not implemented"
	return false
}

func IsAvcBoundary(pkt RtpPacket) bool { _ = "STUB: not implemented"; return false }

// TODO(chef): [fix] 检查数据长度有效性 202211

func IsHevcBoundary(pkt RtpPacket) bool { _ = "STUB: not implemented"; return false }

// TODO(chef): [fix] 检查数据长度有效性 202211

// 注意，这里是后6位，不是中间6位
