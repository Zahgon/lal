// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package remux

import (
	"github.com/q191201771/lal/pkg/base"
)

// MakeDefaultRtmpHeader
//
// 使用场景：一般是输入流转换为输出流时。
// 目的：使得流格式更标准。
// 做法：设置 MsgStreamId 和 Csid，其他字段保持`in`的值。
func MakeDefaultRtmpHeader(in base.RtmpHeader) (out base.RtmpHeader) {
	_ = "STUB: not implemented"
	return *new(base.RtmpHeader)
}

// ---------------------------------------------------------------------------------------------------------------------

// LazyRtmpChunkDivider 在必要时，有且仅有一次做切分成chunk的操作
type LazyRtmpChunkDivider struct {
	msg              base.RtmpMsg
	chunksWithSdf    []byte
	chunksWithoutSdf []byte
}

func (lcd *LazyRtmpChunkDivider) Init(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

func (lcd *LazyRtmpChunkDivider) GetEnsureWithSdf() []byte { _ = "STUB: not implemented"; return nil }

func (lcd *LazyRtmpChunkDivider) GetEnsureWithoutSdf() []byte {
	_ = "STUB: not implemented"
	return nil
}
