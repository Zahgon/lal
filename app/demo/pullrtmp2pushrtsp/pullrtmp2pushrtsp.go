// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"github.com/q191201771/lal/pkg/rtprtcp"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/remux"
	"github.com/q191201771/lal/pkg/rtmp"
	"github.com/q191201771/lal/pkg/rtsp"
	"github.com/q191201771/lal/pkg/sdp"
	"github.com/q191201771/naza/pkg/nazalog"
)

func main() {
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.AssertBehavior = nazalog.AssertFatal
	})
	base.LogoutStartInfo()

	inRtmpUrl, outRtspUrl, overTcp := parseFlag()

	pushSession := rtsp.NewPushSession(func(option *rtsp.PushSessionOption) {
		option.OverTcp = overTcp == 1
	})

	remuxer := remux.NewRtmp2RtspRemuxer(
		func(sdpCtx sdp.LogicContext) {
			// remuxer完成前期工作，生成sdp并开始push
			nazalog.Info("start push.")
			err := pushSession.WithSdpLogicContext(sdpCtx).Start(outRtspUrl)
			nazalog.Assert(nil, err)
			nazalog.Info("push succ.")

		},
		func(pkt rtprtcp.RtpPacket) {
			_ = pushSession.WriteRtpPacket(pkt) // remuxer的数据给push发送
		},
	)

	pullSession := rtmp.NewPullSession().WithOnReadRtmpAvMsg(remuxer.FeedRtmpMsg)

	nazalog.Info("start pull.")
	err := pullSession.Start(inRtmpUrl)
	nazalog.Assert(nil, err)
	nazalog.Info("pull succ.")

	select {
	case err := <-pullSession.WaitChan():
		nazalog.Fatalf("pull stopped. err=%+v", err)
	case err := <-pushSession.WaitChan():
		nazalog.Fatalf("push stopped. err=%+v", err)
	}
}

func parseFlag() (inRtmpUrl string, outRtspUrl string, overTcp int) {
	_ = "STUB: not implemented"
	return "", "", 0
}
