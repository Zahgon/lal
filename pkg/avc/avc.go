// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package avc

import (
	"io"
)

// Annexb:
//   keywords: MPEG-2 transport stream, ElementaryStream(ES),
//   nalu with start code.
//   e.g. ts
//
// AVCC:
//   keywords: AVC1, MPEG-4, extradata, sequence header, AVCDecoderConfigurationRecord
//   nalu with length prefix.
//   e.g. rtmp, flv

var (
	NaluStartCode3 = []byte{0x0, 0x0, 0x1}
	NaluStartCode4 = []byte{0x0, 0x0, 0x0, 0x1}

	// AudNalu aud nalu
	AudNalu = []byte{0x00, 0x00, 0x00, 0x01, 0x09, 0xf0}
)

// NaluTypeMapping
//
// H.264-AVC-ISO_IEC_14496-15.pdf
// Table 1 - NAL unit types in elementary streams
var NaluTypeMapping = map[uint8]string{
	1:  "SLICE",
	5:  "IDR",
	6:  "SEI",
	7:  "SPS",
	8:  "PPS",
	9:  "AUD",
	12: "FD",
}

var SliceTypeMapping = map[uint8]string{
	0: "P",
	1: "B",
	2: "I",
	3: "SP",
	4: "SI",
	5: "P",
	6: "B",
	7: "I",
	8: "SP",
	9: "SI",
}

const (
	NaluTypeSlice    uint8 = 1
	NaluTypeIdrSlice uint8 = 5
	NaluTypeSei      uint8 = 6
	NaluTypeSps      uint8 = 7
	NaluTypePps      uint8 = 8
	NaluTypeAud      uint8 = 9  // Access Unit Delimiter
	NaluTypeFd       uint8 = 12 // Filler Data
)

const (
	SliceTypeP  uint8 = 0
	SliceTypeB  uint8 = 1
	SliceTypeI  uint8 = 2
	SliceTypeSp uint8 = 3
	SliceTypeSi uint8 = 4
)

type Context struct {
	Profile uint8
	Level   uint8
	Width   uint32
	Height  uint32

	Sps Sps
}

// DecoderConfigurationRecord
//
// H.264-AVC-ISO_IEC_14496-15.pdf
// 5.2.4 Decoder configuration information
type DecoderConfigurationRecord struct {
	ConfigurationVersion uint8
	AvcProfileIndication uint8
	ProfileCompatibility uint8
	AvcLevelIndication   uint8
	LengthSizeMinusOne   uint8
	NumOfSps             uint8
	SpsLength            uint16
	NumOfPps             uint8
	PpsLength            uint16
}

// Sps
//
// ISO-14496-10.pdf
// 7.3.2.1 Sequence parameter set RBSP syntax
// 7.4.2.1 Sequence parameter set RBSP semantics
type Sps struct {
	ProfileIdc         uint8
	ConstraintSet0Flag uint8
	ConstraintSet1Flag uint8
	ConstraintSet2Flag uint8
	LevelIdc           uint8
	SpsId              uint32

	ChromaFormatIdc            uint32
	ResidualColorTransformFlag uint8
	BitDepthLuma               uint32
	BitDepthChroma             uint32
	TransFormBypass            uint8

	Log2MaxFrameNumMinus4 uint32
	PicOrderCntType       uint32
	Log2MaxPicOrderCntLsb uint32

	NumRefFrames                   uint32 // num_ref_frames
	GapsInFrameNumValueAllowedFlag uint8  // gaps_in_frame_num_value_allowed_flag
	PicWidthInMbsMinusOne          uint32 // pic_width_in_mbs_minus1
	PicHeightInMapUnitsMinusOne    uint32 // pic_height_in_map_units_minus1

	FrameMbsOnlyFlag         uint8 // frame_mbs_only_flag
	MbAdaptiveFrameFieldFlag uint8 // mb_adaptive_frame_field_flag

	Direct8X8InferenceFlag uint8 // direct_8x8_inference_flag

	FrameCroppingFlag     uint8  // frame_cropping_flag
	FrameCropLeftOffset   uint32 // frame_crop_left_offset
	FrameCropRightOffset  uint32 // frame_crop_right_offset
	FrameCropTopOffset    uint32 // frame_crop_top_offset
	FrameCropBottomOffset uint32 // frame_crop_bottom_offset

	SarNum int
	SarDen int
}

func ParseNaluType(v uint8) uint8 { _ = "STUB: not implemented"; return 0 }

func ParseSliceType(nalu []byte) (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

// skip first_mb_in_slice

// range: [0, 9]

func ParseNaluTypeReadable(v uint8) string { _ = "STUB: not implemented"; return "" }

func ParseSliceTypeReadable(nalu []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

// 这些类型不属于视频帧数据类型，没有slice type

// SpsPpsSeqHeader2Annexb
//
// AVCC Seq Header转换为Annexb格式。
//
// @param payload:
//
//	rtmp message的payload部分或者flv tag的payload部分。
//	注意，包含了头部2字节类型以及3字节的cts。
//
// @return 返回的内存块为内部独立新申请。
func SpsPpsSeqHeader2Annexb(payload []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// TODO(chef): [refactor] 这里没有使用 ParseSpsPpsFromSeqHeaderWithoutMalloc
	// 因为遇到了sps>1个的情况
	// 需要重构相关的代码
	return nil, nil
}

// ParseSpsPpsFromSeqHeader
//
// 见func ParseSpsPpsFromSeqHeaderWithoutMalloc
//
// @return sps, pps: 内存块为内部独立新申请
func ParseSpsPpsFromSeqHeader(payload []byte) (sps, pps []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// BuildSpsPps2Annexb
//
// 根据sps pps构建payload
func BuildSpsPps2Annexb(sps, pps []byte) []byte { _ = "STUB: not implemented"; return nil }

// ParseSpsPpsFromSeqHeaderWithoutMalloc
//
// 从AVCC格式的Seq Header中得到SPS和PPS内存块。
//
// @param payload: rtmp message的payload部分或者flv tag的payload部分。
//
//	注意，包含了头部2字节类型以及3字节的cts。
//
// @return sps, pps: 复用传入参数`payload`的内存块
func ParseSpsPpsFromSeqHeaderWithoutMalloc(payload []byte) (sps, pps []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// BuildSeqHeaderFromSpsPps
//
// @return 内存块为内部独立新申请
func BuildSeqHeaderFromSpsPps(sps, pps []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// H.264-AVC-ISO_IEC_14496-15.pdf
// 5.2.4 Decoder configuration information
// configurationVersion

// AvcProfileIndication
// profile_compatibility
// AvcLevelIndication
// lengthSizeMinusOne '111111'b | (4-1)
// numOfSequenceParameterSets '111'b | 1

// sequenceParameterSetLength

// numOfPictureParameterSets 1

// sequenceParameterSetLength

// CaptureAvcc2Annexb
//
// AVCC转换为Annexb格式。
//
// @param payload: rtmp message的payload部分或者flv tag的payload部分。
//
//	注意，包含了头部2字节类型以及3字节的cts
func CaptureAvcc2Annexb(w io.Writer, payload []byte) error {
	_ = "STUB: not implemented"
	// sps pps
	return nil
}

// TODO(chef): [refactor] 使用IterateNaluAvcc
// payload中可能存在多个nalu

// IterateNaluStartCode
//
// 遍历直到找到第一个nalu start code的位置。
//
// @param start: 从`nalu`的start位置开始查找。
//
// @return pos: start code的起始位置（包含start code自身）。
//
// @return length:
//
//	start code的长度，可能是3或者4。
//	注意，如果找不到start code，则返回-1, -1。
func IterateNaluStartCode(nalu []byte, start int) (pos, length int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// SplitNaluAnnexb
//
// 遍历Annexb格式，去掉start code，获取nal包，正常情况下可能为1个或多个，异常情况下可能一个也没有
//
// 具体见单元测试
//
// @return nalList: 内存块元素引用输入参数`nals`的内存
func SplitNaluAnnexb(nals []byte) (nalList [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SplitNaluAvcc
//
// 遍历AVCC格式，去掉4字节长度，获取nal包，正常情况下可能返回1个或多个，异常情况下可能一个也没有
//
// 具体见单元测试
func SplitNaluAvcc(nals []byte) (nalList [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IterateNaluAnnexb
//
// @param handler: 回调函数中的`nal`参数引用`nals`中的内存
func IterateNaluAnnexb(nals []byte, handler func(nal []byte)) error {
	_ = "STUB: not implemented"
	return nil
}

func IterateNaluAvcc(nals []byte, handler func(nal []byte)) error {
	_ = "STUB: not implemented"
	return nil
}

// 非最后一个

// length为0的直接过滤掉

// 最后一个

func Avcc2Annexb(nals []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// TODO(chef): 增加原地转换，不申请内存的方式 202206
	return nil, nil
}

func Annexb2Avcc(nals []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// TODO(chef): 增加原地转换，不申请内存的方式。考虑原地内存不够大的情况 202206
	return nil, nil
}

// perf: start code是三字节0 0 1时，转换时每个nal会多需要一个字节，预先申请16个字节，减少后续扩容的可能性

// ---------------------------------------------------------------------------------------------------------------------

// parseSpsPpsListFromSeqHeaderWithoutMalloc
//
// 从AVCC格式的Seq Header中得到SPS和PPS内存块。
//
// @param payload:
//
//	rtmp message的payload部分或者flv tag的payload部分。
//	注意，包含了头部2字节类型以及3字节的cts。
//
// @return spsList, ppsList:
//
//	复用传入参数`payload`的内存块
func parseSpsPpsListFromSeqHeaderWithoutMalloc(payload []byte) (spsList, ppsList [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// skip 10

// pps和sps的逻辑一样，再一套一层循环处理

// TODO(chef): 考虑nazabits中支持网络序操作

// pps和sps的读取逻辑一样
