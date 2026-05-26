// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtprtcp

import (
	"math/rand"
	"time"

	"github.com/q191201771/lal/pkg/base"
)

type RtpPacker struct {
	payloadPacker IRtpPackerPayload
	clockRate     int
	ssrc          uint32
	option        RtpPackerOption

	seq uint16
}

type RtpPackerOption struct {
	MaxPayloadSize int
	FirstSeq       uint16 // 初始seq，如果不设置，则随机产生
}

var defaultRtpPackerOption = RtpPackerOption{
	MaxPayloadSize: 1200, // TODO(chef) 这个值弄个更合适的
}

type ModRtpPackerOption func(option *RtpPackerOption)

func NewRtpPacker(payloadPacker IRtpPackerPayload, clockRate int, ssrc uint32, modOptions ...ModRtpPackerOption) *RtpPacker {
	_ = "STUB: not implemented"
	return nil
}

// Pack
//
// @param pkt:
//
// - pkt.Timestamp   绝对时间戳，单位毫秒。
// - pkt.PayloadType rtp包头中的packet type。
func (r *RtpPacker) Pack(pkt base.AvPacket) (out []RtpPacket) {
	_ = "STUB: not implemented"
	return nil
}

func (r *RtpPacker) genSeq() (ret uint16) { _ = "STUB: not implemented"; return 0 }

func init() {
	rand.Seed(time.Now().UnixNano())
}
