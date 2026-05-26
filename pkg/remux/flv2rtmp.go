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
	"github.com/q191201771/lal/pkg/httpflv"
)

func FlvTagHeader2RtmpHeader(in httpflv.TagHeader) (out base.RtmpHeader) {
	_ = "STUB: not implemented"
	return *new(base.RtmpHeader)
}

// FlvTag2RtmpMsg @return msg: 返回的内存块引用参数`tag`的内存块
func FlvTag2RtmpMsg(tag httpflv.Tag) (msg base.RtmpMsg) {
	_ = "STUB: not implemented"
	return *new(base.RtmpMsg)
}

// FlvTag2RtmpChunks @return 返回的内存块为内部新申请
func FlvTag2RtmpChunks(tag httpflv.Tag) []byte { _ = "STUB: not implemented"; return nil }
