// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

import (
	"context"
	"sync"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/sdp"
	"github.com/q191201771/naza/pkg/connection"
	"github.com/q191201771/naza/pkg/nazahttp"
	"github.com/q191201771/naza/pkg/nazanet"
)

type ClientCommandSessionType int

const (
	readBufSize                 = 256
	writeGetParameterIntervalMs = 10000
)

const (
	CcstPullSession ClientCommandSessionType = iota
	CcstPushSession
)

type ClientCommandSessionOption struct {
	DoTimeoutMs int
	OverTcp     bool
}

var defaultClientCommandSessionOption = ClientCommandSessionOption{
	DoTimeoutMs: 10000,
	OverTcp:     false,
}

type IClientCommandSessionObserver interface {
	OnConnectResult()

	// OnDescribeResponse only for PullSession
	OnDescribeResponse(sdpCtx sdp.LogicContext)

	OnSetupWithConn(uri string, rtpConn, rtcpConn *nazanet.UdpConnection)
	OnSetupWithChannel(uri string, rtpChannel, rtcpChannel int)
	OnSetupResult()

	OnInterleavedPacket(packet []byte, channel int)
}

// ClientCommandSession Push和Pull共用，封装了客户端底层信令信令部分。
// 业务方应该使用PushSession和PullSession，而不是直接使用ClientCommandSession，除非你确定要这么做。
type ClientCommandSession struct {
	uniqueKey string
	t         ClientCommandSessionType
	observer  IClientCommandSessionObserver
	option    ClientCommandSessionOption

	rawUrl string
	urlCtx base.UrlContext
	conn   connection.Connection

	cseq                        int
	methodGetParameterSupported bool
	auth                        Auth

	sdpCtx sdp.LogicContext

	sessionId string
	channel   int

	disposeOnce sync.Once
}

type ModClientCommandSessionOption func(option *ClientCommandSessionOption)

func NewClientCommandSession(t ClientCommandSessionType, uniqueKey string, observer IClientCommandSessionObserver, modOptions ...ModClientCommandSessionOption) *ClientCommandSession {
	_ = "STUB: not implemented"
	return nil
}

// InitWithSdp only for PushSession
func (session *ClientCommandSession) InitWithSdp(sdpCtx sdp.LogicContext) {
	_ = "STUB: not implemented"
	return
}

func (session *ClientCommandSession) Start(rawUrl string) error {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
// IClientSessionLifecycle interface
// ---------------------------------------------------------------------------------------------------------------------

// Dispose 文档请参考： IClientSessionLifecycle interface
func (session *ClientCommandSession) Dispose() error { _ = "STUB: not implemented"; return nil }

// WaitChan 文档请参考： IClientSessionLifecycle interface
func (session *ClientCommandSession) WaitChan() <-chan error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------

func (session *ClientCommandSession) WriteInterleavedPacket(packet []byte, channel int) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *ClientCommandSession) RemoteAddr() string { _ = "STUB: not implemented"; return "" }

func (session *ClientCommandSession) Url() string { _ = "STUB: not implemented"; return "" }

func (session *ClientCommandSession) AppName() string { _ = "STUB: not implemented"; return "" }

func (session *ClientCommandSession) StreamName() string { _ = "STUB: not implemented"; return "" }

func (session *ClientCommandSession) RawQuery() string { _ = "STUB: not implemented"; return "" }

func (session *ClientCommandSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

func (session *ClientCommandSession) doContext(ctx context.Context, rawUrl string) error {
	_ = "STUB: not implemented"
	return nil
}

func (session *ClientCommandSession) runReadLoop() { _ = "STUB: not implemented"; return }

// TCP模式，需要收取数据进行处理

// not over tcp
// 接收TCP对端关闭FIN信号

// 对端支持get_parameter，需要定时向对端发送get_parameter进行保活

// noop

// not over tcp

func (session *ClientCommandSession) connect(rawUrl string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// # 建立连接

func (session *ClientCommandSession) writeOptions() error { _ = "STUB: not implemented"; return nil }

func (session *ClientCommandSession) writeDescribe() error { _ = "STUB: not implemented"; return nil }

func (session *ClientCommandSession) writeAnnounce() error { _ = "STUB: not implemented"; return nil }

func (session *ClientCommandSession) writeSetup() error { _ = "STUB: not implemented"; return nil }

// 461情况下尝试切换UDP重试

// 461情况尝试切换TCP重试

// can't else if

func (session *ClientCommandSession) writeOneSetup(setupUri string) error {
	_ = "STUB: not implemented"
	return nil
}

// 切换transport尝试继续

// 增强兼容性逻辑
// 有用户反馈，存在对端不返回server_port的情况，对端是easydrawin
// 其实在pull的情况下，没有对端端口也可以，因为不发数据，或者需要发送时，使用接收时获取到的对端地址即可

// noop

func (session *ClientCommandSession) writeOneSetupTcp(setupUri string) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO chef: 这里没有解析回传的channel id了，因为我假定了它和request中的是一致的

func (session *ClientCommandSession) writePlay() error { _ = "STUB: not implemented"; return nil }

func (session *ClientCommandSession) writeRecord() error { _ = "STUB: not implemented"; return nil }

func (session *ClientCommandSession) writeCmd(method, uri string, headers map[string]string, body string) error {
	_ = "STUB: not implemented"
	return nil
}

// 鉴权时固定用RawUrlWithoutUserInfo

//Log.Debugf("[%s] > write %s.", session.uniqueKey, method)

// @param headers 可以为nil
// @param body 可以为空
func (session *ClientCommandSession) writeCmdReadResp(method, uri string, headers map[string]string, body string) (ctx nazahttp.HttpRespMsgCtx, err error) {
	_ = "STUB: not implemented"
	return *new(nazahttp.HttpRespMsgCtx), nil
}

func (session *ClientCommandSession) dispose(err error) error {
	_ = "STUB: not implemented"
	return nil
}
