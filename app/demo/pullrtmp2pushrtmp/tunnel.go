// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"errors"
	"time"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/rtmp"
)

// 注意，当前的策略是，当推流有多个地址时，任意一个失败就会退出整个任务

var ErrClosedByCaller = errors.New("tunnel closed by caller")

type Tunnel struct {
	uk          string
	inUrl       string
	outUrlList  []string
	startTime   time.Time
	startEcChan chan ErrorCode
	pullEcChan  chan ErrorCode
	pushEcChan  chan ErrorCode
	closeChan   chan ErrorCode
	waitChan    chan ErrorCode
	rtmpMsgQ    chan base.RtmpMsg

	pullSession     *rtmp.PullSession
	pushSessionList []*rtmp.PushSession
}

type ErrorCode struct {
	code int // -1表示拉流失败或者结束，>=0表示推流失败或者结束，值对应outUrlList的下标
	err  error
}

// @param inUrl      拉流rtmp url地址
// @param outUrlList 推流rtmp url地址列表
func NewTunnel(inUrl string, outUrlList []string) *Tunnel { _ = "STUB: not implemented"; return nil }

// Start
//
// @return ErrorCode.err:
//   - 为nil时，表示任务启动成功，拉流和推流通道都已成功建立，并开始转推数据。
//   - 不为nil时，表示任务失败，可以通过`code`得到是拉流还是推流失败。
func (t *Tunnel) Start() (ret ErrorCode) { _ = "STUB: not implemented"; return *new(ErrorCode) }

// 最后清理所有session

// 将多个pushSession wait事件聚合在一起

// 主事件监听

// 逐个开启push session

// 只有有一个失败就直接退出

// 加入的都是成功的

//option.ReuseReadMessageBufferFlag = false

// pull失败就直接退出

// `Start`函数调用成功后，可调用`Wait`函数，等待任务结束
func (t *Tunnel) Wait() chan ErrorCode {
	_ = "STUB: not implemented"

	// `Start`函数调用成功后，可调用`Close`函数，主动关闭转推任务
	// `Close`函数允许调用多次
	return nil
}

func (t *Tunnel) Close() { _ = "STUB: not implemented"; return }

func (t *Tunnel) notifyClose() { _ = "STUB: not implemented"; return }

func (t *Tunnel) notifyWait(ec ErrorCode) { _ = "STUB: not implemented"; return }

func (t *Tunnel) notifyStartEc(ec ErrorCode) { _ = "STUB: not implemented"; return }

func (t *Tunnel) notifyPushEc(ec ErrorCode) { _ = "STUB: not implemented"; return }

func (t *Tunnel) notifyPullEc(ec ErrorCode) { _ = "STUB: not implemented"; return }

func (ec *ErrorCode) Stringify() string { _ = "STUB: not implemented"; return "" }
