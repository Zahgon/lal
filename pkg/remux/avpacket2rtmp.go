// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package remux

import (
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/rtmp"
	"github.com/q191201771/lal/pkg/rtprtcp"
	"github.com/q191201771/lal/pkg/sdp"
)

// AvPacket2RtmpRemuxer AvPacket转换为RTMP
//
// 目前AvPacket来自:
//
// - RTSP:       sdp以及rtp的合帧包
// - gb28181 ps: rtp的合帧包
// - customize:  业务方通过接口向lalserver输入的流
// - ffmpeg:     ffmpeg编码后的数据
// - 理论上也支持webrtc，后续接入webrtc时再验证
type AvPacket2RtmpRemuxer struct {
	option    base.AvPacketStreamOption
	onRtmpMsg rtmp.OnReadRtmpAvMsg

	hasEmittedMetadata bool
	audioType          base.AvPacketPt
	videoType          base.AvPacketPt

	vps []byte // 从AvPacket数据中获取
	sps []byte
	pps []byte

	hasAdts2Asc bool
}

func NewAvPacket2RtmpRemuxer() *AvPacket2RtmpRemuxer { _ = "STUB: not implemented"; return nil }

// WithOption
//
// TODO(chef): [refactor] 返回*AvPacket2RtmpRemuxer 202208
func (r *AvPacket2RtmpRemuxer) WithOption(modOption func(option *base.AvPacketStreamOption)) {
	_ = "STUB: not implemented"
	return

	// TODO(chef): [log] 打印所有option 202301
}

func (r *AvPacket2RtmpRemuxer) WithOnRtmpMsg(onRtmpMsg rtmp.OnReadRtmpAvMsg) *AvPacket2RtmpRemuxer {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// OnRtpPacket OnSdp OnAvPacket
//
// 实现RTSP回调数据的接口 rtsp.IBaseInSessionObserver ，使得接入时方便些
func (r *AvPacket2RtmpRemuxer) OnRtpPacket(pkt rtprtcp.RtpPacket) {
	_ = "STUB: not implemented"
	// noop
	return
}

func (r *AvPacket2RtmpRemuxer) OnSdp(sdpCtx sdp.LogicContext) { _ = "STUB: not implemented"; return }

func (r *AvPacket2RtmpRemuxer) OnAvPacket(pkt base.AvPacket) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

// InitWithAvConfig rtsp场景下，有时sps、pps等信息只包含在sdp中，有时包含在rtp包中，
// 这里提供输入sdp的sps、pps等信息的机会，如果没有，可以不调用
//
// 内部不持有输入参数的内存块
func (r *AvPacket2RtmpRemuxer) InitWithAvConfig(asc, vps, sps, pps []byte) {
	_ = "STUB: not implemented"
	return
}

// FeedAvPacket
//
// 输入 base.AvPacket 数据
//
// @param pkt:
//   - 如果是aac，格式是裸数据或带adts头，具体取决于前面的配置。
//   - 如果是h264，格式是avcc或Annexb，具体取决于前面的配置。
//     内部不持有该内存块。
func (r *AvPacket2RtmpRemuxer) FeedAvPacket(pkt base.AvPacket) { _ = "STUB: not implemented"; return }

// 如果有sps，pps，先把它们抽离出来进行缓存

// 注意，由于sps空值时，可能是nil也可能是[0:0]，所以这里不用nil做判断，而用len

// 凑齐了，发送video seq header
//
// TODO(chef): 是否应该判断sps、pps是连续的，比如rtp seq的关系，或者timestamp是相等的

//if !AvPacket2RtmpRemuxerAddSpsPps2KeyFrameFlag {
//	r.clearVideoSeqHeader()
//}

// 重组实际数据

//if AvPacket2RtmpRemuxerAddSpsPps2KeyFrameFlag {
//	// 关键帧 组合sps vps与数据帧
//	nal = append(append(avc.BuildSpsPps2Annexb(r.sps, r.pps)[4:], hevc.NaluStartCode4...), nal...)
//	// 考虑feed时 无sps 与pps数据
//	if len(pkt.Payload) < len(nal) {
//		maxLength = len(nal) + pos
//		payload = make([]byte, maxLength)
//	}
//}

//if !AvPacket2RtmpRemuxerAddSpsPps2KeyFrameFlag {
//	r.clearVideoSeqHeader()
//}

//if AvPacket2RtmpRemuxerAddSpsPps2KeyFrameFlag {
//	// 关键帧 组合vps sps pps与数据帧
//	annexb, err := hevc.BuildVpsSpsPps2Annexb(r.vps, r.sps, r.pps)
//	if err != nil {
//		Log.Errorf("build hevc vps sps pps data failed. err=%+v", err)
//	}
//	nal = append(append(annexb[4:], hevc.NaluStartCode4...), nal...)
//	// 考虑feed时 无sps 与pps数据
//	if len(pkt.Payload) < len(nal) {
//		maxLength = len(nal) + pos
//		payload = make([]byte, maxLength)
//	}
//}

// 有实际数据

// TODO(chef) 处理此处的魔数0xAF

// -7+2

// ffmpeg是固定值

// ffmpeg是固定值

// codecid=13, 44kHz、16bits、Stereo

// ---------------------------------------------------------------------------------------------------------------------

func (r *AvPacket2RtmpRemuxer) emitRtmpAvMsg(isAudio bool, payload []byte, timestamp int64) {
	_ = "STUB: not implemented"
	return
}

// TODO(chef): 此处简化了从sps中获取宽高写入metadata的逻辑

func (r *AvPacket2RtmpRemuxer) setVps(b []byte) { _ = "STUB: not implemented"; return }

func (r *AvPacket2RtmpRemuxer) setSps(b []byte) { _ = "STUB: not implemented"; return }

func (r *AvPacket2RtmpRemuxer) setPps(b []byte) { _ = "STUB: not implemented"; return }

func (r *AvPacket2RtmpRemuxer) clearVideoSeqHeader() { _ = "STUB: not implemented"; return }
