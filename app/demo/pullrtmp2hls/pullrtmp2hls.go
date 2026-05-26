// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"github.com/q191201771/lal/pkg/remux"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/hls"
	"github.com/q191201771/lal/pkg/rtmp"
	"github.com/q191201771/naza/pkg/nazalog"
)

func main() {
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.AssertBehavior = nazalog.AssertFatal
	})
	defer nazalog.Sync()
	base.LogoutStartInfo()

	url, hlsOutPath, fragmentDurationMs, fragmentNum := parseFlag()
	nazalog.Infof("parse flag succ. url=%s, hlsOutPath=%s, fragmentDurationMs=%d, fragmentNum=%d",
		url, hlsOutPath, fragmentDurationMs, fragmentNum)

	hlsMuxerConfig := hls.MuxerConfig{
		OutPath:            hlsOutPath,
		FragmentDurationMs: fragmentDurationMs,
		FragmentNum:        fragmentNum,
	}

	ctx, err := base.ParseRtmpUrl(url)
	if err != nil {
		nazalog.Fatalf("parse rtmp url failed. url=%s, err=%+v", url, err)
	}
	streamName := ctx.LastItemOfPath

	hlsMuexer := hls.NewMuxer(streamName, &hlsMuxerConfig, nil)
	hlsMuexer.Start()

	rtmp2Mpegts := remux.NewRtmp2MpegtsRemuxer(hlsMuexer)

	pullSession := rtmp.NewPullSession(func(option *rtmp.PullSessionOption) {
		option.PullTimeoutMs = 10000
		option.ReadAvTimeoutMs = 10000
	}).WithOnReadRtmpAvMsg(rtmp2Mpegts.FeedRtmpMessage)
	err = pullSession.Start(url)

	if err != nil {
		nazalog.Fatalf("pull rtmp failed. err=%+v", err)
	}
	err = <-pullSession.WaitChan()
	nazalog.Errorf("< session.Wait [%s] err=%+v", pullSession.UniqueKey(), err)
}

func parseFlag() (url string, hlsOutPath string, fragmentDurationMs int, fragmentNum int) {
	_ = "STUB: not implemented"
	return "", "", 0, 0
}
