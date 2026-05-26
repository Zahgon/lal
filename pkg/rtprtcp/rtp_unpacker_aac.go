// Copyright 2020, Chef.  All rights reserved.
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

type RtpUnpackerAac struct {
	payloadType base.AvPacketPt
	clockRate   int
	onAvPacket  OnAvPacket
}

func NewRtpUnpackerAac(payloadType base.AvPacketPt, clockRate int, onAvPacket OnAvPacket) *RtpUnpackerAac {
	_ = "STUB: not implemented"
	return nil
}

func (unpacker *RtpUnpackerAac) CalcPositionIfNeeded(pkt *RtpPacket) {
	_ = "STUB: not implemented"
	// noop
	return
}

func (unpacker *RtpUnpackerAac) TryUnpackOne(list *RtpPacketList) (unpackedFlag bool, unpackedSeq uint16) {
	_ = "STUB: not implemented"
	// rfc3640 2.11.  Global Structure of Payload Format
	//
	// +---------+-----------+-----------+---------------+
	// | RTP     | AU Header | Auxiliary | Access Unit   |
	// | Header  | Section   | Section   | Data Section  |
	// +---------+-----------+-----------+---------------+
	//
	//           <----------RTP Packet Payload----------->
	//
	// rfc3640 3.2.1.  The AU Header Section
	//
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+- .. -+-+-+-+-+-+-+-+-+-+
	// |AU-headers-length|AU-header|AU-header|      |AU-header|padding|
	// |                 |   (1)   |   (2)   |      |   (n)   | bits  |
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+- .. -+-+-+-+-+-+-+-+-+-+
	//
	// rfc3640 3.3.6.  High Bit-rate AAC
	//
	// rtp_parse_mp4_au()
	//
	//
	// 3.2.3.1.  Fragmentation
	//
	//   A packet SHALL carry either one or more complete Access Units, or a
	//   single fragment of an Access Unit.  Fragments of the same Access Unit
	//   have the same time stamp but different RTP sequence numbers.  The
	//   marker bit in the RTP header is 1 on the last fragment of an Access
	//   Unit, and 0 on all other fragments.
	//
	return false, 0
}

// first

// 只有一个描述

// 描述的音频帧完整的在当前的rtp packet中，没有跨越到下个rtp packet

// one complete access unit

// fragmented
// 注意，这里我们参考size和rtp包头中的timestamp，不参考rtp包头中的mark位

// 注意，非第一个fragment，也会包含au，au的size和第一个fragment里au的size应该相等

// can reach here

// more complete access unit

// TODO chef: 这里1024的含义

type au struct {
	size uint32 // 该音频帧的大小
	pos  uint32 // 相对rtp body的位置
}

func parseAu(b []byte) (ret []au) {
	_ = "STUB: not implemented"
	// TODO(chef): [fix] 解析b时，没有判断长度有效性 202207
	return nil
}

// AU Header Section

// TODO chef: 这里的2是写死的，正常是外部传入auSize和auIndex所占位数的和

// 有多少个AU-Header

// AU Header pos
// AU pos

// TODO chef: auSize和auIndex所在的位数是写死的13bit，3bit，标准的做法应该从外部传入，比如从sdp中获取后传入
// 13bit

// 注意，fragment时，auIndex并不可靠。见TestAacCase1
//auIndex := b[pauh+1] & 0x7
//Log.Debugf("~ %d %d", auSize, auIndex)
