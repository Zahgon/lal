// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

import (
	"github.com/q191201771/naza/pkg/nazanet"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/sdp"
)

type IPubSessionObserver interface {
	IBaseInSessionObserver
}

type PubSession struct {
	urlCtx        base.UrlContext
	cmdSession    *ServerCommandSession
	baseInSession *BaseInSession

	observer IPubSessionObserver
}

func NewPubSession(urlCtx base.UrlContext, cmdSession *ServerCommandSession) *PubSession {
	_ = "STUB: not implemented"
	return nil
}

func (session *PubSession) InitWithSdp(sdpCtx sdp.LogicContext) { _ = "STUB: not implemented"; return }

func (session *PubSession) SetObserver(observer IPubSessionObserver) {
	_ = "STUB: not implemented"
	return
}

func (session *PubSession) SetupWithConn(uri string, rtpConn, rtcpConn *nazanet.UdpConnection) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *PubSession) SetupWithChannel(uri string, rtpChannel, rtcpChannel int) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *PubSession) Dispose() error { _ = "STUB: not implemented"; return nil }

func (session *PubSession) GetSdp() sdp.LogicContext {
	_ = "STUB: not implemented"
	return *new(sdp.LogicContext)
}

func (session *PubSession) HandleInterleavedPacket(b []byte, channel int) {
	_ = "STUB: not implemented"
	return
}

func (session *PubSession) Url() string { _ = "STUB: not implemented"; return "" }

func (session *PubSession) AppName() string { _ = "STUB: not implemented"; return "" }

func (session *PubSession) StreamName() string { _ = "STUB: not implemented"; return "" }

func (session *PubSession) RawQuery() string { _ = "STUB: not implemented"; return "" }

func (session *PubSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// ----- ISessionStat --------------------------------------------------------------------------------------------------

func (session *PubSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *new(base.StatSession)
}

func (session *PubSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

func (session *PubSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

// ---------------------------------------------------------------------------------------------------------------------

// WriteInterleavedPacket IInterleavedPacketWriter, callback by BaseInSession
func (session *PubSession) WriteInterleavedPacket(packet []byte, channel int) error {
	_ = "STUB: not implemented"
	return nil
}
