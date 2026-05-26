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

	"github.com/q191201771/naza/pkg/connection"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/naza/pkg/nazahttp"
)

type IServerCommandSessionObserver interface {
	// OnNewRtspPubSession
	//
	// @brief  Announce阶段回调
	// @return 如果返回非nil，则表示上层要强制关闭这个推流请求
	//
	OnNewRtspPubSession(session *PubSession) error

	// OnNewRtspSubSessionDescribe
	//
	// Describe阶段回调
	//
	// 上层的几种逻辑对应的返回值的组合情况：
	//
	// 1. 强制关闭这个session：`ok`设置为false
	// 2. （当前有sdp）在回调中设置sdp，让session按正常逻辑往下走：`ok`设置为true，`sdp`设置为对应的值
	// 3. （当前没有sdp）后续在回调外通过 ServerCommandSession.FeedSdp 可以让session按正常逻辑往下走：`ok`设置为true，`sdp`设置为nil
	//
	// TODO(chef): bool参数类型统一为error类型 202206
	//
	OnNewRtspSubSessionDescribe(session *SubSession) (ok bool, sdp []byte)

	// OnNewRtspSubSessionPlay
	//
	// @brief Play阶段回调
	// @return ok  如果返回非nil，则表示上层要强制关闭这个拉流请求
	//
	OnNewRtspSubSessionPlay(session *SubSession) error
}

type ServerCommandSession struct {
	uniqueKey    string                        // const after ctor
	observer     IServerCommandSessionObserver // const after ctor
	conn         connection.Connection
	prevConnStat connection.Stat
	staleStat    *connection.Stat
	stat         base.StatSession
	authConf     ServerAuthConfig
	auth         Auth

	pubSession *PubSession
	subSession *SubSession

	describeSeq  string // only for sub session
	isWebSocket  bool
	websocketKey string
}

func NewServerCommandSession(observer IServerCommandSessionObserver, conn net.Conn, authConf ServerAuthConfig, iswebsocket bool, websocketKey string) *ServerCommandSession {
	_ = "STUB: not implemented"
	return nil
}

func (session *ServerCommandSession) RunLoop() error { _ = "STUB: not implemented"; return nil }

func (session *ServerCommandSession) Dispose() error { _ = "STUB: not implemented"; return nil }

func (session *ServerCommandSession) FeedSdp(b []byte) {
	_ = "STUB: not implemented"

	// WriteInterleavedPacket
	//
	// 使用RTSP TCP命令连接，向对端发送RTP数据
	return
}

func (session *ServerCommandSession) WriteInterleavedPacket(packet []byte, channel int) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *ServerCommandSession) RemoteAddr() string { _ = "STUB: not implemented"; return "" }

// ----- ISessionStat --------------------------------------------------------------------------------------------------

func (session *ServerCommandSession) UpdateStat(intervalSec uint32) {
	_ = "STUB: not implemented"
	// TODO(chef): 梳理interleaved模式下，command session的ISessionStat 202205
	return
}

func (session *ServerCommandSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *new(base.StatSession)
}

func (session *ServerCommandSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

// ----- IObject -------------------------------------------------------------------------------------------------------

func (session *ServerCommandSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------

func (session *ServerCommandSession) runCmdLoop() error { _ = "STUB: not implemented"; return nil }

// 解析出websocket的body信息

// 读取一个message

// 读取一个message

// pub, sub

// pub

// sub

// pub, sub

// pub

// sub

// pub

func (session *ServerCommandSession) handleOptions(requestCtx nazahttp.HttpReqMsgCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *ServerCommandSession) handleAnnounce(requestCtx nazahttp.HttpReqMsgCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *ServerCommandSession) handleDescribe(requestCtx nazahttp.HttpReqMsgCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// 鉴权处理

func (session *ServerCommandSession) feedSdp(rawSdp []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *ServerCommandSession) handleAuthorized(requestCtx nazahttp.HttpReqMsgCtx) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// 解析出的鉴权方式需要与配置的鉴权方式一致,防止鉴权降级

// TODO(chef): [refactor] 错误放入base/error.go中 202205

// Basic鉴权

// Digest鉴权

// TODO(chef): [refactor] 错误放入base/error.go中 202205

// 一次SETUP对应一路流（音频或视频）
func (session *ServerCommandSession) handleSetup(requestCtx nazahttp.HttpReqMsgCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// 是否为interleaved模式

func (session *ServerCommandSession) handleRecord(requestCtx nazahttp.HttpReqMsgCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *ServerCommandSession) handlePlay(requestCtx nazahttp.HttpReqMsgCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// 没有收到前面的信令，直接收到Play信令

// TODO(chef): [opt] 上层关闭，可以考虑回复非200状态码再关闭

func (session *ServerCommandSession) handleTeardown(requestCtx nazahttp.HttpReqMsgCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *ServerCommandSession) writeWsFrameHeader(respLen int) {
	_ = "STUB: not implemented"
	return
}
