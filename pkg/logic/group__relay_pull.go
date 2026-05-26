// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/rtsp"

	"github.com/q191201771/lal/pkg/rtmp"
)

// StartPull 外部命令主动触发pull拉流
func (group *Group) StartPull(info base.ApiCtrlStartRelayPullReq) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// StopPull
//
// @return 如果PullSession存在，返回它的unique key
func (group *Group) StopPull() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------

type pullProxy struct {
	staticRelayPullEnable    bool // 是否开启pull TODO(chef): refactor 这两个bool可以考虑合并成一个
	apiEnable                bool
	pullUrl                  string
	pullTimeoutMs            int
	pullRetryNum             int
	autoStopPullAfterNoOutMs int // 没有观看者时，是否自动停止pull
	rtspMode                 int
	debugDumpPacket          string

	startCount   int
	lastHasOutTs int64

	isSessionPulling bool // 是否正在pull，注意，这是一个内部状态，表示的是session的状态，而不是整体任务应该处于的状态
	rtmpSession      *rtmp.PullSession
	rtspSession      *rtsp.PullSession
}

// initRelayPullByConfig 根据配置文件中的静态回源配置来初始化回源设置
func (group *Group) initRelayPullByConfig() { _ = "STUB: not implemented"; return }

func (group *Group) setRtmpPullSession(session *rtmp.PullSession) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) setRtspPullSession(session *rtsp.PullSession) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) resetRelayPullSession() { _ = "STUB: not implemented"; return }

func (group *Group) getStatPull() base.StatPull {
	_ = "STUB: not implemented"
	return *new(base.StatPull)
}

func (group *Group) disposeInactivePullSession() { _ = "STUB: not implemented"; return }

func (group *Group) updatePullSessionStat() { _ = "STUB: not implemented"; return }

func (group *Group) isPullModuleAlive() bool { _ = "STUB: not implemented"; return false }

func (group *Group) tickPullModule() { _ = "STUB: not implemented"; return }

func (group *Group) hasPullSession() bool { _ = "STUB: not implemented"; return false }

func (group *Group) pullSessionUniqueKey() string { _ = "STUB: not implemented"; return "" }

// kickPull
//
// @return 返回true，表示找到对应的session，并关闭
func (group *Group) kickPull(sessionId string) bool { _ = "STUB: not implemented"; return false }

// 判断是否需要pull从远端拉流至本地，如果需要，则触发pull
//
// 当前调用时机：
// 1. 添加新sub session
// 2. 外部命令，比如http api
// 3. 定时器，比如pull的连接断了，通过定时器可以重启触发pull
func (group *Group) pullIfNeeded() (string, error) { _ = "STUB: not implemented"; return "", nil }

// TODO(chef): 处理数据回调，是否应该等待Add成功之后。避免竞态条件中途加入了其他in session

func (group *Group) stopPull() string {
	_ = "STUB: not implemented"
	// 关闭时，清空用于重试的计数
	return ""
}

func (group *Group) shouldStartPull() (bool, error) {
	_ = "STUB: not implemented"
	// 如果本地已经有输入型的流，就不需要pull了
	return false, nil
}

// 已经在pull中，就不需要pull了

// 没人观看自动停的逻辑，是否满足并且需要触发

// 检查重试次数

// 负数永远都重试

// shouldAutoStopPull 是否需要自动停，根据没人观看停的逻辑
func (group *Group) shouldAutoStopPull() bool {
	_ = "STUB: not implemented"
	// 没开启
	return false
}

// 还有观众

// 没有观众，并且设置为立即关闭

// 是否达到时间阈值
