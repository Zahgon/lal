// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package gb28181

import (
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/rtprtcp"

	"github.com/q191201771/naza/pkg/nazabytes"
)

// PsUnpacker 解析ps(Program Stream)流
type PsUnpacker struct {
	list     rtprtcp.RtpPacketList
	buf      *nazabytes.Buffer
	audioBuf []byte
	videoBuf []byte

	audioStreamType  uint8
	videoStreamType  uint8
	audioPayloadType base.AvPacketPt
	videoPayloadType base.AvPacketPt

	preAudioPts int64
	preVideoPts int64
	preAudioDts int64
	preVideoDts int64

	preAudioRtpts int64
	preVideoRtpts int64

	onAvPacket base.OnAvPacketFunc

	waitSpsFlag bool

	feedPacketCount     int
	feedBodyCount       int
	onAvPacketWrapCount int
	onAvPacketCount     int
}

func NewPsUnpacker() *PsUnpacker { _ = "STUB: not implemented"; return nil }

// WithOnAvPacket
//
// @param onAvPacket: 回调函数中 base.AvPacket 字段说明：
// PayloadType AvPacketPt 见 base.AvPacketPt。
// Timestamp   int64      dts，单位毫秒。
// Pts         int64      pts，单位毫秒。
// Payload     []byte
// 对于视频，h264和h265是AnnexB格式。
// 对于音频，AAC是前面携带adts的格式。
func (p *PsUnpacker) WithOnAvPacket(onAvPacket base.OnAvPacketFunc) *PsUnpacker {
	_ = "STUB: not implemented"
	return nil
}

// FeedRtpPacket
//
// 注意，内部会处理丢包、乱序等问题
//
// @param b: rtp包，注意，包含rtp包头部分，内部不持有该内存块
func (p *PsUnpacker) FeedRtpPacket(b []byte) error {
	_ = "STUB: not implemented"
	// TODO(chef): [opt] 当前遇到的场景都是，音频和视频共用一个ssrc，并且音频和视频的seq是打在一起的，是否存在两个ssrc的情况？ 202209
	return nil
}

//defer func() {
//	nazalog.Debugf("<<<<<<<<<< PsUnpacker. list=%s", p.list.DebugString())
//}()

//nazalog.Debugf(">>>>>>>>>> PsUnpacker FeedRtpPacket. h=%+v, len=%d, body=%s",
//	ipkt.Header, len(ipkt.Raw), hex.Dump(nazabytes.Prefix(ipkt.Raw[12:], 8)))

// 处理丢包、乱序、重复

// 过期了直接丢掉

//nazalog.Debugf("PsUnpacker NOTICE stale, drop. %d", ipkt.Header.Seq)

// 插入队列
//nazalog.Debugf("PsUnpacker FeedRtpPacket insert. %d", ipkt.Header.Seq)

// 循环判断头部是否是顺序的

// 如果头一个是顺序的，取出来，喂入解析器

//nazalog.Debugf("PsUnpacker FeedRtpBody. %d", opkt.Header.Seq)

// 不是顺序的，如果还没达到容器阈值，就先缓存在容器中，直接退出了
// 注意，如果队列为空，也会走到这，然后通过!full退出

//nazalog.Debugf("PsUnpacker exit check !full.")

// 如果达到容器阈值了，就丢弃一部分
// 丢弃哪些呢？
// 先丢第一个，因为满了至少要丢一个了。
//
// 再丢弃连续的，直到下一个可解析帧位置
// 因为不连续的话，没法判断和正在丢弃的是否同一帧的，可以给个机会看后续是否能收到

//nazalog.Debugf("PsUnpacker NOTICE drop. %d", prev.Header.Seq)

//nazalog.Debugf("PsUnpacker exit drop !sequential.")

//nazalog.Debugf("PsUnpacker exit drop start.")

// 注意，这里需要设置done seq，确保这个seq在以后的判断中可被使用

//nazalog.Debugf("PsUnpacker NOTICE drop. %d", prev.Header.Seq)

// 注意，缓存的数据也需要清除

// FeedRtpBody 注意，传入的数据应该是连续的，属于完整帧的
func (p *PsUnpacker) FeedRtpBody(rtpBody []byte, rtpts uint32) error {
	_ = "STUB: not implemented"

	//nazalog.Debugf("> FeedRtpBody. len=%d, prev buf=%d", len(rtpBody), p.buf.Len())
	return nil
}

// ISO/IEC iso13818-1
//
// 2.5 Program Stream bitstream requirements
//
// TODO(chef): [fix] 有些没做有效长度判断

//nazalog.Debugf("----------pack header----------")

//nazalog.Debugf("----------system header----------")
// 2.5.3.5 System header
// Table 2-32 - Program Stream system header
//

//nazalog.Debugf("----------program stream map----------")

//nazalog.Debugf("----------audio stream----------")

//nazalog.Debugf("----------video stream----------")

//nazalog.Errorf("----------skip----------. %s", hex.Dump(nazabytes.Prefix(rb[i-4:], 32)))

//nazalog.Debugf("----------hik stream----------consumed=%d", consumed)

// TODO(chef): [opt] 所有code都处理后，不符合格式的code可以考虑重置unpacker，重新处理新喂入的数据 202207

// 消费失败并不一定是数据有问题，可能是数据不完整需要等待下一个rtp包

//nazalog.Debugf("skip. %d", i+consumed)

func (p *PsUnpacker) Dispose() { _ = "STUB: not implemented"; return }

func (p *PsUnpacker) parsePsm(rb []byte, index int) int {
	_ = "STUB: not implemented"
	// 2.5.4 Program Stream map
	// Table 2-35 - Program Stream map
	return 0
}

// skip program_stream_map_length
// skip current_next_indicator
// skip reserved
// skip program_stream_map_version
// skip reverved
// skip marked_bit
// 2 + 1 + 1

// program_stream_info_length

// elementary_stream_map_length

//nazalog.Debugf("l=%d, esml=%d", l, esml)

// elementary_stream_info_length

//nazalog.Debugf("streamType=%d, streamId=%d, esil=%d", streamType, streamId, esil)

// skip

func (p *PsUnpacker) parseAvStream(code int, rtpts uint32, rb []byte, index int) int {
	_ = "STUB: not implemented"

	// 注意，由于length是两字节，所以存在一个帧分成多个pes包的情况
	return 0
}

//nazalog.Debugf("parseAvStream. code=%d, expected=%d, actual=%d", code, length, len(rb)-i)

// pes header data length

//nazalog.Debugf("parseAvStream. code=%d, length=%d, pts=%d, dts=%d", code, length, pts, dts)

// 注意，处理音频的逻辑和处理视频的类似，参考处理视频的注释

//nazalog.Debugf("audio code=%d, length=%d, ptsDtsFlag=%d, phdl=%d, pts=%d, dts=%d,type=%d", code, length, ptsDtsFlag, phdl, pts, dts, p.audioStreamType)

// noop

// TODO(chef): [perf] 复用内存块 202209

// noop

// noop

// 判断出当前pes是否是新的帧，然后将缓存中的帧回调给上层

// 当前pes包没有pts

// 整个流的第一帧，啥也不干

// 整个流没有pts字段，退化成使用rtpts

// 使用rtp的时间戳回调，但是，并不用rtp的时间戳更新pts

// 同一帧，啥也不干

// 当前pes包没有pts，而前一个有
// 这种情况我们认为是同一帧，啥也不干

// 当前pes包有pts

// 当前pes包是新的帧，将缓存中的帧回调给上层

// 两种情况：
// 1. pts != prev && prev == -1 也即第一帧
// 2. pts == prev && prev != -1 也即同一帧（前后两个pes包的时间戳相同）
// 这两种情况，都啥也不干

// 注意，是处理完之前的数据后，再将当前pes包存入缓存中

// parsePackHeader 注意，`rb[index:]`为待解析的内存块
func parsePackHeader(rb []byte, index int) int {
	_ = "STUB: not implemented"
	// 2.5.3.3 Pack layer of Program Stream
	// Table 2-33 - Program Stream pack header
	return 0
}

// TODO(chef): 这里按MPEG-2处理，还需要处理MPEG-1 202206

// skip system clock reference(SCR)
// skip PES program mux rate

// skip stuffing

func parsePackStreamBody(rb []byte, index int) int { _ = "STUB: not implemented"; return 0 }

// iterateNaluByStartCode 通过nal start code分隔缓存数据，将nal回调给上层
func (p *PsUnpacker) iterateNaluByStartCode(code int, pts, dts int64) {
	_ = "STUB: not implemented"
	return
}

// 找到下一个，则取两个start code之间的内容
// 找不到下一个，则取startcode到末尾的内容
// 不管是否找到下一个，都回调

func (p *PsUnpacker) onAvPacketWrap(packet *base.AvPacket) { _ = "STUB: not implemented"; return }

//nazalog.Debugf("PsUnpacker > onAvPacketWrap. packet=%s", packet.DebugString())

//nazalog.Debugf("PsUnpacker onAvPacketWrap. type=%d", typ)
// TODO(chef): [opt] 等待sps等信息再开始回调，这个逻辑不完整简化了 202209

//nazalog.Debugf("PsUnpacker onAvPacketWrap. drop. %d", typ)

//nazalog.Debugf("PsUnpacker onAvPacketWrap. drop. %d", typ)

//nazalog.Debugf("PsUnpacker > onAvPacket. packet=%s", packet.DebugString())

// ---------------------------------------------------------------------------------------------------------------------

// TODO(chef): [refactor] 以下代码拷贝来自package mpegts，重复了

// Pes -----------------------------------------------------------
// <iso13818-1.pdf>
// <2.4.3.6 PES packet> <page 49/174>
// <Table E.1 - PES packet header example> <page 142/174>
// <F.0.2 PES packet> <page 144/174>
// packet_start_code_prefix  [24b] *** always 0x00, 0x00, 0x01
// stream_id                 [8b]  *
// PES_packet_length         [16b] **
// '10'                      [2b]
// PES_scrambling_control    [2b]
// PES_priority              [1b]
// data_alignment_indicator  [1b]
// copyright                 [1b]
// original_or_copy          [1b]  *
// PTS_DTS_flags             [2b]
// ESCR_flag                 [1b]
// ES_rate_flag              [1b]
// DSM_trick_mode_flag       [1b]
// additional_copy_info_flag [1b]
// PES_CRC_flag              [1b]
// PES_extension_flag        [1b]  *
// PES_header_data_length    [8b]  *
// -----------------------------------------------------------
type Pes struct {
	pscp       uint32
	sid        uint8
	ppl        uint16
	pad1       uint8
	ptsDtsFlag uint8
	pad2       uint8
	phdl       uint8
	pts        int64
	dts        int64
}

func ParsePes(b []byte) (pes Pes, length int) { _ = "STUB: not implemented"; return *new(Pes), 0 }

//pes.pscp, _ = br.ReadBits32(24)
//pes.sid, _ = br.ReadBits8(8)
//pes.ppl, _ = br.ReadBits16(16)

// 处理得不是特别标准

//pes.pts = (pes.pts - delay) / 90
//pes.dts = (pes.dts - delay) / 90

// read pts or dts
func readPts(b []byte) (fb uint8, pts int64) { _ = "STUB: not implemented"; return 0, 0 }
