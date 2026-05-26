// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package httpflv

import (
	"io"
)

type TagHeader struct {
	Type      uint8  // type
	DataSize  uint32 // body大小，不包含 header 和 prev tag size 字段
	Timestamp uint32 // 绝对时间戳，单位毫秒
	StreamId  uint32 // always 0
}

type Tag struct {
	Header TagHeader
	Raw    []byte // 结构为 (11字节的 tag header) + (body) + (4字节的 prev tag size)
}

// Payload 只包含数据部分，去除了前面11字节的tag header和后面4字节的prev tag size
func (tag *Tag) Payload() []byte { _ = "STUB: not implemented"; return nil }

func (tag *Tag) IsMetadata() bool { _ = "STUB: not implemented"; return false }

func (tag *Tag) IsAvc() bool { _ = "STUB: not implemented"; return false }

func (tag *Tag) IsHevc() bool { _ = "STUB: not implemented"; return false }

func (tag *Tag) IsAvcKeySeqHeader() bool { _ = "STUB: not implemented"; return false }

func (tag *Tag) IsHevcKeySeqHeader() bool { _ = "STUB: not implemented"; return false }

// IsVideoKeySeqHeader AVC或HEVC的seq header
func (tag *Tag) IsVideoKeySeqHeader() bool { _ = "STUB: not implemented"; return false }

func (tag *Tag) IsAvcKeyNalu() bool { _ = "STUB: not implemented"; return false }

func (tag *Tag) IsHevcKeyNalu() bool { _ = "STUB: not implemented"; return false }

// IsVideoKeyNalu AVC或HEVC的关键帧
func (tag *Tag) IsVideoKeyNalu() bool { _ = "STUB: not implemented"; return false }

func (tag *Tag) IsAacSeqHeader() bool { _ = "STUB: not implemented"; return false }

func (tag *Tag) clone() (out Tag) { _ = "STUB: not implemented"; return *new(Tag) }

func (tag *Tag) ModTagTimestamp(timestamp uint32) { _ = "STUB: not implemented"; return }

// PackHttpflvTag 打包一个序列化后的 tag 二进制buffer，包含 tag header，body，prev tag size
func PackHttpflvTag(t uint8, timestamp uint32, in []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// ReadTag 从`rd`中读取数据并解析至`tag`
func ReadTag(rd io.Reader) (tag Tag, err error) { _ = "STUB: not implemented"; return *new(Tag), nil }

func parseTagHeader(rawHeader []byte) TagHeader { _ = "STUB: not implemented"; return *new(TagHeader) }
