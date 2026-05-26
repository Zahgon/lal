// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package aac

// AudioSpecificConfig(asc)
// keywords: Seq Header,
// e.g.  rtmp, flv
//
// ADTS(Audio Data Transport Stream)
// e.g. es, ts
//
// StreamMuxConfig
//

const (
	AdtsHeaderLength = 7

	AscSamplingFrequencyIndex96000 = 0
	AscSamplingFrequencyIndex88200 = 1
	AscSamplingFrequencyIndex64000 = 2
	AscSamplingFrequencyIndex48000 = 3
	AscSamplingFrequencyIndex44100 = 4
	AscSamplingFrequencyIndex32000 = 5
	AscSamplingFrequencyIndex24000 = 6
	AscSamplingFrequencyIndex22050 = 7
	AscSamplingFrequencyIndex16000 = 8
	AscSamplingFrequencyIndex12000 = 9
	AscSamplingFrequencyIndex11025 = 10
	AscSamplingFrequencyIndex8000  = 11
	AscSamplingFrequencyIndex7350  = 12
)

const (
	minAscLength = 2
)

// AscContext
//
// <ISO_IEC_14496-3.pdf>
// <1.6.2.1 AudioSpecificConfig>, <page 33/110>
// <1.5.1.1 Audio Object type definition>, <page 23/110>
// <1.6.3.3 samplingFrequencyIndex>, <page 35/110>
// <1.6.3.4 channelConfiguration>
// --------------------------------------------------------
// audio object type      [5b] 1=AAC MAIN  2=AAC LC
// samplingFrequencyIndex [4b] 0=96000, 1=88200, 2=64000, 3=48000, 4=44100, 5=32000, 6=24000, 7=22050, 8=16000, 9=12000, 10=11025, 11=11025, 12=7350
// channelConfiguration   [4b] 1=center front speaker  2=left, right front speakers
type AscContext struct {
	AudioObjectType        uint8 // [5b]
	SamplingFrequencyIndex uint8 // [4b]
	ChannelConfiguration   uint8 // [4b]
}

func NewAscContext(asc []byte) (*AscContext, error) { _ = "STUB: not implemented"; return nil, nil }

// Unpack
//
// @param asc: 2字节的AAC Audio Specifc Config。
//
//	注意，如果是rtmp/flv的message/tag，应去除Seq Header头部的2个字节。
//	函数调用结束后，内部不持有该内存块。
func (ascCtx *AscContext) Unpack(asc []byte) error { _ = "STUB: not implemented"; return nil }

// Pack
//
// @return asc: 内存块为独立新申请；函数调用结束后，内部不持有该内存块
func (ascCtx *AscContext) Pack() (asc []byte) { _ = "STUB: not implemented"; return nil }

// PackAdtsHeader
//
// 获取ADTS头，由于ADTS头中的字段依赖包的长度，而每个包的长度可能不同，所以每个包的ADTS头都需要独立生成。
//
// @param frameLength: raw aac frame的大小。
//
//	注意，如果是rtmp/flv的message/tag，应去除Seq Header头部的2个字节。
//
// @return h: 内存块为独立新申请；函数调用结束后，内部不持有该内存块
func (ascCtx *AscContext) PackAdtsHeader(frameLength int) (out []byte) {
	_ = "STUB: not implemented"
	return nil
}

// PackToAdtsHeader
//
// @param out: 函数调用结束后，内部不持有该内存块
func (ascCtx *AscContext) PackToAdtsHeader(out []byte, frameLength int) error {
	_ = "STUB: not implemented"
	return nil
}

// <ISO_IEC_14496-3.pdf>
// <1.A.2.2.1 Fixed Header of ADTS>, <page 75/110>
// <1.A.2.2.2 Variable Header of ADTS>, <page 76/110>
// <1.A.3.2.1 Definitions: Bitstream elements for ADTS>
// ----------------------------------------------------
// Syncword                 [12b] '1111 1111 1111'
// ID                       [1b]  1=MPEG-2 AAC 0=MPEG-4
// Layer                    [2b]
// protection_absent        [1b]  1=no crc check
// Profile_ObjectType       [2b]
// sampling_frequency_index [4b]
// private_bit              [1b]
// channel_configuration    [3b]
// origin/copy              [1b]
// home                     [1b]
// Emphasis???
// ------------------------------------
// copyright_identification_bit   [1b]
// copyright_identification_start [1b]
// aac_frame_length               [13b]
// adts_buffer_fullness           [11b]
// no_raw_data_blocks_in_frame    [2b]

// Syncword 0(8) 1(4)

// ID, Layer, protection_absent 1(4)

// 2(2)
// <ISO_IEC_14496-3.pdf>
// profile_ObjectType: The interpretation of this bitstream element depends on the value of the ID bit. If ID is
// equal to ‘1’ this field holds the same information as the profile field in the ADTS stream defined in ISO/IEC
// 13818-7. If ID is equal to ‘0’ this element denotes the MPEG-4 Audio Object Type according to the table
// defined in subclause 5.1.1.

// 2(4)

// private_bit 2(1)

// 2(1) 3(2)

// origin/copy, home, copyright_identification_bit, copyright_identification_start 3(4)

// 3(2) 4(8) 5(3)

// adts_buffer_fullness 5(5) 6(6)

// no_raw_data_blocks_in_frame 6(2)

func (ascCtx *AscContext) GetSamplingFrequency() (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type AdtsHeaderContext struct {
	AscCtx AscContext

	AdtsLength uint16 // 字段中的值，包含了adts header + adts frame
}

func NewAdtsHeaderContext(adtsHeader []byte) (*AdtsHeaderContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unpack
//
// @param adtsHeader: 函数调用结束后，内部不持有该内存块
func (ctx *AdtsHeaderContext) Unpack(adtsHeader []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// <ISO_IEC_14496-3.pdf>
// profile_ObjectType: The interpretation of this bitstream element depends on the value of the ID bit. If ID is
// equal to ‘1’ this field holds the same information as the profile field in the ADTS stream defined in ISO/IEC
// 13818-7. If ID is equal to ‘0’ this element denotes the MPEG-4 Audio Object Type according to the table
// defined in subclause 5.1.1.

// MakeAscWithAdtsHeader
//
// @param adtsHeader: 函数调用结束后，内部不持有该内存块
//
// @return asc: 内存块为独立新申请；函数调用结束后，内部不持有该内存块
func MakeAscWithAdtsHeader(adtsHeader []byte) (asc []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
