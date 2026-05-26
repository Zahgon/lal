// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"sync"

	"github.com/q191201771/lal/pkg/sdp"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/rtprtcp"
	"github.com/q191201771/lal/pkg/rtsp"
	"github.com/q191201771/naza/pkg/nazalog"
)

type RtspTunnel struct {
	pullUrl     string
	pushUrl     string
	pullOverTcp bool
	pushOverTcp bool

	uniqueKey string

	rtpPacketChan chan rtprtcp.RtpPacket

	disposeOnce sync.Once
	waitChan    chan error

	pullSession *rtsp.PullSession
	pushSession *rtsp.PushSession
}

func NewRtspTunnel(pullUrl string, pushUrl string, pullOverTcp bool, pushOverTcp bool) *RtspTunnel {
	_ = "STUB: not implemented"
	return nil
}

// Start 开启任务，阻塞直到任务开启成功或失败。
//
// @return: 如果为nil，表示任务启动成功，此时数据已经在后台转发
func (r *RtspTunnel) Start() error { _ = "STUB: not implemented"; return nil }

// Dispose 主动关闭tunnel时调用
//
// 注意，只有 Start 成功后的tunnel才能调用，否则行为未定义
//
// 更详细的说明参考 IClientSessionLifecycle interface
func (r *RtspTunnel) Dispose() error { _ = "STUB: not implemented"; return nil }

// WaitChan Start 成功后，可使用这个channel来接收tunnel结束的消息
//
// 更详细的说明参考 IClientSessionLifecycle interface
func (r *RtspTunnel) WaitChan() chan error {
	_ = "STUB: not implemented"

	// ---------------------------------------------------------------------------------------------------------------------
	return nil
}

func (r *RtspTunnel) OnRtpPacket(pkt rtprtcp.RtpPacket) { _ = "STUB: not implemented"; return }

func (r *RtspTunnel) OnSdp(sdpCtx sdp.LogicContext) {
	_ = "STUB: not implemented"
	// noop
	return
}

func (r *RtspTunnel) OnAvPacket(pkt base.AvPacket) {
	_ = "STUB: not implemented"
	// noop
	return
}

func (r *RtspTunnel) dispose(err error) error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------

func main() {
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.AssertBehavior = nazalog.AssertFatal
	})
	defer nazalog.Sync()
	base.LogoutStartInfo()

	inUrl, outUrl, pullOverTcp, pushOverTcp := parseFlag()

	rtspTunnel := NewRtspTunnel(inUrl, outUrl, pullOverTcp == 1, pushOverTcp == 1)
	err := rtspTunnel.Start()
	if err != nil {
		nazalog.Errorf("start tunnel failed. err=%+v", err)
		return
	}

	//go func() {
	//	time.Sleep(5 * time.Second)
	//	_ = rtspTunnel.Dispose()
	//}()

	err = <-rtspTunnel.WaitChan()
	nazalog.Errorf("tunnel stopped. err=%+v", err)
}

func parseFlag() (inUrl string, outUrl string, pullOverTcp int, pushOverTcp int) {
	_ = "STUB: not implemented"
	return "", "", 0, 0
}
