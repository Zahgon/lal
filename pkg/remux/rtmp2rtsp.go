// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package remux

import (
	"math/rand"
	"time"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/rtprtcp"
	"github.com/q191201771/lal/pkg/sdp"
)

// TODO(chef): refactor 将analyze部分独立出来作为一个filter
// TODO(chef): fix 如果前面来的音频和视频数据没有seq header，都是gop中间的数据，那么analyze分析的结果可能是音频和视频都没有

var (
	// config
	// TODO(chef): 提供option，另外还有ssrc和pt都支持自定义
	maxAnalyzeAvMsgSize = 16
)

// Rtmp2RtspRemuxer 提供rtmp数据向sdp+rtp数据的转换
type Rtmp2RtspRemuxer struct {
	onSdp       OnSdp
	onRtpPacket OnRtpPacket

	analyzeDone        bool
	msgCache           []base.RtmpMsg
	vps, sps, pps, asc []byte
	audioPt            base.AvPacketPt
	videoPt            base.AvPacketPt
	audioSampleRate    int

	audioSsrc   uint32
	videoSsrc   uint32
	audioPacker *rtprtcp.RtpPacker
	videoPacker *rtprtcp.RtpPacker
}

type OnSdp func(sdpCtx sdp.LogicContext)
type OnRtpPacket func(pkt rtprtcp.RtpPacket)

// NewRtmp2RtspRemuxer @param onSdp:       每次回调为独立的内存块，回调结束后，内部不再使用该内存块
// @param onRtpPacket: 每次回调为独立的内存块，回调结束后，内部不再使用该内存块
func NewRtmp2RtspRemuxer(onSdp OnSdp, onRtpPacket OnRtpPacket) *Rtmp2RtspRemuxer {
	_ = "STUB: not implemented"
	return nil
}

// FeedRtmpMsg @param msg: 函数调用结束后，内部不持有`msg`内存块
func (r *Rtmp2RtspRemuxer) FeedRtmpMsg(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

// 我们需要先接收一部分rtmp数据，得到音频头、视频头
// 并且考虑，流中只有音频或只有视频的情况
// 我们把前面这个阶段叫做Analyze分析阶段

// 正常阶段

// 音视频头已通过sdp回调，rtp数据中不再包含音视频头
// TODO(chef): [opt] RtspRemuxerAddSpsPps2KeyFrameFlag 开启时，考虑更新sps 202207

func (r *Rtmp2RtspRemuxer) doAnalyze() { _ = "STUB: not implemented"; return }

// aac的采样率以asc为准

// 回调sdp

// 分析阶段缓存的数据

// 是否应该退出Analyze阶段
func (r *Rtmp2RtspRemuxer) isAnalyzeEnough() bool {
	_ = "STUB: not implemented"
	// 音视频头都收集好了
	// 注意，这里故意只判断sps和pps，从而同时支持h264和2h65的情况
	return false
}

// 达到分析包数阈值了

func (r *Rtmp2RtspRemuxer) remux(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

func (r *Rtmp2RtspRemuxer) getAudioPacker() *rtprtcp.RtpPacker {
	_ = "STUB: not implemented"
	return nil

	// TODO(chef): ssrc随机产生，并且整个lal没有在setup信令中传递ssrc
}

func (r *Rtmp2RtspRemuxer) getVideoPacker() *rtprtcp.RtpPacker {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
