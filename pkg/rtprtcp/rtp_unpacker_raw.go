// Copyright 2023, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtprtcp

import "github.com/q191201771/lal/pkg/base"

type RtpUnpackerRaw struct {
	payloadType base.AvPacketPt
	clockRate   int
	onAvPacket  OnAvPacket
}

func NewRtpUnpackerRaw(payloadType base.AvPacketPt, clockRate int, onAvPacket OnAvPacket) *RtpUnpackerRaw {
	_ = "STUB: not implemented"
	return nil
}

func (unpacker *RtpUnpackerRaw) CalcPositionIfNeeded(pkt *RtpPacket) {
	_ = "STUB: not implemented"
	// noop
	return
}

func (unpacker *RtpUnpackerRaw) TryUnpackOne(list *RtpPacketList) (unpackedFlag bool, unpackedSeq uint16) {
	_ = "STUB: not implemented"
	// first
	return false, 0
}

// 暂时认为一个rtp为一帧数据(G711A/G711U)
