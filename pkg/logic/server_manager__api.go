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
)

// server_manager__api.go
//
// 支持http-api功能的部分
//

func (sm *ServerManager) StatLalInfo() base.LalInfo {
	_ = "STUB: not implemented"
	return *new(base.LalInfo)
}

func (sm *ServerManager) StatAllGroup() (sgs []base.StatGroup) {
	_ = "STUB: not implemented"
	return nil
}

func (sm *ServerManager) StatGroup(streamName string) *base.StatGroup {
	_ = "STUB: not implemented"
	return nil
}

// copy

func (sm *ServerManager) CtrlStartRelayPull(info base.ApiCtrlStartRelayPullReq) (ret base.ApiCtrlStartRelayPullResp) {
	_ = "STUB: not implemented"
	return *new(base.ApiCtrlStartRelayPullResp)
}

// 注意，如果group不存在，我们依然relay pull

// CtrlStopRelayPull
//
// TODO(chef): 整理错误值
func (sm *ServerManager) CtrlStopRelayPull(streamName string) (ret base.ApiCtrlStopRelayPullResp) {
	_ = "STUB: not implemented"
	return *new(base.ApiCtrlStopRelayPullResp)
}

// CtrlKickSession
//
// TODO(chef): refactor 不要返回http结果，返回error吧
func (sm *ServerManager) CtrlKickSession(info base.ApiCtrlKickSessionReq) (ret base.ApiCtrlKickSessionResp) {
	_ = "STUB: not implemented"
	return *new(base.ApiCtrlKickSessionResp)
}

func (sm *ServerManager) CtrlAddIpBlacklist(info base.ApiCtrlAddIpBlacklistReq) (ret base.ApiCtrlAddIpBlacklistResp) {
	_ = "STUB: not implemented"
	return *new(base.ApiCtrlAddIpBlacklistResp)
}

func (sm *ServerManager) CtrlStartRtpPub(info base.ApiCtrlStartRtpPubReq) (ret base.ApiCtrlStartRtpPubResp) {
	_ = "STUB: not implemented"
	return *new(base.ApiCtrlStartRtpPubResp)
}

// 注意，如果group不存在，我们依然relay pull
