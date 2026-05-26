// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package remux

import (
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/rtmp"
)

const (
	dummyAudioFilterStageAnalysis = 1
	dummyAudioFilterStageNormal   = 2
	dummyAudioFilterStageDummy    = 3
)

type DummyAudioFilter struct {
	uk          string
	waitAudioMs int
	onPop       rtmp.OnReadRtmpAvMsg

	stage           int
	earlyStageQueue []base.RtmpMsg
	firstVideoTs    uint32
	prevAudioTs     uint32

	audioCount int
}

// NewDummyAudioFilter 检测输入的rtmp流中是否有音频，如果有，则原样返回；如果没有，则制造静音音频数据叠加在rtmp流里面
//
// @param waitAudioMs 等待音频数据时间，如果超出这个时间还没有接收到音频数据，则开始制造静音数据
// @param onPop       注意，所有回调都发生在输入函数调用中
func NewDummyAudioFilter(uk string, waitAudioMs int, onPop rtmp.OnReadRtmpAvMsg) *DummyAudioFilter {
	_ = "STUB: not implemented"
	return nil
}

func (filter *DummyAudioFilter) OnReadRtmpAvMsg(msg base.RtmpMsg) {
	_ = "STUB: not implemented"
	return
}

func (filter *DummyAudioFilter) Feed(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

// 初始阶段，分析是否存在音频
func (filter *DummyAudioFilter) handleAnalysisStage(msg base.RtmpMsg) {
	_ = "STUB: not implemented"
	return
}

// metadata直接入队列

// 原始流中存在音频，将所有缓存数据出队列，进入normal stage

// 分析视频数据累计时长是否达到阈值

// 注意，为了避免seq header的时间戳和视频帧不是线性的（0或其他特殊的值）我们直接入队列并跳过

// 记录首个视频帧的时间戳

// 没有达到阈值

// 达到阈值

// 原始流中存在音频
func (filter *DummyAudioFilter) handleNormalStage(msg base.RtmpMsg) {
	_ = "STUB: not implemented"
	return

	// 原始流中不存在音频
}

func (filter *DummyAudioFilter) handleDummyStage(msg base.RtmpMsg) {
	_ = "STUB: not implemented"
	return
}

// 由于我们已经开始制造静音包了，静音包的编码参数可能会和实际音频参数不一致，所以我们只能过滤掉原始的音频数据了

// TODO(chef): 这里的时间戳可以考虑减1，但是注意处理一些边界条件

func (filter *DummyAudioFilter) cache(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

func (filter *DummyAudioFilter) clearCache() { _ = "STUB: not implemented"; return }

func (filter *DummyAudioFilter) onPopProxy(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

func (filter *DummyAudioFilter) makeAudioSeqHeader(ts uint32) base.RtmpMsg {
	_ = "STUB: not implemented"
	// aac (LC), 48000 Hz, stereo, fltp
	return *new(base.RtmpMsg)
}

func (filter *DummyAudioFilter) makeOneAudio(ts uint32) base.RtmpMsg {
	_ = "STUB: not implemented"
	return *new(base.RtmpMsg)
}

// 注意，前面2字节是seq header头部信息，后面6个字节是AAC静音包

func (filter *DummyAudioFilter) calcAudioDurationMs() uint32 { _ = "STUB: not implemented"; return 0 }
