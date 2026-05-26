// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package base

const (
	// RtmpTypeIdAudio spec-rtmp_specification_1.0.pdf
	// 7.1. Types of Messages
	RtmpTypeIdAudio        uint8 = 8
	RtmpTypeIdVideo        uint8 = 9
	RtmpTypeIdMetadata     uint8 = 18 // RtmpTypeIdDataMessageAmf0
	RtmpTypeIdSetChunkSize uint8 = 1
	// RtmpTypeIdAck 和 RtmpTypeIdWinAckSize 的含义：
	//
	// 一端向另一端发送 RtmpTypeIdWinAckSize ，要求对端每收够一定数据（一定数据的阈值包含在 RtmpTypeIdWinAckSize 信令中）后，向本端回复 RtmpTypeIdAck 。
	//
	// 常见的应用场景：数据发送端要求数据接收端定时发送心跳信令给本端。
	RtmpTypeIdAck         uint8 = 3
	RtmpTypeIdUserControl uint8 = 4
	// RtmpTypeIdWinAckSize 见 RtmpTypeIdAck
	RtmpTypeIdWinAckSize         uint8 = 5
	RtmpTypeIdBandwidth          uint8 = 6
	RtmpTypeIdCommandMessageAmf3 uint8 = 17
	RtmpTypeIdCommandMessageAmf0 uint8 = 20
	RtmpTypeIdAggregateMessage   uint8 = 22

	// RtmpUserControlStreamBegin RtmpUserControlXxx...
	//
	// user control message type
	//
	RtmpUserControlStreamBegin  uint8 = 0
	RtmpUserControlRecorded     uint8 = 4
	RtmpUserControlPingRequest  uint8 = 6
	RtmpUserControlPingResponse uint8 = 7

	// RtmpFrameTypeKey spec-video_file_format_spec_v10.pdf
	// Video tags
	//   VIDEODATA
	//     FrameType UB[4]
	//     CodecId   UB[4]
	//   AVCVIDEOPACKET
	//     AVCPacketType   UI8
	//     CompositionTime SI24
	//     Data            UI8[n]
	RtmpFrameTypeKey   uint8 = 1
	RtmpFrameTypeInter uint8 = 2

	// RtmpCodecIdAvc
	//
	// Video tags -> VIDEODATA -> CodecID
	//
	// 1: JPEG (currently unused)
	// 2: Sorenson H.263
	// 3: Screen video
	// 4: On2 VP6
	// 5: On2 VP6 with alpha channel
	// 6: Screen video version 2
	// 7: AVC
	//
	RtmpCodecIdAvc  uint8 = 7
	RtmpCodecIdHevc uint8 = 12

	// RtmpAvcPacketTypeSeqHeader RtmpAvcPacketTypeNalu RtmpHevcPacketTypeSeqHeader RtmpHevcPacketTypeNalu
	// 注意，按照标准文档上描述，PacketType还有可能为2：
	// 2: AVC end of sequence (lower level NALU sequence ender is not required or supported)
	//
	// 我自己遇到过在流结尾时，对端发送 27 02 00 00 00的情况（比如我们的使用wontcry.flv的单元测试，最后一个包）
	//
	RtmpAvcPacketTypeSeqHeader  uint8 = 0
	RtmpAvcPacketTypeNalu       uint8 = 1
	RtmpHevcPacketTypeSeqHeader       = RtmpAvcPacketTypeSeqHeader
	RtmpHevcPacketTypeNalu            = RtmpAvcPacketTypeNalu

	// enhanced-rtmp packetType https://github.com/veovera/enhanced-rtmp
	RtmpExPacketTypeSequenceStart uint8 = 0
	RtmpExPacketTypeCodedFrames   uint8 = 1 // CompositionTime不为0时有这个类型
	RtmpExPacketTypeSequenceEnd   uint8 = 2
	RtmpExPacketTypeCodedFramesX  uint8 = 3

	// RtmpExFrameTypeKeyFrame RtmpExFrameTypeXXX...
	//
	// The following FrameType values are defined:
	// 0 = reserved
	// 1 = key frame (a seekable frame)
	// 2 = inter frame (a non-seekable frame)
	// ...
	RtmpExFrameTypeKeyFrame uint8 = 1

	RtmpAvcKeyFrame    = RtmpFrameTypeKey<<4 | RtmpCodecIdAvc
	RtmpHevcKeyFrame   = RtmpFrameTypeKey<<4 | RtmpCodecIdHevc
	RtmpAvcInterFrame  = RtmpFrameTypeInter<<4 | RtmpCodecIdAvc
	RtmpHevcInterFrame = RtmpFrameTypeInter<<4 | RtmpCodecIdHevc

	// RtmpSoundFormatAac spec-video_file_format_spec_v10.pdf
	// Audio tags
	//   AUDIODATA
	//     SoundFormat UB[4]
	//     SoundRate   UB[2]
	//     SoundSize   UB[1]
	//     SoundType   UB[1]
	//   AACAUDIODATA
	//     AACPacketType UI8
	//     Data          UI8[n]
	// 注意，视频的CodecId是后4位，音频是前4位
	RtmpSoundFormatG711A uint8 = 7
	RtmpSoundFormatG711U uint8 = 8
	RtmpSoundFormatAac   uint8 = 10
	RtmpSoundFormatOpus  uint8 = 13

	RtmpAacPacketTypeSeqHeader = 0
	RtmpAacPacketTypeRaw       = 1
)

type RtmpHeader struct {
	Csid         int
	MsgLen       uint32 // 不包含header的大小
	MsgTypeId    uint8  // 8 audio 9 video 18 metadata
	MsgStreamId  int
	TimestampAbs uint32 // dts, 经过计算得到的流上的绝对时间戳，单位毫秒
}

type RtmpMsg struct {
	Header  RtmpHeader
	Payload []byte // Payload不包含Header内容。如果需要将RtmpMsg序列化成RTMP chunk，可调用 rtmp.ChunkDivider 相关的函数
}

func (msg RtmpMsg) IsAvcKeySeqHeader() bool { _ = "STUB: not implemented"; return false }

func (msg RtmpMsg) IsHevcKeySeqHeader() bool { _ = "STUB: not implemented"; return false }

func (msg RtmpMsg) IsEnhanced() bool { _ = "STUB: not implemented"; return false }

func (msg RtmpMsg) IsVideoKeySeqHeader() bool { _ = "STUB: not implemented"; return false }

func (msg RtmpMsg) IsAvcKeyNalu() bool { _ = "STUB: not implemented"; return false }

func (msg RtmpMsg) IsHevcKeyNalu() bool { _ = "STUB: not implemented"; return false }

func (msg RtmpMsg) IsEnchanedHevcNalu() bool { _ = "STUB: not implemented"; return false }

func (msg RtmpMsg) GetEnchanedHevcNaluIndex() int { _ = "STUB: not implemented"; return 0 }

// NALU前面有3个字节CompositionTime

func (msg RtmpMsg) IsVideoKeyNalu() bool { _ = "STUB: not implemented"; return false }

func (msg RtmpMsg) IsAacSeqHeader() bool { _ = "STUB: not implemented"; return false }

func (msg RtmpMsg) VideoCodecId() uint8 { _ = "STUB: not implemented"; return 0 }

func (msg RtmpMsg) AudioCodecId() uint8 { _ = "STUB: not implemented"; return 0 }

func (msg RtmpMsg) Clone() (ret RtmpMsg) { _ = "STUB: not implemented"; return *new(RtmpMsg) }

func (msg RtmpMsg) Dts() uint32 { _ = "STUB: not implemented"; return 0 }

// Pts
//
// 注意，只有视频才能调用该函数获取pts，音频的dts和pts都直接使用 RtmpMsg.Header.TimestampAbs
func (msg RtmpMsg) Pts() uint32 { _ = "STUB: not implemented"; return 0 }

func (msg RtmpMsg) Cts() uint32 { _ = "STUB: not implemented"; return 0 }

func (msg RtmpMsg) DebugString() string { _ = "STUB: not implemented"; return "" }

// e.g. RtmpExPacketTypeSequenceStart
