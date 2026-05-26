// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtmp

import (
	"crypto/tls"

	"github.com/q191201771/lal/pkg/base"
)

type OnReadRtmpAvMsg func(msg base.RtmpMsg)

type PullSession struct {
	core *ClientSession
}

type PullSessionOption struct {
	// PullTimeoutMs
	//
	// 从调用Pull函数，到接收音视频数据的前一步，也即收到服务端返回的rtmp play对应结果的信令的超时时间
	// 如果为0，则没有超时时间
	//
	PullTimeoutMs int

	ReadAvTimeoutMs            int
	ReadBufSize                int  // io层读取音视频数据时的缓冲大小，如果为0，则没有缓冲
	ReuseReadMessageBufferFlag bool // 接收Message时，是否复用内存块
	PeerWinAckSize             int

	HandshakeComplexFlag bool
	// TlsConfig
	// rtmps时使用。
	// 不关心可以不填。
	// 业务方可以通过这个字段自定义 tls.Config
	// 注意，如果使用rtmps并且该字段为nil，那么内部会使用 base.DefaultTlsConfigClient 生成 tls.Config
	TlsConfig *tls.Config
}

var defaultPullSessionOption = PullSessionOption{
	PullTimeoutMs:              10000,
	ReadAvTimeoutMs:            0,
	ReadBufSize:                0,
	HandshakeComplexFlag:       false,
	PeerWinAckSize:             0,
	ReuseReadMessageBufferFlag: true,
}

type ModPullSessionOption func(option *PullSessionOption)

func NewPullSession(modOptions ...ModPullSessionOption) *PullSession {
	_ = "STUB: not implemented"
	return nil
}

// WithOnPullSucc Pull成功
//
// 如果你想保证绝对时序，在 WithOnReadRtmpAvMsg 回调音视频数据前，做一些操作，那么使用这个回调替代 Start 返回成功
func (s *PullSession) WithOnPullSucc(onPullResult func()) *PullSession {
	_ = "STUB: not implemented"
	return nil
}

// WithOnReadRtmpAvMsg
//
// @param onReadRtmpAvMsg:
//
//	msg: 关于内存块的说明：
//	  ReuseReadMessageBufferFlag 为true时：
//	    回调结束后，`msg`的内存块会被`PullSession`重复使用。
//	    也即多次回调的`msg`是复用的同一块内存块。
//	    如果业务方需要在回调结束后，依然持有`msg`，那么需要对`msg`进行拷贝，比如调用`msg.Clone()`。
//	    只在回调中使用`msg`，则不需要拷贝。
//	  ReuseReadMessageBufferFlag 为false时：
//	    回调接收后，`PullSession`不再使用该内存块。
//	    业务方可以自由持有释放该内存块。
func (s *PullSession) WithOnReadRtmpAvMsg(onReadRtmpAvMsg OnReadRtmpAvMsg) *PullSession {
	_ = "STUB: not implemented"
	return nil
}

// Start 阻塞直到和对端完成拉流前的所有准备工作（也即收到RTMP Play response），或者发生错误
func (s *PullSession) Start(rawUrl string) error { _ = "STUB: not implemented"; return nil }

// Pull deprecated. use Start instead.
func (s *PullSession) Pull(rawUrl string) error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------
// IClientSessionLifecycle interface
// ---------------------------------------------------------------------------------------------------------------------

// Dispose 文档请参考： IClientSessionLifecycle interface
func (s *PullSession) Dispose() error { _ = "STUB: not implemented"; return nil }

// WaitChan 文档请参考： IClientSessionLifecycle interface
func (s *PullSession) WaitChan() <-chan error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionUrlContext interface
// ---------------------------------------------------------------------------------------------------------------------

// Url 文档请参考： interface ISessionUrlContext
func (s *PullSession) Url() string { _ = "STUB: not implemented"; return "" }

// AppName 文档请参考： interface ISessionUrlContext
func (s *PullSession) AppName() string { _ = "STUB: not implemented"; return "" }

// StreamName 文档请参考： interface ISessionUrlContext
func (s *PullSession) StreamName() string { _ = "STUB: not implemented"; return "" }

// RawQuery 文档请参考： interface ISessionUrlContext
func (s *PullSession) RawQuery() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------
// IObject interface
// ---------------------------------------------------------------------------------------------------------------------

// UniqueKey 文档请参考： interface IObject
func (s *PullSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionStat interface
// ---------------------------------------------------------------------------------------------------------------------

// GetStat 文档请参考： interface ISessionStat
func (s *PullSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *

	// UpdateStat 文档请参考： interface ISessionStat
	new(base.StatSession)
}

func (s *PullSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

// IsAlive 文档请参考： interface ISessionStat
func (s *PullSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}
