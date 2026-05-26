// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtmp

import (
	"context"
	"crypto/tls"
	"sync"

	"github.com/q191201771/lal/pkg/base"

	"github.com/q191201771/naza/pkg/connection"
)

// ClientSession rtmp 客户端类型连接的底层实现
// package rtmp 的使用者应该优先使用基于 ClientSession 实现的 PushSession 和 PullSession
type ClientSession struct {
	onDoResult func()

	// 只有PullSession使用
	onReadRtmpAvMsg OnReadRtmpAvMsg

	option ClientSessionOption

	packer        *MessagePacker
	chunkComposer *ChunkComposer
	urlCtx        base.UrlContext
	hc            IHandshakeClient

	conn                  connection.Connection
	doResultChan          chan struct{}
	errChan               chan error
	hasNotifyDoResultSucc bool

	sessionStat base.BasicSessionStat

	debugLogReadUserCtrlMsgCount int
	debugLogReadUserCtrlMsgMax   int

	recvLastAck uint64
	seqNum      uint32

	disposeOnce sync.Once
	authInfo    AuthInfo
}

type AuthInfo struct {
	challenge string
	salt      string
	opaque    string
}

type ClientSessionOption struct {
	// 单位毫秒，如果为0，则没有超时
	DoTimeoutMs      int // 从发起连接（包含了建立连接的时间）到收到publish或play信令结果的超时
	ReadAvTimeoutMs  int // 读取音视频数据的超时
	WriteAvTimeoutMs int // 发送音视频数据的超时

	ReadBufSize   int // io层读取音视频数据时的缓冲大小，如果为0，则没有缓冲
	WriteBufSize  int // io层发送音视频数据的缓冲大小，如果为0，则没有缓冲
	WriteChanSize int // io层发送音视频数据的异步队列大小，如果为0，则同步发送

	ReuseReadMessageBufferFlag bool // 接收Message时，是否重用内存块
	// PeerWinAckSize
	//
	// 设置发送 base.RtmpTypeIdAck 的触发阈值的默认值。
	// 如果没收到 base.RtmpTypeIdWinAckSize，则使用该默认值作为阈值；
	// 如果收到 base.RtmpTypeIdWinAckSize，则使用收到的值作为阈值。
	PeerWinAckSize int

	HandshakeComplexFlag bool // 握手是否使用复杂模式
	// TlsConfig
	// rtmps时使用。
	// 不关心可以不填。
	// 业务方可以通过这个字段自定义 tls.Config
	// 注意，如果使用rtmps并且该字段为nil，那么内部会使用 base.DefaultTlsConfigClient 生成 tls.Config
	TlsConfig *tls.Config
}

var defaultClientSessOption = ClientSessionOption{
	DoTimeoutMs:                10000,
	ReadAvTimeoutMs:            0,
	WriteAvTimeoutMs:           0,
	ReadBufSize:                0,
	WriteBufSize:               0,
	WriteChanSize:              0,
	HandshakeComplexFlag:       false,
	PeerWinAckSize:             0,
	ReuseReadMessageBufferFlag: true,
}

type ModClientSessionOption func(option *ClientSessionOption)

// NewClientSession @param t: session的类型，只能是推或者拉
func NewClientSession(sessionType base.SessionType, modOptions ...ModClientSessionOption) *ClientSession {
	_ = "STUB: not implemented"
	return nil
}

// Start 阻塞直到收到服务端返回的 publish / play 对应结果的信令或者发生错误
func (s *ClientSession) Start(rawUrl string) error { _ = "STUB: not implemented"; return nil }

func (s *ClientSession) Write(msg []byte) error { _ = "STUB: not implemented"; return nil }

func (s *ClientSession) Flush() error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------
// IClientSessionLifecycle interface
// ---------------------------------------------------------------------------------------------------------------------

// Dispose 文档请参考： IClientSessionLifecycle interface
func (s *ClientSession) Dispose() error { _ = "STUB: not implemented"; return nil }

// WaitChan 文档请参考： IClientSessionLifecycle interface
func (s *ClientSession) WaitChan() <-chan error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------

func (s *ClientSession) Url() string { _ = "STUB: not implemented"; return "" }

func (s *ClientSession) AppName() string { _ = "STUB: not implemented"; return "" }

func (s *ClientSession) StreamName() string { _ = "STUB: not implemented"; return "" }

func (s *ClientSession) RawQuery() string { _ = "STUB: not implemented"; return "" }

func (s *ClientSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// ----- ISessionStat --------------------------------------------------------------------------------------------------

func (s *ClientSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *new(base.StatSession)
}

func (s *ClientSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

func (s *ClientSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *ClientSession) connect() { _ = "STUB: not implemented"; return }

func (s *ClientSession) doContext(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *ClientSession) parseUrl(rawUrl string) (err error) { _ = "STUB: not implemented"; return nil }

func (s *ClientSession) tcUrl() string { _ = "STUB: not implemented"; return "" }

func (s *ClientSession) appName() string { _ = "STUB: not implemented"; return "" }

func (s *ClientSession) streamNameWithRawQuery() string { _ = "STUB: not implemented"; return "" }

func (s *ClientSession) tcpConnect() error { _ = "STUB: not implemented"; return nil }

func (s *ClientSession) handshake() error { _ = "STUB: not implemented"; return nil }

func (s *ClientSession) runReadLoop() { _ = "STUB: not implemented"; return }

func (s *ClientSession) doMsg(stream *Stream) error { _ = "STUB: not implemented"; return nil }

func (s *ClientSession) doAck(stream *Stream) error { _ = "STUB: not implemented"; return nil }

func (s *ClientSession) doUserControl(stream *Stream) error { _ = "STUB: not implemented"; return nil }

func (s *ClientSession) doDataMessageAmf0(stream *Stream) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClientSession) doCommandMessage(stream *Stream) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClientSession) doErrorMessage(stream *Stream, tid int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClientSession) parseAuthorityInfo(auth string) {
	_ = "STUB: not implemented"
	// 解析salt、challenge、opaque字段
	return
}

func (s *ClientSession) dealErrorMessage(description string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// app和tcUrl需要加上streamid、authmod、user

//关闭上一次连接并发起新的连接

// base64(md5(username|salt|password))作为新的salt1

// response = base64(md5(salt1|opaque|challenge))

// app和tcUrl需要加上challenge、response、opaque字段

// 关闭前一个连接并发起新的连接

func (s *ClientSession) doOnStatusMessage(stream *Stream, tid int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClientSession) doResultMessage(stream *Stream, tid int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClientSession) doProtocolControlMessage(stream *Stream) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO chef: 是否需要关注这个信令

// composer内部会自动更新peer chunk size.

func (s *ClientSession) writeAcknowledgementIfNeeded(stream *Stream) error {
	_ = "STUB: not implemented"
	// https://github.com/q191201771/lal/pull/154
	return nil
}

//此次接收小于窗口大小一半，不处理

//当序列号溢出时，将其重置

//时间戳暂时先发0

func (s *ClientSession) notifyDoResultSucc() {
	_ = "STUB: not implemented"
	// 碰上过对端服务器实现有问题，对于play信令回复了两次相同的结果，我们在这里忽略掉非第一次的回复
	return
}

//pull有可能还需要小包发送，不使用缓存

func (s *ClientSession) dispose(err error) error { _ = "STUB: not implemented"; return nil }

func defaultOnPullResult() { _ = "STUB: not implemented"; return }

func defaultOnReadRtmpAvMsg(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }
