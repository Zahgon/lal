// Copyright 2020, Chef.  All rights reserved.
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

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/rtprtcp"
	"github.com/q191201771/lal/pkg/sdp"
	"github.com/q191201771/naza/pkg/nazanet"
)

// 聚合PubSession和PullSession，也即流数据是输入类型的session

// IBaseInSessionObserver
//
// BaseInSession会向上层回调两种格式的数据(本质上是一份数据，业务方可自由选择使用)：
// 1. 原始的rtp packet
// 2. rtp合并后的av packet
type IBaseInSessionObserver interface {
	OnSdp(sdpCtx sdp.LogicContext)

	// OnRtpPacket 回调收到的RTP包
	//
	OnRtpPacket(pkt rtprtcp.RtpPacket)

	// OnAvPacket
	//
	// @param pkt: pkt结构体中字段含义见 rtprtcp.OnAvPacket
	//
	OnAvPacket(pkt base.AvPacket)
}

type BaseInSession struct {
	cmdSession IInterleavedPacketWriter

	observer IBaseInSessionObserver

	audioRtpConn     *nazanet.UdpConnection
	videoRtpConn     *nazanet.UdpConnection
	audioRtcpConn    *nazanet.UdpConnection
	videoRtcpConn    *nazanet.UdpConnection
	audioRtpChannel  int
	audioRtcpChannel int
	videoRtpChannel  int
	videoRtcpChannel int

	sessionStat base.BasicSessionStat

	mu              sync.Mutex
	sdpCtx          sdp.LogicContext // const after set
	avPacketQueue   *AvPacketQueue
	audioRrProducer *rtprtcp.RrProducer
	videoRrProducer *rtprtcp.RrProducer

	audioUnpacker rtprtcp.IRtpUnpacker
	videoUnpacker rtprtcp.IRtpUnpacker

	audioSsrc nazaatomic.Uint32
	videoSsrc nazaatomic.Uint32

	disposeOnce sync.Once
	waitChan    chan error

	dumpReadAudioRtp base.LogDump
	dumpReadVideoRtp base.LogDump
	dumpReadRtcp     base.LogDump
	dumpReadSr       base.LogDump
}

func NewBaseInSession(sessionType base.SessionType, cmdSession IInterleavedPacketWriter) *BaseInSession {
	_ = "STUB: not implemented"
	return nil
}

func NewBaseInSessionWithObserver(sessionType base.SessionType, cmdSession IInterleavedPacketWriter, observer IBaseInSessionObserver) *BaseInSession {
	_ = "STUB: not implemented"
	return nil
}

func (session *BaseInSession) InitWithSdp(sdpCtx sdp.LogicContext) {
	_ = "STUB: not implemented"
	return
}

// SetObserver 如果没有设置回调监听对象，可以通过该函数设置，调用方保证调用该函数发生在调用InitWithSdp之后
func (session *BaseInSession) SetObserver(observer IBaseInSessionObserver) {
	_ = "STUB: not implemented"
	return
}

// 避免在当前协程回调，降低业务方使用负担，不必担心设置监听对象和回调函数中锁重入 TODO(chef): 更好的方式

func (session *BaseInSession) SetupWithConn(uri string, rtpConn, rtcpConn *nazanet.UdpConnection) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *BaseInSession) SetupWithChannel(uri string, rtpChannel, rtcpChannel int) error {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
// IClientSessionLifecycle interface
// ---------------------------------------------------------------------------------------------------------------------

// Dispose 文档请参考： IClientSessionLifecycle interface
func (session *BaseInSession) Dispose() error { _ = "STUB: not implemented"; return nil }

// WaitChan 文档请参考： IClientSessionLifecycle interface
//
// 注意，目前只有一种情况，即上层主动调用Dispose函数，此时error为nil
func (session *BaseInSession) WaitChan() <-chan error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------

func (session *BaseInSession) GetSdp() sdp.LogicContext {
	_ = "STUB: not implemented"
	return *new(sdp.LogicContext)
}

func (session *BaseInSession) HandleInterleavedPacket(b []byte, channel int) {
	_ = "STUB: not implemented"
	return
}

// WriteRtpRtcpDummy 发现pull时，需要先给对端发送数据，才能收到数据
func (session *BaseInSession) WriteRtpRtcpDummy() { _ = "STUB: not implemented"; return }

// ----- ISessionStat --------------------------------------------------------------------------------------------------

func (session *BaseInSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *new(base.StatSession)
}

func (session *BaseInSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

func (session *BaseInSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

// ---------------------------------------------------------------------------------------------------------------------

func (session *BaseInSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// callback by RTPUnpacker
func (session *BaseInSession) onAvPacketUnpacked(pkt base.AvPacket) {
	_ = "STUB: not implemented"
	return
}

// callback by avpacket queue
func (session *BaseInSession) onAvPacket(pkt base.AvPacket) { _ = "STUB: not implemented"; return }

// callback by UDPConnection
func (session *BaseInSession) onReadRtpPacket(b []byte, rAddr *net.UDPAddr, err error) bool {
	_ = "STUB: not implemented"

	// TODO(chef):
	// read udp [::]:30008: use of closed network connection
	// 可以退出loop，看是在上层退还是下层退，但是要注意每次read都判断的开销
	return false
}

// callback by UDPConnection
func (session *BaseInSession) onReadRtcpPacket(b []byte, rAddr *net.UDPAddr, err error) bool {
	_ = "STUB: not implemented"
	return false
}

// @param rAddr 对端地址，往对端发送数据时使用，注意，如果nil，则表示是interleaved模式，我们直接往TCP连接发数据
func (session *BaseInSession) handleRtcpPacket(b []byte, rAddr *net.UDPAddr) error {
	_ = "STUB: not implemented"
	return nil
}

// noop
//
// ffmpeg推流时，会在发送第一个RTP包之前就发送一个SR，所以关闭这个警告日志
//Log.Warnf("[%s] read rtcp sr but senderSsrc invalid. senderSsrc=%d, audio=%d, video=%d",
//	p.uniqueKey, sr.SenderSsrc, p.audioSsrc, p.videoSsrc)

func (session *BaseInSession) handleRtpPacket(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

//Log.Errorf("[%s] handleRtpPacket but type invalid. type=%d", session.UniqueKey(), packetType)

// 接收数据时，保证了sdp的原始类型对应

// noop 因为前面已经判断过type了，所以永远不会走到这

func (session *BaseInSession) dispose(err error) error { _ = "STUB: not implemented"; return nil }
