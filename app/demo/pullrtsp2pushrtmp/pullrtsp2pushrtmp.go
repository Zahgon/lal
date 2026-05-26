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

	"github.com/q191201771/lal/pkg/rtmp"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/remux"
	"github.com/q191201771/lal/pkg/rtsp"
	"github.com/q191201771/naza/pkg/nazalog"
)

func main() {
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.AssertBehavior = nazalog.AssertFatal
	})
	defer nazalog.Sync()
	base.LogoutStartInfo()

	inUrl, outUrl, overTcp := parseFlag()

	pushSession := rtmp.NewPushSession(func(option *rtmp.PushSessionOption) {
		option.PushTimeoutMs = 10000
		option.WriteAvTimeoutMs = 10000
	})

	err := pushSession.Start(outUrl)
	nazalog.Assert(nil, err)
	defer pushSession.Dispose()

	remuxer := remux.NewAvPacket2RtmpRemuxer().WithOnRtmpMsg(func(msg base.RtmpMsg) {
		err = pushSession.Write(rtmp.Message2Chunks(msg.Payload, &msg.Header))
		nazalog.Assert(nil, err)
	})
	pullSession := rtsp.NewPullSession(remuxer, func(option *rtsp.PullSessionOption) {
		option.PullTimeoutMs = 10000
		option.OverTcp = overTcp != 0
	})

	err = pullSession.Start(inUrl)
	nazalog.Assert(nil, err)
	defer pullSession.Dispose()

	go func() {
		for {
			pullSession.UpdateStat(1)
			pullStat := pullSession.GetStat()
			pushSession.UpdateStat(1)
			pushStat := pushSession.GetStat()
			nazalog.Debugf("stat. pull=%+v, push=%+v", pullStat, pushStat)
			time.Sleep(1 * time.Second)
		}
	}()

	select {
	case err = <-pullSession.WaitChan():
		nazalog.Infof("< pullSession.Wait(). err=%+v", err)
		time.Sleep(1 * time.Second)
		return
	case err = <-pushSession.WaitChan():
		nazalog.Infof("< pushSession.Wait(). err=%+v", err)
		time.Sleep(1 * time.Second)
		return
	}
}

func parseFlag() (inUrl string, outUrl string, overTcp int) {
	_ = "STUB: not implemented"
	return "", "", 0
}
