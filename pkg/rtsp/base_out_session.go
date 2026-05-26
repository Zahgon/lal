// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

import (
	"net"
	"sync"

	"github.com/q191201771/naza/pkg/nazaatomic"

	"github.com/q191201771/lal/pkg/rtprtcp"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/sdp"
	"github.com/q191201771/naza/pkg/nazanet"
)

// BaseOutSession out的含义是音视频由本端发送至对端
type BaseOutSession struct {
	cmdSession IInterleavedPacketWriter

	sdpCtx sdp.LogicContext

	audioRtpConn     *nazanet.UdpConnection
	videoRtpConn     *nazanet.UdpConnection
	audioRtcpConn    *nazanet.UdpConnection
	videoRtcpConn    *nazanet.UdpConnection
	audioRtpChannel  int
	audioRtcpChannel int
	videoRtpChannel  int
	videoRtcpChannel int

	sessionStat base.BasicSessionStat

	// only for debug log
	debugLogMaxCount         int
	loggedWriteAudioRtpCount int
	loggedWriteVideoRtpCount int
	loggedReadRtpCount       nazaatomic.Int32 // 因为音频和视频是两个连接，所以需要原子操作
	loggedReadRtcpCount      nazaatomic.Int32

	disposeOnce sync.Once
	waitChan    chan error
}

func NewBaseOutSession(sessionType base.SessionType, cmdSession IInterleavedPacketWriter) *BaseOutSession {
	_ = "STUB: not implemented"
	return nil
}

func (session *BaseOutSession) InitWithSdp(sdpCtx sdp.LogicContext) {
	_ = "STUB: not implemented"
	return
}

func (session *BaseOutSession) SetupWithConn(uri string, rtpConn, rtcpConn *nazanet.UdpConnection) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *BaseOutSession) SetupWithChannel(uri string, rtpChannel, rtcpChannel int) error {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
// IClientSessionLifecycle interface
// ---------------------------------------------------------------------------------------------------------------------

// Dispose 文档请参考： IClientSessionLifecycle interface
func (session *BaseOutSession) Dispose() error { _ = "STUB: not implemented"; return nil }

// WaitChan 文档请参考： IClientSessionLifecycle interface
//
// 注意，目前只有一种情况，即上层主动调用Dispose函数，此时error为nil
func (session *BaseOutSession) WaitChan() <-chan error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------

func (session *BaseOutSession) HandleInterleavedPacket(b []byte, channel int) {
	_ = "STUB: not implemented"
	return
}

func (session *BaseOutSession) WriteRtpPacket(packet rtprtcp.RtpPacket) error {
	_ = "STUB: not implemented"

	// 发送数据时，保证和sdp的原始类型对应
	return nil
}

// ----- ISessionStat --------------------------------------------------------------------------------------------------

func (session *BaseOutSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *new(base.StatSession)
}

func (session *BaseOutSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

func (session *BaseOutSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

// ---------------------------------------------------------------------------------------------------------------------

func (session *BaseOutSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

func (session *BaseOutSession) onReadRtpPacket(b []byte, rAddr *net.UDPAddr, err error) bool {
	_ = "STUB: not implemented"
	// TODO(chef): [fix] 在收到rtp和rtcp的地方，加入stat统计 202205
	return false
}

func (session *BaseOutSession) onReadRtcpPacket(b []byte, rAddr *net.UDPAddr, err error) bool {
	_ = "STUB: not implemented"
	// TODO chef: impl me
	return false
}

func (session *BaseOutSession) dispose(err error) error { _ = "STUB: not implemented"; return nil }
