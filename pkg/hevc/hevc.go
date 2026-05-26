// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package hevc

import (
	"github.com/q191201771/naza/pkg/nazabits"
)

// HVCC
//
// ISO_IEC_23008-2_2013.pdf

// NAL Unit Header
//
// +---------------+---------------+
// |0|1|2|3|4|5|6|7|0|1|2|3|4|5|6|7|
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |F|   Type    |  LayerId  | TID |
// +-------------+-----------------+

var (
	NaluStartCode4 = []byte{0x0, 0x0, 0x0, 0x1}

	// AudNalu aud nalu
	AudNalu = []byte{0x00, 0x00, 0x00, 0x01, 0x46, 0x01, 0x10}
)

var NaluTypeMapping = map[uint8]string{
	NaluTypeSliceTrailN: "TrailN",
	NaluTypeSliceTrailR: "TrailR",
	NaluTypeSliceTsaN:   "TsaN",
	NaluTypeSliceTsaR:   "TsaR",
	NaluTypeSliceStsaN:  "StsaN",
	NaluTypeSliceStsaR:  "StsaR",
	NaluTypeSliceRadlN:  "RadlN",
	NaluTypeSliceRadlR:  "RadlR",
	NaluTypeSliceRaslN:  "RaslN",
	NaluTypeSliceRaslR:  "RaslR",

	NaluTypeSliceBlaWlp:       "BlaWlp",
	NaluTypeSliceBlaWradl:     "BlaWradl",
	NaluTypeSliceBlaNlp:       "BlaNlp",
	NaluTypeSliceIdr:          "IDR",
	NaluTypeSliceIdrNlp:       "IDRNLP",
	NaluTypeSliceCranut:       "CRANUT",
	NaluTypeSliceRsvIrapVcl22: "IrapVcl22",
	NaluTypeSliceRsvIrapVcl23: "IrapVcl23",

	NaluTypeVps:       "VPS",
	NaluTypeSps:       "SPS",
	NaluTypePps:       "PPS",
	NaluTypeAud:       "AUD",
	NaluTypeSei:       "SEI",
	NaluTypeSeiSuffix: "SEISuffix",
}

// ISO_IEC_23008-2_2013.pdf
// Table 7-1 – NAL unit type codes and NAL unit type classes
const (
	NaluTypeSliceTrailN uint8 = 0 // 0x0
	NaluTypeSliceTrailR uint8 = 1 // 0x01
	NaluTypeSliceTsaN   uint8 = 2 // 0x02
	NaluTypeSliceTsaR   uint8 = 3 // 0x03
	NaluTypeSliceStsaN  uint8 = 4 // 0x04
	NaluTypeSliceStsaR  uint8 = 5 // 0x05
	NaluTypeSliceRadlN  uint8 = 6 // 0x06
	NaluTypeSliceRadlR  uint8 = 7 // 0x07
	NaluTypeSliceRaslN  uint8 = 8 // 0x06
	NaluTypeSliceRaslR  uint8 = 9 // 0x09

	NaluTypeSliceBlaWlp       uint8 = 16 // 0x10
	NaluTypeSliceBlaWradl     uint8 = 17 // 0x11
	NaluTypeSliceBlaNlp       uint8 = 18 // 0x12
	NaluTypeSliceIdr          uint8 = 19 // 0x13
	NaluTypeSliceIdrNlp       uint8 = 20 // 0x14
	NaluTypeSliceCranut       uint8 = 21 // 0x15
	NaluTypeSliceRsvIrapVcl22 uint8 = 22 // 0x16
	NaluTypeSliceRsvIrapVcl23 uint8 = 23 // 0x17

	NaluTypeVps       uint8 = 32 // 0x20
	NaluTypeSps       uint8 = 33 // 0x21
	NaluTypePps       uint8 = 34 // 0x22
	NaluTypeAud       uint8 = 35 // 0x23
	NaluTypeSei       uint8 = 39 // 0x27
	NaluTypeSeiSuffix uint8 = 40 // 0x28
)

type Context struct {
	PicWidthInLumaSamples  uint32 // sps
	PicHeightInLumaSamples uint32 // sps

	ConfigurationVersion uint8 // const value: 1

	GeneralProfileSpace              uint8
	GeneralTierFlag                  uint8
	GeneralProfileIdc                uint8
	GeneralProfileCompatibilityFlags uint32 // const value: 0xffffffff
	GeneralConstraintIndicatorFlags  uint64 // const value: 0xffffffffffff
	GeneralLevelIdc                  uint8

	LengthSizeMinusOne uint8 // const value: 3

	NumTemporalLayers uint8
	TemporalIdNested  uint8

	ChromaFormat         uint8
	BitDepthLumaMinus8   uint8
	BitDepthChromaMinus8 uint8
}

func ParseNaluTypeReadable(v uint8) string { _ = "STUB: not implemented"; return "" }

// ParseNaluType
//
// @param v 第一个字节
func ParseNaluType(v uint8) uint8 {
	_ = "STUB: not implemented"
	// 6 bit in middle
	// 0*** ***0
	// or return (nalu[0] >> 1) & 0x3F
	return 0
}

// IsIrapNalu 是否是关键帧
//
// @param typ 帧类型。注意，是经过 ParseNaluType 解析后的帧类型
func IsIrapNalu(typ uint8) bool {
	_ = "STUB: not implemented"
	// [16, 23] irap nal
	// [19, 20] idr nal
	return false
}

// VpsSpsPpsSeqHeader2Annexb
//
// HVCC Seq Header -> Annexb
//
// @return 返回的内存块为内部独立新申请
func VpsSpsPpsSeqHeader2Annexb(payload []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func VpsSpsPpsEnhancedSeqHeader2Annexb(payload []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func BuildVpsSpsPps2Annexb(vps, sps, pps []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseVpsSpsPpsFromSeqHeader
//
// 见func ParseVpsSpsPpsFromSeqHeaderWithoutMalloc
//
// @return vps, sps, pps: 内存块为内部独立新申请
func ParseVpsSpsPpsFromSeqHeader(payload []byte) (vps, sps, pps []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func ParseVpsSpsPpsFromEnhancedSeqHeader(payload []byte) (vps, sps, pps []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// ParseVpsSpsPpsFromSeqHeaderWithoutMalloc
//
// 从HVCC格式的Seq Header中得到VPS，SPS，PPS内存块。
//
// @param payload: rtmp message的payload部分或者flv tag的payload部分。
//
//	注意，包含了头部2字节类型以及3字节的cts。
//
// @return vps, sps, pps: 复用传入参数`payload`的内存块。
func ParseVpsSpsPpsFromSeqHeaderWithoutMalloc(payload []byte) (vps, sps, pps []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

//Log.Debugf("%s", hex.Dump(payload))

//Log.Warnf("parse vps sps pps from seq header failed. try parse from annexb. payload=%s, err=%+v", hex.Dump(payload), err)

// TODO(chef): 函数中vps/sps/pps变量指向的为新申请的内存块，与调用它的函数ParseVpsSpsPpsFromSeqHeaderWithoutMalloc中的WithoutMalloc语义不相符 202405
func parseVpsSpsPpsAnnexbFromRecord(payload []byte) (vps, sps, pps []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func parseVpsSpsPpsFromRecord(payload []byte) (vps, sps, pps []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// 注意，seq header中，是最后6个字节而不是中间6个字节

// BuildSeqHeaderFromVpsSpsPps
//
// @return 内存块为内部独立新申请
func BuildSeqHeaderFromVpsSpsPps(vps, sps, pps []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// unsigned int(8) configurationVersion = 1;

// unsigned int(2) general_profile_space;
// unsigned int(1) general_tier_flag;
// unsigned int(5) general_profile_idc;

// unsigned int(32) general_profile_compatibility_flags

// unsigned int(48) general_constraint_indicator_flags

// unsigned int(8) general_level_idc;

// bit(4) reserved = ‘1111’b;
// unsigned int(12) min_spatial_segmentation_idc;
// bit(6) reserved = ‘111111’b;
// unsigned int(2) parallelismType;
// TODO chef: 这两个字段没有解析

// bit(6) reserved = ‘111111’b;
// unsigned int(2) ChromaFormat;

// bit(5) reserved = ‘11111’b;
// unsigned int(3) BitDepthLumaMinus8;

// bit(5) reserved = ‘11111’b;
// unsigned int(3) BitDepthChromaMinus8;

// bit(16) avgFrameRate;

// bit(2) constantFrameRate;
// bit(3) NumTemporalLayers;
// bit(1) TemporalIdNested;
// unsigned int(2) lengthSizeMinusOne;

// num of vps sps pps

// num of vps

// length

func ParseVps(vps []byte, ctx *Context) error { _ = "STUB: not implemented"; return nil }

// skip
// vps_video_parameter_set_id u(4)
// vps_reserved_three_2bits   u(2)
// vps_max_layers_minus1      u(6)

// skip
// vps_temporal_id_nesting_flag u(1)
// vps_reserved_0xffff_16bits   u(16)

func ParseSps(sps []byte, ctx *Context) error { _ = "STUB: not implemented"; return nil }

// sps_video_parameter_set_id

// sps_temporal_id_nesting_flag

// sps_seq_parameter_set_id

// https://github.com/OpenVisualCloud/SVT-HEVC/issues/319

func parsePtl(br *nazabits.BitReader, ctx *Context, maxSubLayersMinus1 uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func updatePtl(ctx, ptl *Context) { _ = "STUB: not implemented"; return }

func newContext() *Context { _ = "STUB: not implemented"; return nil }

// 4 bytes

func nal2rbsp(nal []byte) []byte {
	_ = "STUB: not implemented"
	// TODO chef:
	// 1. 输出应该可由外部申请
	// 2. 替换性能
	// 3. 该函数应该放入avc中
	return nil
}
