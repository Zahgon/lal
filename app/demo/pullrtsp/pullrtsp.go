// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"time"

	"github.com/q191201771/lal/pkg/rtprtcp"
	"github.com/q191201771/lal/pkg/sdp"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/lal/pkg/remux"
	"github.com/q191201771/lal/pkg/rtsp"
	"github.com/q191201771/naza/pkg/nazalog"
)

// pullrtsp 拉取rtsp流，然后存储为flv文件或者dump文件进行分析

var remuxer *remux.AvPacket2RtmpRemuxer
var dump *base.DumpFile

type Observer struct{}

func (o *Observer) OnSdp(sdpCtx sdp.LogicContext) { _ = "STUB: not implemented"; return }

func (o *Observer) OnRtpPacket(pkt rtprtcp.RtpPacket) { _ = "STUB: not implemented"; return }

func (o *Observer) OnAvPacket(pkt base.AvPacket) {
	_ = "STUB: not implemented"
	// nazalog.Debugf("OnAvPacket %+v", pkt.DebugString())
	return
}

func main() {
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.AssertBehavior = nazalog.AssertFatal
		option.IsToStdout = true
		option.Filename = "pullrtsp.log"
	})
	defer nazalog.Sync()
	base.LogoutStartInfo()

	inUrl, outFilename, overTcp, debugDumpPacket := parseFlag()

	if debugDumpPacket != "" {
		dump = base.NewDumpFile()
		err := dump.OpenToWrite(debugDumpPacket)
		nazalog.Assert(nil, err)
	}

	var fileWriter httpflv.FlvFileWriter
	err := fileWriter.Open(outFilename)
	nazalog.Assert(nil, err)
	defer fileWriter.Dispose()
	err = fileWriter.WriteRaw(httpflv.FlvHeader)
	nazalog.Assert(nil, err)

	remuxer = remux.NewAvPacket2RtmpRemuxer().WithOnRtmpMsg(func(msg base.RtmpMsg) {
		err = fileWriter.WriteTag(*remux.RtmpMsg2FlvTag(msg))
		nazalog.Assert(nil, err)
	})

	var observer Observer

	pullSession := rtsp.NewPullSession(&observer, func(option *rtsp.PullSessionOption) {
		option.PullTimeoutMs = 10000
		option.OverTcp = overTcp != 0
	})

	err = pullSession.Start(inUrl)
	nazalog.Assert(nil, err)

	go func() {
		for {
			pullSession.UpdateStat(1)
			nazalog.Debugf("stat. pull=%+v", pullSession.GetStat())
			time.Sleep(1 * time.Second)
		}
	}()

	// 临时测试一下主动关闭client session
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	err := pullSession.Dispose()
	//	nazalog.Debugf("< session Dispose. err=%+v", err)
	//}()

	err = <-pullSession.WaitChan()
	nazalog.Infof("< pullSession.Wait(). err=%+v", err)
}

func parseFlag() (inUrl string, outFilename string, overTcp int, debugDumpPacket string) {
	_ = "STUB: not implemented"
	return "", "", 0, ""
}
