// Copyright 2023, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"github.com/q191201771/lal/pkg/base"
)

// server_manager__notify.go
//
// NotifyHandler部分。
//
// 注意，所有Notify回调都异步化在单独的协程中执行。
// 使得外部的NotifyHandler实现不会影响到内部逻辑的执行。比如避免回调在锁内触发，回调中调用API又再次拿锁的情况。
//

func (sm *ServerManager) nhInitNotifyHandler() {
	_ = "STUB: not implemented"
	// TODO(chef): [opt] 这里已经做了异步化处理，http notify那边的异步可以去掉 202304
	return
}

// 如果外部没有传入，则使用默认的http notify handler

func (sm *ServerManager) nhOnServerStart(info base.LalInfo) { _ = "STUB: not implemented"; return }

func (sm *ServerManager) nhOnUpdate(info base.UpdateInfo) { _ = "STUB: not implemented"; return }

func (sm *ServerManager) nhOnPubStart(info base.PubStartInfo) { _ = "STUB: not implemented"; return }

func (sm *ServerManager) nhOnPubStop(info base.PubStopInfo) { _ = "STUB: not implemented"; return }

func (sm *ServerManager) nhOnSubStart(info base.SubStartInfo) { _ = "STUB: not implemented"; return }

func (sm *ServerManager) nhOnSubStop(info base.SubStopInfo) { _ = "STUB: not implemented"; return }

func (sm *ServerManager) nhOnRelayPullStart(info base.PullStartInfo) {
	_ = "STUB: not implemented"
	return
}

func (sm *ServerManager) nhOnRelayPullStop(info base.PullStopInfo) {
	_ = "STUB: not implemented"
	return
}

func (sm *ServerManager) nhOnRtmpConnect(info base.RtmpConnectInfo) {
	_ = "STUB: not implemented"
	return
}

func (sm *ServerManager) nhOnHlsMakeTs(info base.HlsMakeTsInfo) { _ = "STUB: not implemented"; return }
