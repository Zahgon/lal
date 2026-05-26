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
	"github.com/q191201771/lal/pkg/rtprtcp"
	"github.com/q191201771/lal/pkg/sdp"
	"github.com/q191201771/naza/pkg/nazanet"
)

type PushSessionOption struct {
	PushTimeoutMs int
	OverTcp       bool
}

var defaultPushSessionOption = PushSessionOption{
	PushTimeoutMs: 10000,
	OverTcp:       false,
}

type PushSession struct {
	cmdSession     *ClientCommandSession
	baseOutSession *BaseOutSession

	sdpCtx *sdp.LogicContext

	disposeOnce sync.Once
	waitChan    chan error
}

type ModPushSessionOption func(option *PushSessionOption)

func NewPushSession(modOptions ...ModPushSessionOption) *PushSession {
	_ = "STUB: not implemented"
	return nil
}

func (session *PushSession) WithSdpLogicContext(sdpCtx sdp.LogicContext) *PushSession {
	_ = "STUB: not implemented"
	return nil
}

// Start 阻塞直到和对端完成推流前，握手部分的工作（也即收到RTSP Record response），或者发生错误
func (session *PushSession) Start(rawUrl string) error { _ = "STUB: not implemented"; return nil }

// Push deprecated. use WithSdpLogicContext and Start instead.
func (session *PushSession) Push(rawUrl string, sdpCtx sdp.LogicContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *PushSession) WriteRtpPacket(packet rtprtcp.RtpPacket) error {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
// IClientSessionLifecycle interface
// ---------------------------------------------------------------------------------------------------------------------

// Dispose 文档请参考： IClientSessionLifecycle interface
func (session *PushSession) Dispose() error { _ = "STUB: not implemented"; return nil }

// WaitChan 文档请参考： IClientSessionLifecycle interface
func (session *PushSession) WaitChan() <-chan error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionUrlContext interface
// ---------------------------------------------------------------------------------------------------------------------

// Url 文档请参考： interface ISessionUrlContext
func (session *PushSession) Url() string { _ = "STUB: not implemented"; return "" }

// AppName 文档请参考： interface ISessionUrlContext
func (session *PushSession) AppName() string { _ = "STUB: not implemented"; return "" }

// StreamName 文档请参考： interface ISessionUrlContext
func (session *PushSession) StreamName() string { _ = "STUB: not implemented"; return "" }

// RawQuery 文档请参考： interface ISessionUrlContext
func (session *PushSession) RawQuery() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionUrlContext IObject
// ---------------------------------------------------------------------------------------------------------------------

// UniqueKey 文档请参考： interface IObject
func (session *PushSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionStat IObject
// ---------------------------------------------------------------------------------------------------------------------

// GetStat 文档请参考： interface ISessionStat
func (session *PushSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *new(base.StatSession)
}

// UpdateStat 文档请参考： interface ISessionStat
func (session *PushSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

// IsAlive 文档请参考： interface ISessionStat
func (session *PushSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

// ---------------------------------------------------------------------------------------------------------------------
// ISessionStat IClientCommandSessionObserver
// ---------------------------------------------------------------------------------------------------------------------

// OnConnectResult callback by ClientCommandSession
func (session *PushSession) OnConnectResult() {
	_ = "STUB: not implemented"
	// noop

	// OnDescribeResponse callback by ClientCommandSession
	return
}

func (session *PushSession) OnDescribeResponse(sdpCtx sdp.LogicContext) {
	_ = "STUB: not implemented"
	// noop

	// OnSetupWithConn callback by ClientCommandSession
	return
}

func (session *PushSession) OnSetupWithConn(uri string, rtpConn, rtcpConn *nazanet.UdpConnection) {
	_ = "STUB: not implemented"
	return
}

// OnSetupWithChannel callback by ClientCommandSession
func (session *PushSession) OnSetupWithChannel(uri string, rtpChannel, rtcpChannel int) {
	_ = "STUB: not implemented"
	return
}

// OnSetupResult callback by ClientCommandSession
func (session *PushSession) OnSetupResult() {
	_ = "STUB: not implemented"
	// noop

	// OnInterleavedPacket callback by ClientCommandSession
	return
}

func (session *PushSession) OnInterleavedPacket(packet []byte, channel int) {
	_ = "STUB: not implemented"
	return
}

// ---------------------------------------------------------------------------------------------------------------------
// ISessionStat IInterleavedPacketWriter
// ---------------------------------------------------------------------------------------------------------------------

// WriteInterleavedPacket callback by BaseOutSession
func (session *PushSession) WriteInterleavedPacket(packet []byte, channel int) error {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

func (session *PushSession) push(rawUrl string) error { _ = "STUB: not implemented"; return nil }

// err是nil时，表示是被PullSession::Dispose主动销毁，那么cmdSession也会被销毁，就不需要我们再调用cmdSession.Dispose了

// select loop

// 第一个错误作为返回值

// for loop

func (session *PushSession) dispose(err error) error { _ = "STUB: not implemented"; return nil }
