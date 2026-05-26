// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

import (
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/naza/pkg/circularqueue"
)

const maxQueueSize = 128

type OnAvPacket func(pkt base.AvPacket)

// AvPacketQueue
//
// 处理音频和视频的时间戳：
// 1. 让音频和视频的时间戳都从0开始（改变原时间戳）
// 2. 让音频和视频的时间戳交替递增输出（不改变原时间戳）
// 3. 当时间戳翻转时，保持时间戳的连续性（改变原时间戳）
//
// 注意，本模块默认音频和视频都存在，如果只有音频或只有视频，则不要使用该模块
//
// TODO(chef): [refactor] 重命名为filter 202305
type AvPacketQueue struct {
	onAvPacket OnAvPacket

	audioQueue *circularqueue.CircularQueue // TODO chef: 特化成AvPacket类型
	videoQueue *circularqueue.CircularQueue

	audioBaseTs int64 // audio base timestamp
	videoBaseTs int64 // video base timestamp

	audioPrevOriginTs   int64
	videoPrevOriginTs   int64
	audioPrevModTs      int64
	videoPrevModTs      int64
	audioPrevIntervalTs int64
	videoPrevIntervalTs int64
}

func NewAvPacketQueue(onAvPacket OnAvPacket) *AvPacketQueue { _ = "STUB: not implemented"; return nil }

func (a *AvPacketQueue) adjustTsHandleRotate(pkt *base.AvPacket) {
	_ = "STUB: not implemented"
	// TODO(chef): [refactor] adjustTsXxx 这一层换成fn这一层 202305
	return
}

// 第一次

// 没有prev，所以没法也不需要计算 prevIntervalTs

// 非第一次

// 时间戳翻滚，并且差值大于阈值了（为了避免B帧导致的小范围翻滚)

// 注意，这个不要放到if判断的前面，因为需要先打印旧值日志

// 用历史差值来更新

func (a *AvPacketQueue) adjustTs(pkt *base.AvPacket) {
	_ = "STUB: not implemented"

	// 时间戳回退了
	return
}

// 第一次

// 根据基准调节

// Feed 注意，调用方保证，音频相较于音频，视频相较于视频，时间戳是线性递增的。
func (a *AvPacketQueue) Feed(pkt base.AvPacket) {
	_ = "STUB: not implemented"
	// Log.Debugf("[AVQ] Feed. t=%d, ts=%d, Q(%d,%d), %s, base(%d,%d)", pkt.PayloadType, pkt.Timestamp, a.audioQueue.Size(), a.videoQueue.Size(), packetsReadable(peekQueuePackets(a)), a.audioBaseTs, a.videoBaseTs)
	return
}

// 如果音频和视频都存在，则按序输出，直到其中一个为空

//Log.Debugf("[AVQ] pop audio by video. a=%d, v=%d", aapkt.Timestamp, vvpkt.Timestamp)

//Log.Debugf("[AVQ] pop video by audio. a=%d, v=%d", aapkt.Timestamp, vvpkt.Timestamp)

// 相等时，我们把早加入的先输出

//Log.Debugf("[AVQ] pop video by audio. a=%d, v=%d", aapkt.Timestamp, vvpkt.Timestamp)

//Log.Debugf("[AVQ] pop audio by video. a=%d, v=%d", aapkt.Timestamp, vvpkt.Timestamp)

// 如果视频满了，则全部输出

// 如果音频满了，则全部输出

func (a *AvPacketQueue) PopAllByForce() { _ = "STUB: not implemented"; return }

// noop

// 这种情况不可能发生
// 因为每次排出时，都会排空一个队列

func (a *AvPacketQueue) popAllAudio() {
	_ = "STUB: not implemented"
	// Log.Debugf("[AVQ] pop all audio. audioQueue=%d, videoQueue=%d", a.audioQueue.Size(), a.videoQueue.Size())
	return
}

func (a *AvPacketQueue) popAllVideo() {
	_ = "STUB: not implemented"
	// Log.Debugf("[AVQ] pop all video. audioQueue=%d, videoQueue=%d", a.audioQueue.Size(), a.videoQueue.Size())
	return
}

// ---------------------------------------------------------------------------------------------------------------------

func packetsReadable(pkts []base.AvPacket) string { _ = "STUB: not implemented"; return "" }

func peekQueuePackets(q *AvPacketQueue) []base.AvPacket { _ = "STUB: not implemented"; return nil }
