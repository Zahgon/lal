// Copyright 2021, Chef.  All rights reserved.
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

type VideoInfo struct {
	VideoPt       base.AvPacketPt
	Vps, Sps, Pps []byte
}

type AudioInfo struct {
	AudioPt           base.AvPacketPt
	SamplingFrequency int
	Asc               []byte
}

func Pack(videoInfo VideoInfo, audioInfo AudioInfo) (ctx LogicContext, err error) {
	_ = "STUB: not implemented"
	// 组装SDP头部
	return *new(LogicContext), nil
}

// 组装视频SDP信息

// 组装音频SDP信息

func buildVideoSdpInfo(videoInfo VideoInfo, streamid int) string {
	_ = "STUB: not implemented"
	return ""
}

func buildAudioSdpInfo(audioInfo AudioInfo, streamid int) string {
	_ = "STUB: not implemented"
	return ""
}
