// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtprtcp

type RtpPacketListItem struct {
	Packet RtpPacket
	Next   *RtpPacketListItem
}

// RtpPacketList rtp packet的有序链表，前面的seq小于后面的seq
//
// 为什么不用红黑树等查找性能更高的有序kv结构？
// 第一，容器有最大值，这个数量级用啥容器都差不多，
// 第二，插入时，99.99%的seq号是当前最大号附近的，遍历找就可以了，
// 注意，这个链表并不是一个定长容器，当数据有序时，容器内缓存的数据是一个帧的数据。
type RtpPacketList struct {
	// TODO(chef): [refactor] 隐藏这两个变量的访问权限 202207
	Head RtpPacketListItem // 哨兵，自身不存放rtp包，第一个rtp包存在在head.next中
	Size int               // 实际元素个数

	doneSeqFlag bool   // 如果为false，则说明我们是初始化阶段，还不知道需要的packet的seq是多少
	doneSeq     uint16 // 已处理的seq号，之后我们需要seq+1的packet

	maxSize int
}

// IsStale 是否过期
func (l *RtpPacketList) IsStale(seq uint16) bool { _ = "STUB: not implemented"; return false }

// 序号太小

// Insert 插入有序链表，并去重
func (l *RtpPacketList) Insert(pkt RtpPacket) {
	_ = "STUB: not implemented"
	// 遍历查找插入位置
	return
}

// TODO(chef): [perf] 考虑优化成从后往前查找，提高查找效率 202207

// 包已经存在，不需要插入了

// noop

// PopFirst 弹出第一个包。注意，调用方保证容器不为空时调用
func (l *RtpPacketList) PopFirst() RtpPacket { _ = "STUB: not implemented"; return *new(RtpPacket) }

// PeekFirst 查看第一个包。注意，调用方保证容器不为空时调用
func (l *RtpPacketList) PeekFirst() RtpPacket {
	_ = "STUB: not implemented"
	return *

	// InitMaxSize 设置容器最大容量
	new(RtpPacket)
}

func (l *RtpPacketList) InitMaxSize(maxSize int) { _ = "STUB: not implemented"; return }

// Full 是否已经满了
func (l *RtpPacketList) Full() bool { _ = "STUB: not implemented"; return false }

// IsFirstSequential 第一个包是否是需要的（与之前已处理的是连续的）
func (l *RtpPacketList) IsFirstSequential() bool { _ = "STUB: not implemented"; return false }

// SetDoneSeq 设置已处理的包序号，比如已经成功合成了，或者主动丢弃到该位置结束丢弃了
func (l *RtpPacketList) SetDoneSeq(seq uint16) { _ = "STUB: not implemented"; return }

func (l *RtpPacketList) Reset() { _ = "STUB: not implemented"; return }

func (l *RtpPacketList) DebugString() string { _ = "STUB: not implemented"; return "" }
