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

// RtmpMsg2FlvTag @return 返回的内存块为新申请的独立内存块
func RtmpMsg2FlvTag(msg base.RtmpMsg) *httpflv.Tag { _ = "STUB: not implemented"; return nil }

// -------------------------------------------------------------------------------------------------------------------

// LazyRtmpMsg2FlvTag 在必要时，有且仅有一次做转换操作
type LazyRtmpMsg2FlvTag struct {
	msg base.RtmpMsg
	//tagWithSdf []byte
	tagWithoutSdf []byte
}

func (l *LazyRtmpMsg2FlvTag) Init(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

func (l *LazyRtmpMsg2FlvTag) GetEnsureWithSdf() []byte {
	_ = "STUB: not implemented"
	// TODO(chef): [refactor] 这个函数目前没有实际用途 202207
	return nil
}

func (l *LazyRtmpMsg2FlvTag) GetEnsureWithoutSdf() []byte { _ = "STUB: not implemented"; return nil }
