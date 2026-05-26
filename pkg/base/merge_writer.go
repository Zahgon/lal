// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package base

import (
	"net"
)

// TODO(chef): feat 通过时间戳（目前是数据大小）来设定合并阈值

// MergeWriter 合并多个内存块，达到阈值后一次性将内存块数组返回给上层
//
// 注意，输入时的单个内存块，回调时不会出现拆分切割的情况
type MergeWriter struct {
	onWritev OnWritev
	size     int

	currSize int
	bs       net.Buffers // TODO(chef): perf 复用外层切片
}

type OnWritev func(bs net.Buffers)

// NewMergeWriter
//
// @param onWritev 回调缓存的1~n个内存块
// @param size     回调阈值
func NewMergeWriter(onWritev OnWritev, size int) *MergeWriter {
	_ = "STUB: not implemented"
	return nil
}

// Write
//
// 注意，函数调用结束后，`b`内存块会被内部持有
func (w *MergeWriter) Write(b []byte) { _ = "STUB: not implemented"; return }

// Flush 强制将内部缓冲的数据全部回调排空
func (w *MergeWriter) Flush() { _ = "STUB: not implemented"; return }

// flush 将内部缓冲的数据全部回调排空
func (w *MergeWriter) flush() {
	_ = "STUB: not implemented"
	// only for debug log
	return
}
