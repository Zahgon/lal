// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package base

import (
	"github.com/q191201771/naza/pkg/connection"
)

type IStatable interface {
	GetStat() connection.Stat // TODO(chef): [refactor] 考虑为 nazanet.UdpConnection 实现这个接口
}

// BasicSessionStat
//
// 包含两部分功能：
// 1. 维护 StatSession 的一些静态信息
// 2. 计算带宽
//
// 计算带宽有两种方式，一种是通过外部的 connection.Connection 获取最新状态，一种是内部自己管理状态
type BasicSessionStat struct {
	stat StatSession

	prevConnStat connection.Stat
	staleStat    *connection.Stat

	currConnStat connection.StatAtomic
}

// ---------------------------------------------------------------------------------------------------------------------

// NewBasicSessionStat
//
// @param remoteAddr: 如果当前未知，填入""空字符串
func NewBasicSessionStat(sessionType SessionType, remoteAddr string) BasicSessionStat {
	_ = "STUB: not implemented"
	return *new(BasicSessionStat)
}

// TODO(chef): [fix] 为customize pub添加 202205

func (s *BasicSessionStat) SetBaseType(baseType string) { _ = "STUB: not implemented"; return }

func (s *BasicSessionStat) SetRemoteAddr(addr string) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

func (s *BasicSessionStat) AddReadBytes(n int) { _ = "STUB: not implemented"; return }

func (s *BasicSessionStat) AddWriteBytes(n int) { _ = "STUB: not implemented"; return }

func (s *BasicSessionStat) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

func (s *BasicSessionStat) UpdateStatWitchConn(conn IStatable, intervalSec uint32) {
	_ = "STUB: not implemented"
	return
}

func (s *BasicSessionStat) GetStat() StatSession {
	_ = "STUB: not implemented"
	return *new(StatSession)
}

func (s *BasicSessionStat) GetStatWithConn(conn IStatable) StatSession {
	_ = "STUB: not implemented"
	return *new(StatSession)
}

func (s *BasicSessionStat) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (s *BasicSessionStat) IsAliveWitchConn(conn IStatable) (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *BasicSessionStat) BaseType() string { _ = "STUB: not implemented"; return "" }

func (s *BasicSessionStat) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------

// updateStat 根据两次调用间隔计算bitrate
func (s *BasicSessionStat) updateStat(readBytesSum, wroteBytesSum uint64, typ string, intervalSec uint32) {
	_ = "STUB: not implemented"
	return
}

// isAlive 根据两次调用间隔计算是否存活
func (s *BasicSessionStat) isAlive(readBytesSum, wroteBytesSum uint64) (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}
