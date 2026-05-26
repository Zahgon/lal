// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package remux

import (
	"github.com/q191201771/lal/pkg/aac"
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/mpegts"
)

const (
	initialVideoOutBufferSize          = 1024 * 1024
	calcFragmentHeaderQueueSize        = 16
	maxAudioCacheDelayByAudio   uint64 = 150 * 90 // 单位（毫秒*90）
	maxAudioCacheDelayByVideo   uint64 = 300 * 90 // 单位（毫秒*90）
)

type IRtmp2MpegtsRemuxerObserver interface {
	// OnPatPmt
	//
	// @param b: const只读内存块，上层可以持有，但是不允许修改
	//
	OnPatPmt(b []byte)

	// OnTsPackets
	//
	// @param tsPackets:
	//  - mpegts数据，有一个或多个188字节的ts数据组成
	//  - 回调结束后，remux.Rtmp2MpegtsRemuxer 不再使用这块内存块
	//
	// @param frame: 各字段含义见 mpegts.Frame 结构体定义
	//
	// @param boundary: 是否到达边界处，根据该字段，上层可判断诸如：
	//  - hls是否允许开启新的ts文件切片
	//  - 新的ts播放流从该位置开始播放
	//
	OnTsPackets(tsPackets []byte, frame *mpegts.Frame, boundary bool)
}

// Rtmp2MpegtsRemuxer 输入rtmp流，输出mpegts流
type Rtmp2MpegtsRemuxer struct {
	uk string

	observer        IRtmp2MpegtsRemuxerObserver
	filter          *rtmp2MpegtsFilter
	videoOut        []byte // Annexb
	spspps          []byte // Annexb 也可能是vps+sps+pps
	ascCtx          *aac.AscContext
	audioCc         uint8
	videoCc         uint8
	timestampFilter Rtmp2MpegtsTimestampFilter

	// audioCacheFrames: 缓存音频packet数据，注意，可能包含多个音频packet
	//
	// audioCacheFirstFramePts: audioCacheFrames中第一个音频packet的时间戳
	//
	// 缓存（也即合并）多个packet并让他们使用第一个packet的时间戳，目的是为了还原单packet的精度。
	//
	// 比如，aac 48000的每帧数据时长为21.333，
	// 由于rtmp的时间戳是整数（单位毫秒），所以它使用如下方式：
	//  duration: 21, 21, 22
	// timestamp: 0, 21, 42, 64
	// 其中duration的第一个和第二个都少了0.33，第三个多了0.66，然后循环
	// timestamp中0和64是对的，21和42是有误差的，然后循环。也即3个packet，timestamp恢复严格意义上的正确。
	// 这个用整数类型毫秒单位没法避免。
	//
	// 当 * 90换算成mpegts的时间戳时，duration的小数部分的误差还是存在，
	// 这将导致部分hls播放器（比如macOS的QuickTime）会出现声音效果不好的情况。
	// 所以我们优化为将多个packet合并，使得原来每个packet之间出现异常的情况，缓解为缓存与缓存间可能出现异常的情况
	//
	// TODO(chef): 可按ffmpeg中av_rescale_delta的手段进行优化 202206
	//
	// 其大致思想是，当duration在一定范围内，使用timestamp*90；当duration超过一定范围，则直接用采样率换算得到的duration。
	// 也即，迁移中方法是21或者22 * 90，后一种直接是21.33*90=1920
	// int64_t av_rescale_delta(AVRational in_tb, int64_t in_ts,  AVRational fs_tb, int duration, int64_t *last, AVRational out_tb)
	// in_tb    1000
	// in_ts    dts
	// fs_tb    48000
	// duration 1024
	// last
	// out_tb   90000
	//
	// 但是ffmpeg的做法也有问题，就是每个packet都转ts，padding会比较多，导致ts码流变大
	//
	//TODO chef: rename to DTS
	//
	audioCacheFrames        []byte
	audioCacheFirstFramePts uint64

	opened bool
}

func NewRtmp2MpegtsRemuxer(observer IRtmp2MpegtsRemuxerObserver) *Rtmp2MpegtsRemuxer {
	_ = "STUB: not implemented"
	return nil
}

// FeedRtmpMessage
//
// @param msg: msg.Payload 调用结束后，函数内部不会持有这块内存
func (s *Rtmp2MpegtsRemuxer) FeedRtmpMessage(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

func (s *Rtmp2MpegtsRemuxer) Dispose() {
	_ = "STUB: not implemented"

	// ---------------------------------------------------------------------------------------------------------------------
	return
}

// FlushAudio
//
// 吐出音频数据的三种情况：
// 1. 收到音频或视频时，音频缓存队列已达到一定长度（内部判断）
// 2. 打开一个新的TS文件切片时
// 3. 输入流关闭时
func (s *Rtmp2MpegtsRemuxer) FlushAudio() { _ = "STUB: not implemented"; return }

// 注意，在回调前设置为空，因为回调中有可能再次调用FlushAudio

// 回调结束后更新cc

func (s *Rtmp2MpegtsRemuxer) UniqueKey() string {
	_ = "STUB: not implemented"

	// ----- implement of iRtmp2MpegtsFilterObserver ----------------------------------------------------------------------------------------------------------------
	return ""
}

// onPatPmt onPop
//
// 实现 iRtmp2MpegtsFilterObserver
func (s *Rtmp2MpegtsRemuxer) onPatPmt(b []byte) { _ = "STUB: not implemented"; return }

func (s *Rtmp2MpegtsRemuxer) onPop(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

func (s *Rtmp2MpegtsRemuxer) feedVideo(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

// 将数据转换成Annexb

// 如果是seq header sps pps，缓存住，然后直接返回

// msg中可能有多个NALU，逐个获取

// 处理 sps pps aud
//
// aud 过滤掉，我们有自己的添加aud的逻辑
//
// sps pps
// 注意，有的流，seq header中的sps和pps是错误的，需要从nals里获取sps pps并更新
// 见 https://github.com/q191201771/lal/issues/143
//
// TODO(chef): rtmp转其他类型的模块也存在这个问题，应该抽象出一个统一处理的地方
//

// TODO(chef): [opt] 考虑缓存下来，放到关键帧前面 202208

// tag中的首个nalu前面写入aud

// 注意，因为前面已经过滤了sps pps aud的信息，所以这里可以认为都是需要用aud分隔的，不需要单独判断了
//if codecId == base.RtmpCodecIdAvc && (nalType == avc.NaluTypeSei || nalType == avc.NaluTypeIdrSlice || nalType == avc.NaluTypeSlice) {

// 关键帧前追加sps pps

// h264的逻辑，一个tag中，多个连续的关键帧只追加一个，不连续则每个关键帧前都追加。为什么要这样处理

// 这里只有P帧，没有SEI。为什么要这样处理

// TODO(chef): [refactor] avc和hevc可以考虑再抽象一层高层的包，使得更上层代码简洁一些

// 这里简化了，只要不是关键帧，就刷新标志

// 如果写入了aud或spspps，则用start code3，否则start code4。为什么要这样处理
// 这里不知为什么要区分写入两种类型的start code

// for loop

// 比如只有SEI nal或者非seq header的msg只有sps, pps, vps

func (s *Rtmp2MpegtsRemuxer) feedAudio(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

//Log.Debugf("[%s] hls: feedAudio. dts=%d len=%d", s.uk, msg.Header.TimestampAbs, len(msg.Payload))

func (s *Rtmp2MpegtsRemuxer) cacheAacSeqHeader(msg base.RtmpMsg) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Rtmp2MpegtsRemuxer) audioSeqHeaderCached() bool { _ = "STUB: not implemented"; return false }

func (s *Rtmp2MpegtsRemuxer) appendSpsPps(out []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Rtmp2MpegtsRemuxer) videoSeqHeaderCached() bool { _ = "STUB: not implemented"; return false }

func (s *Rtmp2MpegtsRemuxer) audioCacheEmpty() bool { _ = "STUB: not implemented"; return false }

func (s *Rtmp2MpegtsRemuxer) resetAudioCache() { _ = "STUB: not implemented"; return }

func (s *Rtmp2MpegtsRemuxer) resetVideoOutBuffer() { _ = "STUB: not implemented"; return }

func (s *Rtmp2MpegtsRemuxer) onFrame(frame *mpegts.Frame) {
	_ = "STUB: not implemented"
	// Log.Debugf("in frame=%s", frame.DebugString())
	return
}

//Log.Debugf("ou frame=%s", frame.DebugString())

// 为了考虑没有视频的情况也能切片，所以这里判断spspps为空时，也建议生成fragment

// 收到视频，可能触发建立fragment的条件是：
// 关键帧数据 &&
// (
//  (没有收到过音频seq header) || 说明 只有视频
//  (收到过音频seq header && fragment没有打开) || 说明 音视频都有，且都已ready
//  (收到过音频seq header && fragment已经打开 && 音频缓存数据不为空) 说明 为什么音频缓存需不为空？
// )

//nazalog.Debugf("> OnTsPackets. frame=%s, boundary=%v, packets=%d", frame.DebugString(), boundary, len(packets))
