// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"github.com/q191201771/lal/pkg/hls"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/lal/pkg/httpts"
	"github.com/q191201771/lal/pkg/rtmp"
	"github.com/q191201771/lal/pkg/rtsp"
)

func (group *Group) AddRtmpSubSession(session *rtmp.ServerSession) {
	_ = "STUB: not implemented"
	return
}

// 加入时，如果上行还没有推过视频（比如还没推流，或者是单音频流），就不需要等待关键帧了
// 也即我们假定上行肯定是以关键帧为开始进行视频发送，假设不是，那么我们按上行的流正常发，而不过滤掉关键帧前面的不包含关键帧的非完整GOP
// TODO(chef):
//   1. 需要仔细考虑单音频无视频的流的情况
//   2. 这里不修改标志，让这个session继续等关键帧也可以

func (group *Group) AddHttpflvSubSession(session *httpflv.SubSession) {
	_ = "STUB: not implemented"
	return
}

// 加入时，如果上行还没有推流过，就不需要等待关键帧了

// AddHttptsSubSession ...
func (group *Group) AddHttptsSubSession(session *httpts.SubSession) {
	_ = "STUB: not implemented"
	return
}

// AddHlsSubSession ...
func (group *Group) AddHlsSubSession(session *hls.SubSession) { _ = "STUB: not implemented"; return }

func (group *Group) HandleNewRtspSubSessionDescribe(session *rtsp.SubSession) (ok bool, sdp []byte) {
	_ = "STUB: not implemented"
	return false, nil
}

func (group *Group) HandleNewRtspSubSessionPlay(session *rtsp.SubSession) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) DelRtmpSubSession(session *rtmp.ServerSession) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) DelHttpflvSubSession(session *httpflv.SubSession) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) DelHttptsSubSession(session *httpts.SubSession) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) DelRtspSubSession(session *rtsp.SubSession) { _ = "STUB: not implemented"; return }

func (group *Group) DelHlsSubSession(session *hls.SubSession) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

func (group *Group) delRtmpSubSession(session *rtmp.ServerSession) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) delHttpflvSubSession(session *httpflv.SubSession) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) delHttptsSubSession(session *httpts.SubSession) {
	_ = "STUB: not implemented"
	return
}

func (group *Group) delRtspSubSession(session *rtsp.SubSession) { _ = "STUB: not implemented"; return }

func (group *Group) delHlsSubSession(session *hls.SubSession) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

func (group *Group) addSub() { _ = "STUB: not implemented"; return }
