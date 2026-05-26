// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

import (
	"sync"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/sdp"
	"github.com/q191201771/naza/pkg/nazanet"
)

type IPullSessionObserver interface {
	IBaseInSessionObserver
}

type PullSessionOption struct {
	// 从调用Pull函数，到接收音视频数据的前一步，也即收到rtsp play response的超时时间
	// 如果为0，则没有超时时间
	PullTimeoutMs int

	OverTcp bool // 是否使用interleaved模式，也即是否通过rtsp command tcp连接传输rtp/rtcp数据
}

var defaultPullSessionOption = PullSessionOption{
	PullTimeoutMs: 10000,
	OverTcp:       false,
}

type PullSession struct {
	onDescribeResponse func()

	cmdSession    *ClientCommandSession
	baseInSession *BaseInSession

	disposeOnce sync.Once
	waitChan    chan error
}

type ModPullSessionOption func(option *PullSessionOption)

func NewPullSession(observer IPullSessionObserver, modOptions ...ModPullSessionOption) *PullSession {
	_ = "STUB: not implemented"
	// TODO(chef): refactor 把observer从New中移除到With的函数中
	return nil
}

func (session *PullSession) WithOnDescribeResponse(onDescribeResponse func()) *PullSession {
	_ = "STUB: not implemented"
	return nil
}

// Start 阻塞直到和对端完成拉流前，握手部分的工作（也即收到RTSP Play response），或者发生错误
func (session *PullSession) Start(rawUrl string) error { _ = "STUB: not implemented"; return nil }

// 管理内部的多个资源，确保:
// 1. 一个资源销毁后，其他资源也被销毁
// 2. 所有资源都销毁后才通知上层

// err是nil时，表示是被PullSession::Dispose主动销毁，那么cmdSession也会被销毁，就不需要我们再调用cmdSession.Dispose了

// select loop

// 第一个错误作为返回值

// for loop

// Pull deprecated. use Start instead.
func (session *PullSession) Pull(rawUrl string) error { _ = "STUB: not implemented"; return nil }

func (session *PullSession) GetSdp() sdp.LogicContext {
	_ = "STUB: not implemented"
	return *new(sdp.LogicContext)
}

// ---------------------------------------------------------------------------------------------------------------------
// IClientSessionLifecycle interface
// ---------------------------------------------------------------------------------------------------------------------

// Dispose 文档请参考： IClientSessionLifecycle interface
func (session *PullSession) Dispose() error { _ = "STUB: not implemented"; return nil }

// WaitChan 文档请参考： IClientSessionLifecycle interface
func (session *PullSession) WaitChan() <-chan error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionUrlContext interface
// ---------------------------------------------------------------------------------------------------------------------

// Url 文档请参考： interface ISessionUrlContext
func (session *PullSession) Url() string { _ = "STUB: not implemented"; return "" }

// AppName 文档请参考： interface ISessionUrlContext
func (session *PullSession) AppName() string { _ = "STUB: not implemented"; return "" }

// StreamName 文档请参考： interface ISessionUrlContext
func (session *PullSession) StreamName() string { _ = "STUB: not implemented"; return "" }

// RawQuery 文档请参考： interface ISessionUrlContext
func (session *PullSession) RawQuery() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------
// IObject interface
// ---------------------------------------------------------------------------------------------------------------------

// UniqueKey 文档请参考： interface IObject
func (session *PullSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionStat interface
// ---------------------------------------------------------------------------------------------------------------------

// GetStat 文档请参考： interface ISessionStat
func (session *PullSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *new(base.StatSession)
}

// UpdateStat 文档请参考： interface ISessionStat
func (session *PullSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

// IsAlive 文档请参考： interface ISessionStat
func (session *PullSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

// ---------------------------------------------------------------------------------------------------------------------
// IClientCommandSessionObserver interface
// ---------------------------------------------------------------------------------------------------------------------

// OnConnectResult callback by ClientCommandSession
func (session *PullSession) OnConnectResult() {
	_ = "STUB: not implemented"
	// noop

	// OnDescribeResponse callback by ClientCommandSession
	return
}

func (session *PullSession) OnDescribeResponse(sdpCtx sdp.LogicContext) {
	_ = "STUB: not implemented"
	return
}

// OnSetupWithConn callback by ClientCommandSession
func (session *PullSession) OnSetupWithConn(uri string, rtpConn, rtcpConn *nazanet.UdpConnection) {
	_ = "STUB: not implemented"
	return
}

// OnSetupWithChannel callback by ClientCommandSession
func (session *PullSession) OnSetupWithChannel(uri string, rtpChannel, rtcpChannel int) {
	_ = "STUB: not implemented"
	return
}

// OnSetupResult callback by ClientCommandSession
func (session *PullSession) OnSetupResult() { _ = "STUB: not implemented"; return }

// OnInterleavedPacket callback by ClientCommandSession
func (session *PullSession) OnInterleavedPacket(packet []byte, channel int) {
	_ = "STUB: not implemented"
	return
}

// ---------------------------------------------------------------------------------------------------------------------
// IInterleavedPacketWriter interface
// ---------------------------------------------------------------------------------------------------------------------

// WriteInterleavedPacket callback by BaseInSession
func (session *PullSession) WriteInterleavedPacket(packet []byte, channel int) error {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

func (session *PullSession) dispose(err error) error { _ = "STUB: not implemented"; return nil }

func defaultOnDescribeResponse() { _ = "STUB: not implemented"; return }
