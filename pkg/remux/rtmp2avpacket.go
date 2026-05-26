// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package remux

import (
	"github.com/q191201771/lal/pkg/base"
)

// TODO(chef): 该文件处于开发阶段，请不要直接使用
// TODO(chef): 支持音频 202206

// Rtmp2AvPacketRemuxer
//
// 用途：
// - 将rtmp流中的视频转换成ffmpeg可解码的格式
type Rtmp2AvPacketRemuxer struct {
	option     Rtmp2AvPacketRemuxerOption
	onAvPacket func(pkt base.AvPacket, arg interface{})

	spspps []byte // annexb格式
}

type Rtmp2AvPacketRemuxerOption struct {
	// TODO(chef): impl me 202206
	TryInPlaceFlag bool // 尝试在原有内存上直接修改
	EraseSeiFlag   bool
}

var defaultRtmp2AvPacketRemuxerOption = Rtmp2AvPacketRemuxerOption{
	TryInPlaceFlag: false,
	EraseSeiFlag:   true,
}

func NewRtmp2AvPacketRemuxer() *Rtmp2AvPacketRemuxer { _ = "STUB: not implemented"; return nil }

func (r *Rtmp2AvPacketRemuxer) WithOption(modOption func(option *Rtmp2AvPacketRemuxerOption)) *Rtmp2AvPacketRemuxer {
	_ = "STUB: not implemented"
	return nil
}

// WithOnAvPacket
//
// @param onAvPacket: pkt 内存由内部新申请，回调后内部不再使用
func (r *Rtmp2AvPacketRemuxer) WithOnAvPacket(onAvPacket func(pkt base.AvPacket, arg interface{})) *Rtmp2AvPacketRemuxer {
	_ = "STUB: not implemented"
	return nil
}

func (r *Rtmp2AvPacketRemuxer) FeedRtmpMsg(msg base.RtmpMsg, arg interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

func (r *Rtmp2AvPacketRemuxer) feedVideo(msg base.RtmpMsg, arg interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// noop

// noop

// ---------------------------------------------------------------------------------------------------------------------

func defaultOnAvPacket(pkt base.AvPacket, arg interface{}) {
	_ = "STUB: not implemented"
	// noop
	return
}
