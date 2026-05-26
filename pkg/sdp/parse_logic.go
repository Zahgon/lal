// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package sdp

import (
	"github.com/q191201771/lal/pkg/base"
)

type LogicContext struct {
	RawSdp []byte

	AudioClockRate int
	VideoClockRate int

	Asc []byte
	Vps []byte
	Sps []byte
	Pps []byte

	audioPayloadTypeBase base.AvPacketPt // lal内部定义的类型
	videoPayloadTypeBase base.AvPacketPt

	audioPayloadTypeOrigin int // 原始类型，sdp或rtp中的类型
	videoPayloadTypeOrigin int
	audioAControl          string
	videoAControl          string

	// 没有用上的
	hasAudio bool
	hasVideo bool
}

func (lc *LogicContext) IsAudioPayloadTypeOrigin(t int) bool {
	_ = "STUB: not implemented"
	return false
}

func (lc *LogicContext) IsVideoPayloadTypeOrigin(t int) bool {
	_ = "STUB: not implemented"
	return false
}

func (lc *LogicContext) IsPayloadTypeOrigin(t int) bool { _ = "STUB: not implemented"; return false }

func (lc *LogicContext) IsAudioUnpackable() bool { _ = "STUB: not implemented"; return false }

func (lc *LogicContext) IsVideoUnpackable() bool { _ = "STUB: not implemented"; return false }

func (lc *LogicContext) IsAudioUri(uri string) bool { _ = "STUB: not implemented"; return false }

func (lc *LogicContext) IsVideoUri(uri string) bool { _ = "STUB: not implemented"; return false }

func (lc *LogicContext) HasAudioAControl() bool { _ = "STUB: not implemented"; return false }

func (lc *LogicContext) HasVideoAControl() bool { _ = "STUB: not implemented"; return false }

func (lc *LogicContext) MakeAudioSetupUri(uri string) string { _ = "STUB: not implemented"; return "" }

func (lc *LogicContext) MakeVideoSetupUri(uri string) string { _ = "STUB: not implemented"; return "" }

func (lc *LogicContext) GetAudioPayloadTypeBase() base.AvPacketPt {
	_ = "STUB: not implemented"
	return *new(base.AvPacketPt)
}

func (lc *LogicContext) GetVideoPayloadTypeBase() base.AvPacketPt {
	_ = "STUB: not implemented"
	return *new(base.AvPacketPt)
}

func (lc *LogicContext) makeSetupUri(uri string, aControl string) string {
	_ = "STUB: not implemented"
	return ""
}

func ParseSdp2LogicContext(b []byte) (LogicContext, error) {
	_ = "STUB: not implemented"
	return *new(LogicContext), nil
}

// 例子:a=rtpmap:8 PCMA/8000/1
// rtmpmap中有PCMA字段表示G711A

// ffmpeg推流情况下不会填充rtpmap字段,m中pt值为8也可以表示是PCMA,采样率默认为8000Hz
// RFC3551中表明G711A固定pt值为8

// afmtp不存在，也即没法从sdp中解析出sps、pps。
// 这种情况是存在的，sps、pps可以在后续的rtp数据包中传输。
// 所以这里只打印警告。
