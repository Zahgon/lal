// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtmp

import (
	"github.com/q191201771/naza/pkg/nazabytes"

	"github.com/q191201771/lal/pkg/base"
)

// ----- Stream --------------------------------------------------------------------------------------------------------

type Stream struct {
	header base.RtmpHeader
	msg    StreamMsg

	absTsFlag bool   // 标记当这个stream收到新的msg的时候，是否收到过绝对时间
	timestamp uint32 // 注意，是rtmp chunk协议header中的时间戳，可能是绝对的，也可能是相对的。上层不应该使用这个字段，而应该使用Header.TimestampAbs
}

func NewStream() *Stream { _ = "STUB: not implemented"; return nil }

// 序列化成可读字符串，一般用于发生错误时打印日志
func (stream *Stream) toDebugString() string { _ = "STUB: not implemented"; return "" }

func (stream *Stream) toAvMsg() base.RtmpMsg {
	_ = "STUB: not implemented"
	// TODO chef: 考虑可能出现header中的len和buf的大小不一致的情况
	return *new(base.RtmpMsg)
}

// ----- StreamMsg -----------------------------------------------------------------------------------------------------

type StreamMsg struct {
	// TODO(chef): [refactor] 考虑外部(chunk_composer)不要直接访问buff，封装一层 202206
	buff *nazabytes.Buffer
}

// Grow 确保可写空间，如果不够会扩容
func (msg *StreamMsg) Grow(n uint32) { _ = "STUB: not implemented"; return }

func (msg *StreamMsg) Len() uint32 { _ = "STUB: not implemented"; return 0 }

func (msg *StreamMsg) Flush(n uint32) { _ = "STUB: not implemented"; return }

func (msg *StreamMsg) Skip(n uint32) { _ = "STUB: not implemented"; return }

func (msg *StreamMsg) Reset() { _ = "STUB: not implemented"; return }

func (msg *StreamMsg) ResetAndFree() { _ = "STUB: not implemented"; return }

func (msg *StreamMsg) peekStringWithType() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (msg *StreamMsg) readStringWithType() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (msg *StreamMsg) readNumberWithType() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (msg *StreamMsg) readObjectWithType() (ObjectPairArray, error) {
	_ = "STUB: not implemented"
	return *new(ObjectPairArray), nil
}

func (msg *StreamMsg) readNull() error { _ = "STUB: not implemented"; return nil }
