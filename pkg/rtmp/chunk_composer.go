// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtmp

import (
	"io"
)

// ChunkComposer
//
// 读取chunk，并合并chunk，生成message返回给上层
type ChunkComposer struct {
	peerChunkSize   uint32
	reuseBufferFlag bool // TODO(chef): [fix] RtmpTypeIdAggregateMessage时，reuseBufferFlag==false的处理 202206

	csid2stream map[int]*Stream
}

func NewChunkComposer() *ChunkComposer { _ = "STUB: not implemented"; return nil }

func (c *ChunkComposer) SetReuseBufferFlag(val bool) { _ = "STUB: not implemented"; return }

func (c *ChunkComposer) SetPeerChunkSize(val uint32) { _ = "STUB: not implemented"; return }

type OnCompleteMessage func(stream *Stream) error

// RunLoop 将rtmp chunk合并为message
//
// @param cb:
//
//	 @param cb.Stream.msg:
//	  注意，回调结束后，`msg`的内存块会被`ChunkComposer`重复使用。
//		 也即多次回调的`msg`是复用的同一块内存块。
//		 如果业务方需要在回调结束后，依然持有`msg`，那么需要对`msg`进行拷贝。
//		 只在回调中使用`msg`，则不需要拷贝。
//		 @return(回调函数`cb`的返回值): 如果cb返回的error不为nil，则`RunLoop`停止阻塞，并返回这个错误。
//
// @return 阻塞直到发生错误
//
// TODO chef: msglen支持最大阈值，超过可以认为对端是非法的
func (c *ChunkComposer) RunLoop(reader io.Reader, cb OnCompleteMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// 5.3.1.1. Chunk Basic Header
// 读取fmt和csid

// csid可能是变长的

// noop

// 5.3.1.2. Chunk Message Header
// 当前chunk的fmt不同，Message Header包含的字段也不同，是变长

// 包头中为绝对时间戳

// 包头中为相对时间戳

//stream.header.TimestampAbs += stream.header.Timestamp

// 包头中为相对时间戳

//stream.header.TimestampAbs += stream.header.Timestamp

// noop

// 5.3.1.3 Extended Timestamp
// 使用ffmpeg推流时，发现时间戳超过3字节最大值后，即使是fmt3(即包头大小为0)，依然存在ext ts字段
// 所以这里我将 `==` 的判断改成了 `>=`
// TODO chef:
// - 测试其他客户端和ext ts相关的表现
// - 这部分可能还有问题，需要根据具体的case调整
//if stream.header.Timestamp == maxTimestampInMessageHeader {

// noop

// 对端设置了chunk size

// 这么处理相当于取最后一个chunk的时间戳差值，有的协议栈是取的第一个，正常来说都可以

// 懒初始化

// 读取sub message的头

// 计算时间戳

// message包体

// sub message回调给上层

// 跳过prev size字段

// TODO(chef): 这里应该永远执行不到，可以删除掉

func (c *ChunkComposer) getOrCreateStream(csid int) *Stream { _ = "STUB: not implemented"; return nil }

// 临时存放一些rtmp推流case在这，便于理解，以及修改后，回归用
//
// 场景：ffmpeg推送test.flv至lalserver
// 关注点：message超过chunk时，fmt和timestamp的值
//
// ChunkComposer chunk fmt:1 header:{Csid:6 MsgLen:143 Timestamp:40 MsgTypeId:9 MsgStreamId:1 TimestampAbs:520} csid:6 len:143 ts:520
// ChunkComposer chunk fmt:1 header:{Csid:6 MsgLen:4511 Timestamp:40 MsgTypeId:9 MsgStreamId:1 TimestampAbs:560} csid:6 len:4511 ts:560
// ChunkComposer chunk fmt:3 header:{Csid:6 MsgLen:4511 Timestamp:40 MsgTypeId:9 MsgStreamId:1 TimestampAbs:560} csid:6 len:4511 ts:560
// 此处应只给上层返回一次，也即一个message，时间戳应该是560
// ChunkComposer chunk fmt:1 header:{Csid:6 MsgLen:904 Timestamp:40 MsgTypeId:9 MsgStreamId:1 TimestampAbs:600} csid:6 len:904 ts:600
