// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package httpflv

import (
	"context"
	"net/http"
	"sync"

	"github.com/q191201771/lal/pkg/base"

	"github.com/q191201771/naza/pkg/connection"
)

type PullSessionOption struct {
	// 从调用Pull函数，到接收音视频数据的前一步，也即发送完HTTP请求的超时时间
	// 如果为0，则没有超时时间
	PullTimeoutMs int

	ReadTimeoutMs int // 接收数据超时，单位毫秒，如果为0，则不设置超时
}

var defaultPullSessionOption = PullSessionOption{
	PullTimeoutMs: 10000,
	ReadTimeoutMs: 0,
}

type PullSession struct {
	option PullSessionOption // const after ctor

	conn        connection.Connection
	sessionStat base.BasicSessionStat

	onReadFlvTag OnReadFlvTag

	urlCtx base.UrlContext

	disposeOnce sync.Once
}

type ModPullSessionOption func(option *PullSessionOption)

func NewPullSession(modOptions ...ModPullSessionOption) *PullSession {
	_ = "STUB: not implemented"
	return nil
}

// OnReadFlvTag @param tag: 底层保证回调上来的Raw数据长度是完整的（但是不会分析Raw内部的编码数据）
type OnReadFlvTag func(tag Tag)

// WithOnReadFlvTag
//
// @param onReadFlvTag 读取到 flv tag 数据时回调。回调结束后，PullSession 不会再使用这块 <tag> 数据。
func (session *PullSession) WithOnReadFlvTag(onReadFlvTag OnReadFlvTag) *PullSession {
	_ = "STUB: not implemented"
	return nil
}

// Start 阻塞直到和对端完成拉流前，握手部分的工作，或者发生错误。
//
// 注意，握手指的是发送完HTTP Request，不包含接收任何数据，因为有的httpflv服务端，如果流不存在不会发送任何内容，此时我们也应该认为是握手完成了。
//
// @param rawUrl 支持如下两种格式（当然，关键点是对端支持）：
//  1. `http://{domain}/{app_name}/{stream_name}.flv`
//  2. `http://{ip}/{domain}/{app_name}/{stream_name}.flv`
func (session *PullSession) Start(rawUrl string) error { _ = "STUB: not implemented"; return nil }

// Pull deprecated. use WithOnReadFlvTag and Start instead.
func (session *PullSession) Pull(rawUrl string, onReadFlvTag OnReadFlvTag) error {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
// IClientSessionLifecycle interface
// ---------------------------------------------------------------------------------------------------------------------

// Dispose 文档请参考： IClientSessionLifecycle interface
func (session *PullSession) Dispose() error { _ = "STUB: not implemented"; return nil }

// WaitChan 文档请参考： IClientSessionLifecycle interface
func (session *PullSession) WaitChan() <-chan error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionUrlContext interface
// ---------------------------------------------------------------------------------------------------------------------

// Url 文档请参考： interface ISessionUrlContext
func (session *PullSession) Url() string { _ = "STUB: not implemented"; return "" }

// AppName 文档请参考： interface ISessionUrlContext
func (session *PullSession) AppName() string { _ = "STUB: not implemented"; return "" }

// StreamName 文档请参考： interface ISessionUrlContext
func (session *PullSession) StreamName() string { _ = "STUB: not implemented"; return "" }

// RawQuery 文档请参考： interface ISessionUrlContext
func (session *PullSession) RawQuery() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------
// IObject interface
// ---------------------------------------------------------------------------------------------------------------------

// UniqueKey 文档请参考： interface IObject
func (session *PullSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionStat interface
// ---------------------------------------------------------------------------------------------------------------------

// UpdateStat 文档请参考： interface ISessionStat
func (session *PullSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

// GetStat 文档请参考： interface ISessionStat
func (session *PullSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *new(base.StatSession)
}

// IsAlive 文档请参考： interface ISessionStat
func (session *PullSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

// ---------------------------------------------------------------------------------------------------------------------

func (session *PullSession) pull(rawUrl string) error { _ = "STUB: not implemented"; return nil }

func (session *PullSession) pullContext(ctx context.Context, rawUrl string, onReadFlvTag OnReadFlvTag) error {
	_ = "STUB: not implemented"
	return nil
}

// 异步握手

// 处理跳转

// 等待握手结果，或者超时通知

// 注意，如果超时，可能连接已经建立了，要dispose避免泄漏

// 握手消息，不为nil则握手失败

// 握手成功，开启收数据协程

func (session *PullSession) connect(rawUrl string) (err error) {
	_ = "STUB: not implemented"
	// TODO(chef): refactor 可以考虑抽象出一个http client，负责http拉流的建连、https、302等功能
	return nil
}

// TODO chef: 为什么是 Read 赋值给 Write

func (session *PullSession) writeHttpRequest() error {
	_ = "STUB: not implemented"
	// # 发送 http GET 请求
	return nil
}

func (session *PullSession) readHttpRespHeader() (statusCode string, headers http.Header, err error) {
	_ = "STUB: not implemented"
	return "", *new(http.Header), nil
}

func (session *PullSession) readFlvHeader() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO chef: check flv header's value

func (session *PullSession) readTag() (Tag, error) {
	_ = "STUB: not implemented"
	return *new(Tag), nil
}

func (session *PullSession) runReadLoop(onReadFlvTag OnReadFlvTag) {
	_ = "STUB: not implemented"
	return
}

func (session *PullSession) dispose(err error) error { _ = "STUB: not implemented"; return nil }
