// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"net/http"

	"github.com/q191201771/lal/pkg/base"
)

// TODO(chef): refactor 配置参数供外部传入
// TODO(chef): refactor maxTaskLen修改为能表示是阻塞任务的意思
var (
	maxTaskLen       = 1024
	notifyTimeoutSec = 3
)

type PostTask struct {
	url  string
	info interface{}
}

type HttpNotify struct {
	cfg HttpNotifyConfig

	serverId string

	taskQueue chan PostTask
	client    *http.Client
}

func NewHttpNotify(cfg HttpNotifyConfig, serverId string) *HttpNotify {
	_ = "STUB: not implemented"
	return nil
}

// TODO(chef): Dispose

// ---------------------------------------------------------------------------------------------------------------------

func (h *HttpNotify) NotifyServerStart(info base.LalInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) NotifyUpdate(info base.UpdateInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) NotifyPubStart(info base.PubStartInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) NotifyPubStop(info base.PubStopInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) NotifySubStart(info base.SubStartInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) NotifySubStop(info base.SubStopInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) NotifyPullStart(info base.PullStartInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) NotifyPullStop(info base.PullStopInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) NotifyRtmpConnect(info base.RtmpConnectInfo) {
	_ = "STUB: not implemented"
	return
}

func (h *HttpNotify) NotifyOnHlsMakeTs(info base.HlsMakeTsInfo) { _ = "STUB: not implemented"; return }

// ----- implement INotifyHandler interface ----------------------------------------------------------------------------

func (h *HttpNotify) OnServerStart(info base.LalInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) OnUpdate(info base.UpdateInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) OnPubStart(info base.PubStartInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) OnPubStop(info base.PubStopInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) OnSubStart(info base.SubStartInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) OnSubStop(info base.SubStopInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) OnRelayPullStart(info base.PullStartInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) OnRelayPullStop(info base.PullStopInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) OnRtmpConnect(info base.RtmpConnectInfo) { _ = "STUB: not implemented"; return }

func (h *HttpNotify) OnHlsMakeTs(info base.HlsMakeTsInfo) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

func (h *HttpNotify) RunLoop() { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

func (h *HttpNotify) asyncPost(url string, info interface{}) { _ = "STUB: not implemented"; return }

// noop

func (h *HttpNotify) post(url string, info interface{}) { _ = "STUB: not implemented"; return }
