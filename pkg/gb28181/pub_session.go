// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package gb28181

import (
	"net"
	"sync"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/naza/pkg/nazanet"
)

type OnReadPacket func(b []byte)

type PubSession struct {
	unpacker *PsUnpacker

	streamName string

	hookOnReadPacket OnReadPacket

	isTcpFlag bool

	disposeOnce sync.Once
	udpConn     *nazanet.UdpConnection
	listener    net.Listener
	tcpConn     net.Conn
	sessionStat base.BasicSessionStat
}

func NewPubSession() *PubSession { _ = "STUB: not implemented"; return nil }

// WithOnAvPacket 设置音视频的回调。
//
//	@param onAvPacket: 见 PsUnpacker.WithOnAvPacket 的注释
func (session *PubSession) WithOnAvPacket(onAvPacket base.OnAvPacketFunc) *PubSession {
	_ = "STUB: not implemented"
	return nil
}

func (session *PubSession) WithStreamName(streamName string) *PubSession {
	_ = "STUB: not implemented"
	return nil
}

// WithHookReadPacket
//
// 将接收的数据返回给上层。
// 注意，底层的解析逻辑依然走。
// 可以用这个方式来截取数据进行调试。
func (session *PubSession) WithHookReadPacket(fn OnReadPacket) *PubSession {
	_ = "STUB: not implemented"
	return nil
}

// Listen 非阻塞函数
//
// 注意，当`port`参数为0时，内部会自动选择一个可用端口监听，并通过返回值返回该端口
func (session *PubSession) Listen(port int, isTcpFlag bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// RunLoop 阻塞函数
func (session *PubSession) RunLoop() error { _ = "STUB: not implemented"; return nil }

// ----- IServerSessionLifecycle ---------------------------------------------------------------------------------------

func (session *PubSession) Dispose() error { _ = "STUB: not implemented"; return nil }

// ----- ISessionUrlContext --------------------------------------------------------------------------------------------

func (session *PubSession) Url() string { _ = "STUB: not implemented"; return "" }

func (session *PubSession) AppName() string { _ = "STUB: not implemented"; return "" }

func (session *PubSession) StreamName() string {
	_ = "STUB: not implemented"
	// 如果stream name没有设置，则使用session的unique key作为stream name
	return ""
}

func (session *PubSession) RawQuery() string { _ = "STUB: not implemented"; return "" }

// ----- IObject -------------------------------------------------------------------------------------------------------

func (session *PubSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// ----- ISessionStat --------------------------------------------------------------------------------------------------

func (session *PubSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

func (session *PubSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *new(base.StatSession)
}

func (session *PubSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

// ---------------------------------------------------------------------------------------------------------------------

func (session *PubSession) listenUdp(port int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (session *PubSession) listenTcp(port int) (int, error) {
	_ = "STUB: not implemented"

	// TODO(chef): [refactor] 考虑挪到naza中，udp在naza中有类似的实现 202209
	return 0, nil
}

func (session *PubSession) runLoopUdp() error { _ = "STUB: not implemented"; return nil }

func (session *PubSession) runLoopTcp() error { _ = "STUB: not implemented"; return nil }

// TODO(chef): [fix] reset unpack 202209

// 初始1500，如果不够会扩容

func (session *PubSession) feedPacket(b []byte) { _ = "STUB: not implemented"; return }

func (session *PubSession) dispose(err error) error { _ = "STUB: not implemented"; return nil }
