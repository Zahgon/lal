// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"net"

	"github.com/q191201771/lal/pkg/mpegts"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/rtprtcp"
	"github.com/q191201771/lal/pkg/sdp"
)

// group__streaming.go
//
// 包含group中音视频数据转发、转封装协议的逻辑
//

// ---------------------------------------------------------------------------------------------------------------------

// OnReadRtmpAvMsg
//
// 输入rtmp数据.
// 来自 rtmp.ServerSession(Pub), rtmp.PullSession, CustomizePubSessionContext(remux.AvPacket2RtmpRemuxer), (remux.DummyAudioFilter) 的回调.
func (group *Group) OnReadRtmpAvMsg(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

// OnSdp OnRtpPacket OnAvPacket
//
// 输入rtsp(rtp)和rtp合帧之后的数据.
// 来自 rtsp.PubSession 的回调.
func (group *Group) OnSdp(sdpCtx sdp.LogicContext) { _ = "STUB: not implemented"; return }

// OnRtpPacket ...
func (group *Group) OnRtpPacket(pkt rtprtcp.RtpPacket) { _ = "STUB: not implemented"; return }

// OnAvPacket ...
func (group *Group) OnAvPacket(pkt base.AvPacket) { _ = "STUB: not implemented"; return }

// 注意，由于rtsp pub的tcp命令连接和udp接收数据连接是并行的，
// 可能发生rtsp pub已经回调告知结束，数据依然回调的现象，
// 出于性能考虑，底层不判断，由上层按需判断

// ---------------------------------------------------------------------------------------------------------------------

// OnAvPacketFromPsPubSession
//
// 来自 gb28181.PubSession 的回调.
func (group *Group) OnAvPacketFromPsPubSession(pkt *base.AvPacket) {
	_ = "STUB: not implemented"
	// TODO(chef): [refactor] 统一所有回调，AvPacket和*AvPacket 202208
	return
}

//Log.Debugf("Group::OnAvPacketFromPsPubSession. pkt=%s", pkt.DebugString())

// ---------------------------------------------------------------------------------------------------------------------

// OnPatPmt OnTsPackets
//
// 输入mpegts数据.
// 来自 remux.Rtmp2MpegtsRemuxer 的回调.
func (group *Group) OnPatPmt(b []byte) { _ = "STUB: not implemented"; return }

// OnTsPackets ...
func (group *Group) OnTsPackets(tsPackets []byte, frame *mpegts.Frame, boundary bool) {
	_ = "STUB: not implemented"
	return
}

// ---------------------------------------------------------------------------------------------------------------------

// onRtmpMsgFromRemux
//
// 输入rtmp数据.
// 来自 remux.AvPacket2RtmpRemuxer 的回调.
func (group *Group) onRtmpMsgFromRemux(msg base.RtmpMsg) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

// onSdpFromRemux onRtpPacketFromRemux
//
// 输入rtsp(rtp)数据.
// 来自 remux.Rtmp2RtspRemuxer 的回调.
func (group *Group) onSdpFromRemux(sdpCtx sdp.LogicContext) { _ = "STUB: not implemented"; return }

// onRtpPacketFromRemux ...
func (group *Group) onRtpPacketFromRemux(pkt rtprtcp.RtpPacket) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

// OnFragmentOpen
//
// 来自 hls.Muxer 的回调
func (group *Group) OnFragmentOpen() { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

// broadcastByRtmpMsg
//
// 使用rtmp类型的数据做为输入，广播给各协议的输出
//
// @param msg 调用结束后，内部不持有msg.Payload内存块
func (group *Group) broadcastByRtmpMsg(msg base.RtmpMsg) {
	_ = "STUB: not implemented"
	// Log.Debugf("> broadcastByRtmpMsg. %s", msg.DebugString())
	return
}

// 设置好用于发送的 rtmp 头部信息

// # 数据有效性检查

// TODO(chef): 暂时不打开，因为过滤掉了innertest中rtmp和flv的输出和输入就不完全相同了
//if msg.Header.MsgTypeId == base.RtmpTypeIdAudio {
//	if len(msg.Payload) <= 2 {
//		// 注意，ffmpeg有时会发送这几种空数据，这种情况我们直接返回，不打印日志
//		if bytes.Equal(msg.Payload, []byte{0xaf, 0x0}) {
//			// noop
//			return
//		}
//		Log.Errorf("[%s] invalid rtmp audio message. header=%+v, payload=%s",
//			group.UniqueKey, msg.Header, hex.Dump(msg.Payload))
//		return
//	}
//} else if msg.Header.MsgTypeId == base.RtmpTypeIdVideo {
//	if len(msg.Payload) <= 5 {
//		if bytes.Equal(msg.Payload, []byte{0x27, 0x02, 0x0, 0x0, 0x0}) ||
//			bytes.Equal(msg.Payload, []byte{0x17, 0x02, 0x0, 0x0, 0x0}) {
//			// noop
//			return
//		}
//		Log.Errorf("[%s] invalid rtmp video message. header=%+v, payload=%s",
//			group.UniqueKey, msg.Header, hex.Dump(msg.Payload))
//		return
//	}
//}

// # mpegts remuxer

// # rtsp

// # 广播。遍历所有 rtmp sub session，转发数据
// ## 如果是新的 sub session，发送已缓存的信息

// TODO chef: 头信息和full gop也可以在SubSession刚加入时发送

// GOP缓存中肯定包含了关键帧

// 有新加入的sub session（本次循环的第一个新加入的sub session），把rtmp buf writer中的缓存数据全部广播发送给老的sub session
// 从而确保新加入的sub session不会发送这部分脏的数据
// 注意，此处可能被调用多次，但是只有第一次会实际flush缓存数据

// 有sub session在等待关键帧，并且当前是关键帧
// 把rtmp buf writer中的缓存数据全部广播发送给老的sub session
// 并且修改这个sub session的标志
// 让rtmp buf writer来发送这个关键帧

// for loop iterate rtmpSubSessionSet

// ## 转发本次数据

// TODO chef: rtmp sub, rtmp push, httpflv sub 的发送逻辑都差不多，可以考虑封装一下

// # 广播。遍历所有 httpflv sub session，转发数据

// GOP缓存中肯定包含了关键帧

// 是否在等待关键帧

// # 录制flv文件

// # 缓存关键信息，以及gop

// 注意，因为withSdf实际上用不上，而且我们也没实现，所以全部用without了

// # 记录stat

// ---------------------------------------------------------------------------------------------------------------------

func (group *Group) feedRtpPacket(pkt rtprtcp.RtpPacket) {
	_ = "STUB: not implemented"
	// 如果配置项 OutWaitKeyFrameFlag 为false，则音频和视频都直接发送。（音频和视频都不等待视频关键帧，都不等待任何数据）
	return
}

// 是否是视频GOP起始位置
// 保证遍历sub session时，只在必要时检查0次或1次，减少性能开销

// session的 ShouldWaitVideoKeyFrame 为false，那么可能有两种情况：
// 1. 对输入流做智能检测时，判定为流内没有视频
// 2. 该输出流已经发送过了GOP起始数据
//
// 这两种情况下，音频或视频数据都直接发送，不需要等了

// 注意，不是avc和hevc时，直接发送

// ---------------------------------------------------------------------------------------------------------------------

func (group *Group) feedTsPackets(tsPackets []byte, frame *mpegts.Frame, boundary bool) {
	_ = "STUB: not implemented"
	// 注意，hls的处理放在前面，让hls先判断是否打开新的fragment并flush audio
	return
}

// # 遍历 httpts sub session

// ## 如果是新加入者

// 发送头

// 如果有缓存，发送缓存
// 并且设置标志，后续都实时转发就行了

// 新加入逻辑只用走一次

// ## 转发本次数据

// 需要继续等

// for loop iterate httptsSubSessionSet

// ---------------------------------------------------------------------------------------------------------------------

func (group *Group) write2RtmpSubSessions(b []byte) { _ = "STUB: not implemented"; return }

func (group *Group) writev2RtmpSubSessions(bs net.Buffers) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

func (group *Group) feedWaitRtspSubSessions() { _ = "STUB: not implemented"; return }
