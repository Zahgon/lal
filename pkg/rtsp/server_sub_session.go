// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

import (
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/rtprtcp"
	"github.com/q191201771/lal/pkg/sdp"
	"github.com/q191201771/naza/pkg/nazaatomic"
	"github.com/q191201771/naza/pkg/nazanet"
)

type SubSessionStage int

const (
	SubSessionStageReadDescribe int32 = 0 // 初时阶段，已收到 describe
	SubSessionStageWriteSdp           = 1 // 已发送 sdp
	SubSessionStageReadPlay           = 2 // 已收到 play
)

type SubSession struct {
	urlCtx         base.UrlContext
	cmdSession     *ServerCommandSession
	baseOutSession *BaseOutSession

	ShouldWaitVideoKeyFrame bool

	Stage nazaatomic.Int32 // 见 SubSessionStageReadDescribe 等常量定义
}

func NewSubSession(urlCtx base.UrlContext, cmdSession *ServerCommandSession) *SubSession {
	_ = "STUB: not implemented"
	return nil
}

// FeedSdp 供上层调用
func (session *SubSession) FeedSdp(sdpCtx sdp.LogicContext) { _ = "STUB: not implemented"; return }

// InitWithSdp 供 ServerCommandSession 调用
func (session *SubSession) InitWithSdp(sdpCtx sdp.LogicContext) { _ = "STUB: not implemented"; return }

func (session *SubSession) SetupWithConn(uri string, rtpConn, rtcpConn *nazanet.UdpConnection) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *SubSession) SetupWithChannel(uri string, rtpChannel, rtcpChannel int) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *SubSession) WriteRtpPacket(packet rtprtcp.RtpPacket) {
	_ = "STUB: not implemented"
	return
}

//Log.Warnf("[%s] write rtp packet is not as expected, stage is not ready yet.. stage=%d", session.UniqueKey(), stage)

func (session *SubSession) Dispose() error { _ = "STUB: not implemented"; return nil }

func (session *SubSession) HandleInterleavedPacket(b []byte, channel int) {
	_ = "STUB: not implemented"
	return
}

func (session *SubSession) Url() string { _ = "STUB: not implemented"; return "" }

func (session *SubSession) AppName() string { _ = "STUB: not implemented"; return "" }

func (session *SubSession) StreamName() string { _ = "STUB: not implemented"; return "" }

func (session *SubSession) RawQuery() string { _ = "STUB: not implemented"; return "" }

func (session *SubSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

func (session *SubSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *new(base.StatSession)
}

func (session *SubSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

func (session *SubSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

// WriteInterleavedPacket IInterleavedPacketWriter, callback by BaseOutSession
func (session *SubSession) WriteInterleavedPacket(packet []byte, channel int) error {
	_ = "STUB: not implemented"
	return nil
}
