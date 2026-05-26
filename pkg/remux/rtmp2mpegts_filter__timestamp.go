// Copyright 2023, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package remux

import (
	"github.com/q191201771/lal/pkg/mpegts"
)

// Rtmp2MpegtsTimestampFilter
//
// (1). 通过dts和cts计算pts。
// (2). 将数据包中的时间戳转换为从零开始的时间戳。
//
// TODO(chef): [opt] 音频和视频公用一个base，避免以下情况:
// A:           6 7 ...
// V: 1 2 3 4 5 6 7 ...
// 另外，需要考虑音频比base还小的情况
type Rtmp2MpegtsTimestampFilter struct {
	uk string

	basicAudioDts uint64
	basicVideoDts uint64
}

func (f *Rtmp2MpegtsTimestampFilter) Init(uk string) { _ = "STUB: not implemented"; return }

// Do
//
// @param frame: 直接修改frame中的dts和pts
func (f *Rtmp2MpegtsTimestampFilter) Do(frame *mpegts.Frame) { _ = "STUB: not implemented"; return }
