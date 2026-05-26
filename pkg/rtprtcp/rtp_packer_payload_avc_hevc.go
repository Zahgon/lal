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

type RtpPackerPayloadAvcHevcType int

const (
	RtpPackerPayloadAvcHevcTypeNalu   RtpPackerPayloadAvcHevcType = 1
	RtpPackerPayloadAvcHevcTypeAvcc                               = 2
	RtpPackerPayloadAvcHevcTypeAnnexb                             = 3
)

type RtpPackerPayloadAvcHevcOption struct {
	Typ RtpPackerPayloadAvcHevcType
}

var defaultRtpPackerPayloadAvcHevcOption = RtpPackerPayloadAvcHevcOption{
	Typ: RtpPackerPayloadAvcHevcTypeNalu,
}

type RtpPackerPayloadAvcHevc struct {
	payloadType base.AvPacketPt
	option      RtpPackerPayloadAvcHevcOption
}

type ModRtpPackerPayloadAvcHevcOption func(option *RtpPackerPayloadAvcHevcOption)

func NewRtpPackerPayloadAvc(modOptions ...ModRtpPackerPayloadAvcHevcOption) *RtpPackerPayloadAvcHevc {
	_ = "STUB: not implemented"
	return nil
}

func NewRtpPackerPayloadHevc(modOptions ...ModRtpPackerPayloadAvcHevcOption) *RtpPackerPayloadAvcHevc {
	_ = "STUB: not implemented"
	return nil
}

func NewRtpPackerPayloadAvcHevc(payloadType base.AvPacketPt, modOptions ...ModRtpPackerPayloadAvcHevcOption) *RtpPackerPayloadAvcHevc {
	_ = "STUB: not implemented"
	return nil
}

// Pack @param in: AVCC格式
//
// @return out: 内存块为独立新申请；函数返回后，内部不再持有该内存块
func (r *RtpPackerPayloadAvcHevc) Pack(in []byte, maxSize int) (out [][]byte) {
	_ = "STUB: not implemented"
	return nil
}

// RtpPackerPayloadAvcHevcTypeNalu

func (r *RtpPackerPayloadAvcHevc) PackNal(nal []byte, maxSize int) (out [][]byte) {
	_ = "STUB: not implemented"
	// pack逻辑
	//
	// avc
	//
	// 输入
	// nri     [01, 02]
	// nalType [03, 07]
	//
	// 输出
	// nri     [01, 02]
	// 28      [03, 07]    28是avc fua的nal type
	// start   [10]
	// end     [11]
	// nalType [13, 17]
	//
	// hevc
	//
	// 输入
	// nalType [01, 06]
	//
	// 输出
	// 49      [01, 06] 49是hevc fua的nal type
	// 1       [10, 17]
	// start   [20]
	// end     [21]
	// nalType [22, 27] 注意，和输入的nalType的所在type字节的位位置不同
	//
	return nil
}

// single

// FU-A

// var
// rtp payload大小

// const after set

// start-end标志所在位置

// only avc

// 注意，跳过输入的nal type那个字节，使用FU-A自己的两个字节的头，避免重复

// 前面的包

// ffmpeg, rtpenc_h264_hevc.c, func nal_send

// 当前帧切割后的首个RTP包

// start

//

// 最后一包

// end
