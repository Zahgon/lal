// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtprtcp

type RtpUnpackContainer struct {
	unpackerProtocol IRtpUnpackerProtocol

	list RtpPacketList
}

func NewRtpUnpackContainer(maxSize int, unpackerProtocol IRtpUnpackerProtocol) *RtpUnpackContainer {
	_ = "STUB: not implemented"
	return nil
}

// Feed 输入收到的rtp包
func (r *RtpUnpackContainer) Feed(pkt RtpPacket) {
	_ = "STUB: not implemented"
	// 过期的包
	return
}

// 计算位置

// 根据序号插入有序链表

// 尽可能多的合成顺序的帧

// 合成顺序的帧成功了，直接返回

// 缓存达到最大值

// 尝试合成一帧发生跳跃的帧

// 合成失败了，丢弃一包过期数据

// 合成成功了，再次尝试，尽可能多的合成顺序的帧

// tryUnpackOneSequential 从队列头部，尝试合成一个完整的帧。保证这次合成的帧的首个seq和上次合成帧的尾部seq是连续的
func (r *RtpUnpackContainer) tryUnpackOneSequential() bool { _ = "STUB: not implemented"; return false }

// tryUnpackOne 从队列头部，尝试合成一个完整的帧。不保证这次合成的帧的首个seq和上次合成帧的尾部seq是连续的
func (r *RtpUnpackContainer) tryUnpackOne() bool { _ = "STUB: not implemented"; return false }
