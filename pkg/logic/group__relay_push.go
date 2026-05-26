// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"github.com/q191201771/lal/pkg/rtmp"
)

// TODO(chef): [refactor] 参照relay pull，整体重构一次relay push 202205

func (group *Group) AddRtmpPushSession(url string, session *rtmp.PushSession) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) DelRtmpPushSession(url string, session *rtmp.PushSession) {
	_ = "STUB: not implemented"
	return
}

// ---------------------------------------------------------------------------------------------------------------------

type pushProxy struct {
	isPushing   bool
	pushSession *rtmp.PushSession
}

func (group *Group) initRelayPushByConfig() { _ = "STUB: not implemented"; return }

// startPushIfNeeded 必要时进行replay push转推
func (group *Group) startPushIfNeeded() {
	_ = "STUB: not implemented"
	// push转推功能没开
	return
}

// 没有pub发布者
// TODO(chef): [refactor] 判断所有pub是否存在的方式 202208

// relay push时携带rtmp pub的参数
// TODO chef: 这个逻辑放这里不太好看

// 正在转推中

func (group *Group) stopPushIfNeeded() { _ = "STUB: not implemented"; return }
