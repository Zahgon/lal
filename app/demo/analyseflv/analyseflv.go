// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"strings"
	"time"

	"github.com/q191201771/lal/pkg/base"

	"github.com/q191201771/naza/pkg/bitrate"

	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/naza/pkg/nazalog"
)

// 分析诊断HTTP-FLV流以及FLV文件的小工具。
//
// 功能：
// - 时间戳回退检查
//     - 当音频时间戳出现回退时打error日志
//     - 当视频时间戳出现回退时打error日志
//     - 将音频和视频时间戳看成一个整体，出现回退时打error日志
// - 定时打印：
//     - 总体带宽
//     - 音频带宽
//     - 视频带宽
//     - 视频DTS和PTS不相等的计数
//     - I帧间隔时间
// - metadata
// - H264
//     - 打印每个tag的类型：key seq header...
//     - 打印每个tag中有多少个帧：SPS PPS SEI IDR SLICE...
//     - 打印每个SLICE的类型：I、P、B...
// - AAC
//     - 解析seq header
//

// TODO
// - 检查时间戳正向大的跳跃
// - 打印GOP中帧数量？
// - slice_num?

var (
	timestampCheckFlag   = true
	printStatFlag        = true
	printEveryTagFlag    = true
	printMetaData        = true
	analysisVideoTagFlag = true
)

var (
	prevAudioTs = int64(-1)
	prevVideoTs = int64(-1)
	prevTs      = int64(-1)
	prevIdrTs   = int64(-1)
	diffIdrTs   = int64(-1)
)

var brTotal = bitrate.New(func(option *bitrate.Option) {
	option.WindowMs = 5000
})

var brAudio = bitrate.New(func(option *bitrate.Option) {
	option.WindowMs = 5000
})

var brVideo = bitrate.New(func(option *bitrate.Option) {
	option.WindowMs = 5000
})

var videoCtsNotZeroCount = 0

func handleTags(tag httpflv.Tag) bool { _ = "STUB: not implemented"; return false }

//nazalog.Debugf("header=%+v, body=%s", tag.Header, hex.Dump(nazabytes.Prefix(tag.Payload(), 128)))

func main() {
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.AssertBehavior = nazalog.AssertFatal
	})
	defer nazalog.Sync()
	base.LogoutStartInfo()

	in := parseFlag()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			if printStatFlag {
				nazalog.Debugf("stat. total=%dKb/s, audio=%dKb/s, video=%dKb/s, videoCtsNotZeroCount=%d, diffIdrTs=%d",
					int(brTotal.Rate()), int(brAudio.Rate()), int(brVideo.Rate()), videoCtsNotZeroCount, diffIdrTs)
			}
		}
	}()

	if strings.HasPrefix(in, "http") || strings.HasPrefix(in, "https") {
		session := httpflv.NewPullSession().WithOnReadFlvTag(func(tag httpflv.Tag) {
			handleTags(tag)
		})

		// TODO(chef): [refactor] 统一 PullSession 和 FilePump 的回调格式 202211
		err := session.Start(in)
		nazalog.Assert(nil, err)

		// 临时测试一下主动关闭client session
		//go func() {
		//	time.Sleep(5 * time.Second)
		//	_ = session.Dispose()
		//}()

		err = <-session.WaitChan()
		nazalog.Errorf("< session.WaitChan. err=%+v", err)
	} else {
		err := httpflv.NewFlvFilePump().Pump(in, handleTags)
		nazalog.Assert(nil, err)
	}
}

const (
	typeUnknown uint8 = 1
	typeAvc     uint8 = 2
	typeHevc    uint8 = 3
)

var t uint8 = typeUnknown

func analysisVideoTag(tag httpflv.Tag) { _ = "STUB: not implemented"; return }

// SeiDelayMs 注意，SEI的内容是自定义格式，解析的代码不具有通用性
func SeiDelayMs(seiNalu []byte) int {
	_ = "STUB: not implemented"
	// nazalog.Debugf("sei: %s", hex.Dump(seiNalu))
	return 0
}

func parseFlag() string { _ = "STUB: not implemented"; return "" }
