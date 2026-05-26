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

type RtpUnpackerAvcHevc struct {
	payloadType base.AvPacketPt
	clockRate   int
	onAvPacket  OnAvPacket
}

func NewRtpUnpackerAvcHevc(payloadType base.AvPacketPt, clockRate int, onAvPacket OnAvPacket) *RtpUnpackerAvcHevc {
	_ = "STUB: not implemented"
	return nil
}

func (unpacker *RtpUnpackerAvcHevc) CalcPositionIfNeeded(pkt *RtpPacket) {
	_ = "STUB: not implemented"
	return
}

func (unpacker *RtpUnpackerAvcHevc) TryUnpackOne(list *RtpPacketList) (unpackedFlag bool, unpackedSeq uint16) {
	_ = "STUB: not implemented"
	return false, 0
}

// 跳过前面的字节，并且将多nalu前的2字节长度，替换成4字节长度
// skip后：
// rtp中的数据格式 [<2字节的nalu长度>, <nalu>, <2字节的nalu长度>, <nalu> ...]
// 转变后的数据格式 [<4字节的nalu长度>, <nalu>, <4字节的nalu长度>, <nalu> ...]

// 使用两次遍历，第一次遍历找出总大小，第二次逐个拷贝，目的是使得内存块一次就申请好，不用动态扩容造成额外性能开销

// ffmpeg rtpdec_hevc.c
// 取buf[0]的头尾各1位

// 使用两次遍历，第一次遍历找出总大小，第二次逐个拷贝，目的是使得内存块一次就申请好，不用动态扩容造成额外性能开销

// naluTypeLen表示的是合帧之后的nalu type的长度。而+1，是在rtp时头的长度。
// 比如h264，帧数据时naluTypeLen是1字节，rtp包时长度是2字节。

// 三部分组成： len + type + data

// len

// type

// data

// 不应该出现其他类型

// noop

// noop

func calcPositionIfNeededAvc(pkt *RtpPacket) {
	_ = "STUB: not implemented"

	// rfc3984 5.3.  NAL Unit Octet Usage
	//
	// +---------------+
	// |0|1|2|3|4|5|6|7|
	// +-+-+-+-+-+-+-+-+
	// |F|NRI|  Type   |
	// +---------------+
	return
}

// rfc3984 5.8.  Fragmentation Units (FUs)
//
// 0                   1                   2                   3
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// | FU indicator  |   FU header   |                               |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+                               |
// |                                                               |
// |                         FU payload                            |
// |                                                               |
// |                               +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                               :...OPTIONAL RTP padding        |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//
// FU indicator:
// +---------------+
// |0|1|2|3|4|5|6|7|
// +-+-+-+-+-+-+-+-+
// |F|NRI|  Type   |
// +---------------+
//
// Fu header:
// +---------------+
// |0|1|2|3|4|5|6|7|
// +-+-+-+-+-+-+-+-+
// |S|E|R|  Type   |
// +---------------+

func calcPositionIfNeededHevc(pkt *RtpPacket) {
	_ = "STUB: not implemented"

	// +---------------+---------------+
	// |0|1|2|3|4|5|6|7|0|1|2|3|4|5|6|7|
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// |F|   Type    |  LayerId  | TID |
	// +-------------+-----------------+
	return
}

// Figure 1: The Structure of the HEVC NAL Unit Header

// 0                   1                   2                   3
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |    PayloadHdr (Type=49)       |   FU header   | DONL (cond)   |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-|
// | DONL (cond)   |                                               |
// |-+-+-+-+-+-+-+-+                                               |
// |                         FU payload                            |
// |                                                               |
// |                               +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                               :...OPTIONAL RTP padding        |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

// Figure 9: The Structure of an FU

// +---------------+
// |0|1|2|3|4|5|6|7|
// +-+-+-+-+-+-+-+-+
// |S|E|  FuType   |
// +---------------+

// Figure 10: The Structure of FU Header
