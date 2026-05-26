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

type PushSession struct {
	IsFresh bool

	core *ClientSession
}

type PushSessionOption struct {
	// 从调用Push函数，到可以发送音视频数据的前一步，也即收到服务端返回的rtmp publish对应结果的信令的超时时间
	// 如果为0，则没有超时时间
	PushTimeoutMs int

	WriteAvTimeoutMs int
	WriteBufSize     int // io层发送音视频数据的缓冲大小，如果为0，则没有缓冲
	WriteChanSize    int // io层发送音视频数据的异步队列大小，如果为0，则同步发送

	HandshakeComplexFlag bool
	// TlsConfig
	// rtmps时使用。
	// 不关心可以不填。
	// 业务方可以通过这个字段自定义 tls.Config
	// 注意，如果使用rtmps并且该字段为nil，那么内部会使用 base.DefaultTlsConfigClient 生成 tls.Config
	TlsConfig *tls.Config
}

var defaultPushSessionOption = PushSessionOption{
	PushTimeoutMs:        10000,
	WriteAvTimeoutMs:     0,
	WriteBufSize:         0,
	WriteChanSize:        0,
	HandshakeComplexFlag: false,
}

type ModPushSessionOption func(option *PushSessionOption)

func NewPushSession(modOptions ...ModPushSessionOption) *PushSession {
	_ = "STUB: not implemented"
	return nil
}

// Start 阻塞直到和对端完成推流前，握手部分的工作（也即收到RTMP Publish response），或者发生错误
func (s *PushSession) Start(rawUrl string) error { _ = "STUB: not implemented"; return nil }

// Push deprecated. use Start instead.
func (s *PushSession) Push(rawUrl string) error { _ = "STUB: not implemented"; return nil }

// Write 发送数据
//
// @param b:
//
//	注意，`b`数据应该是已经打包成rtmp chunk格式的数据。这里的数据就对应socket发送的数据，内部不会再修改数据内容。
//	如果要发送 base.RtmpMsg 数据，请使用 WriteMsg 函数。
func (s *PushSession) Write(b []byte) error {
	_ = "STUB: not implemented"
	// TODO(chef): [opt] 使用Write函数时确保metadata有@SetDataFrame 202207
	return nil
}

// WriteMsg
//
// 内部会根据 msg 的包头字段和包体数据，打包成 rtmp chunk 格式的数据，然后发送。
// 如果想要自己控制打包过程，请使用 Write 函数直接发送数据。
func (s *PushSession) WriteMsg(msg base.RtmpMsg) error { _ = "STUB: not implemented"; return nil }

// Flush 将缓存的数据立即刷新发送
// 是否有缓存策略，请参见配置及内部实现
func (s *PushSession) Flush() error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------
// IClientSessionLifecycle interface
// ---------------------------------------------------------------------------------------------------------------------

// Dispose 文档请参考： IClientSessionLifecycle interface
func (s *PushSession) Dispose() error { _ = "STUB: not implemented"; return nil }

// WaitChan 文档请参考： IClientSessionLifecycle interface
func (s *PushSession) WaitChan() <-chan error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionUrlContext interface
// ---------------------------------------------------------------------------------------------------------------------

// Url 文档请参考： interface ISessionUrlContext
func (s *PushSession) Url() string { _ = "STUB: not implemented"; return "" }

// AppName 文档请参考： interface ISessionUrlContext
func (s *PushSession) AppName() string { _ = "STUB: not implemented"; return "" }

// StreamName 文档请参考： interface ISessionUrlContext
func (s *PushSession) StreamName() string { _ = "STUB: not implemented"; return "" }

// RawQuery 文档请参考： interface ISessionUrlContext
func (s *PushSession) RawQuery() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------
// IObject interface
// ---------------------------------------------------------------------------------------------------------------------

// UniqueKey 文档请参考： interface IObject
func (s *PushSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionStat interface
// ---------------------------------------------------------------------------------------------------------------------

// GetStat 文档请参考： interface ISessionStat
func (s *PushSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *

	// UpdateStat 文档请参考： interface ISessionStat
	new(base.StatSession)
}

func (s *PushSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

// IsAlive 文档请参考： interface ISessionStat
func (s *PushSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}
