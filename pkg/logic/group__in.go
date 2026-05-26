// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"github.com/q191201771/lal/pkg/gb28181"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/rtmp"
	"github.com/q191201771/lal/pkg/rtsp"
)

func (group *Group) AddCustomizePubSession(streamName string) (ICustomizePubSessionContext, error) {
	_ = "STUB: not implemented"
	return *new(ICustomizePubSessionContext), nil
}

func (group *Group) AddRtmpPubSession(session *rtmp.ServerSession) error {
	_ = "STUB: not implemented"
	return nil
}

// AddRtspPubSession TODO chef: rtsp package中，增加回调返回值判断，如果是false，将连接关掉
func (group *Group) AddRtspPubSession(session *rtsp.PubSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (group *Group) StartRtpPub(req base.ApiCtrlStartRtpPubReq) (ret base.ApiCtrlStartRtpPubResp) {
	_ = "STUB: not implemented"
	return *new(base.ApiCtrlStartRtpPubResp)
}

// TODO(chef): [fix] 处理已经有输入session的情况 202207

func (group *Group) AddRtmpPullSession(session *rtmp.PullSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (group *Group) AddRtspPullSession(session *rtsp.PullSession) error {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

func (group *Group) DelPsPubSession(session *gb28181.PubSession) { _ = "STUB: not implemented"; return }

func (group *Group) DelCustomizePubSession(sessionCtx ICustomizePubSessionContext) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) DelRtmpPubSession(session *rtmp.ServerSession) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) DelRtspPubSession(session *rtsp.PubSession) { _ = "STUB: not implemented"; return }

func (group *Group) DelRtmpPullSession(session *rtmp.PullSession) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) DelRtspPullSession(session *rtsp.PullSession) {
	_ = "STUB: not implemented"
	return
}

// ---------------------------------------------------------------------------------------------------------------------

func (group *Group) delPsPubSession(session *gb28181.PubSession) { _ = "STUB: not implemented"; return }

func (group *Group) delCustomizePubSession(sessionCtx ICustomizePubSessionContext) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) delRtmpPubSession(session *rtmp.ServerSession) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) delRtspPubSession(session *rtsp.PubSession) { _ = "STUB: not implemented"; return }

func (group *Group) delPullSession(session base.IObject) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

// addIn 有pub或pull的输入型session加入时，需要调用该函数
func (group *Group) addIn() { _ = "STUB: not implemented"; return }

// delIn 有pub或pull的输入型session离开时，需要调用该函数
func (group *Group) delIn() {
	_ = "STUB: not implemented"
	// 注意，remuxer放前面，使得有机会将内部缓存的数据吐出来
	return
}
