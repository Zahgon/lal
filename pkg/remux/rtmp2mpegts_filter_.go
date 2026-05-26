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
)

// rtmp2MpegtsFilter
//
// 缓存流起始的一些数据，判断流中是否存在音频、视频，以及编码格式，生成正确的mpegts PatPmt头信息
//
// 一旦判断结束，该队列变成直进直出，不再有实际缓存
//
// TODO(chef): [refactor] 文件名改为 rtmp2mpegts_filter__probe.go 202304
type rtmp2MpegtsFilter struct {
	maxMsgSize int
	data       []base.RtmpMsg
	observer   iRtmp2MpegtsFilterObserver

	audioCodecId int
	videoCodecId int
	done         bool
}

type iRtmp2MpegtsFilterObserver interface {
	// OnPatPmt
	//
	// 该回调一定发生在数据回调之前
	// 只会返回两种格式，h264和h265
	//
	// TODO(chef): [opt] 当没有视频时，不应该返回h264的格式
	// TODO(chef) 这里可以考虑换成只通知drain，由上层完成FragmentHeader的组装逻辑
	//
	onPatPmt(b []byte)

	onPop(msg base.RtmpMsg)
}

// NewRtmp2MpegtsFilter
//
// @param maxMsgSize: 最大缓存多少个包
func newRtmp2MpegtsFilter(maxMsgSize int, observer iRtmp2MpegtsFilterObserver) *rtmp2MpegtsFilter {
	_ = "STUB: not implemented"
	return nil
}

// Push
//
// @param msg: 函数调用结束后，内部不持有该内存块
func (q *rtmp2MpegtsFilter) Push(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

func (q *rtmp2MpegtsFilter) drain() { _ = "STUB: not implemented"; return }
