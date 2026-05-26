// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package avc

import (
	"github.com/q191201771/naza/pkg/nazabits"
)

func ParseSps(payload []byte, ctx *Context) error { _ = "STUB: not implemented"; return nil }

//if err := parseSpsBeta(&br, &sps); err != nil {
//	// 注意，这里不将错误返回给上层，因为可能是Beta自身解析的问题
//}

// 注意，这里不将错误返回给上层，因为可能是Beta自身解析的问题

// TryParsePps 尝试解析PPS所有字段，实验中，请勿直接使用该函数
func TryParsePps(payload []byte) error {
	_ = "STUB: not implemented"
	// ISO-14496-10.pdf
	// 7.3.2.2 Picture parameter set RBSP syntax
	return nil
}

// TODO impl me

// TryParseSeqHeader 尝试解析SeqHeader所有字段，实验中，请勿直接使用该函数。
//
// @param payload:
// rtmp message的payload部分或者flv tag的payload部分。
// 注意，包含了头部2字节类型以及3字节的cts。
func TryParseSeqHeader(payload []byte) error { _ = "STUB: not implemented"; return nil }

// H.264-AVC-ISO_IEC_14496-15.pdf
// 5.2.4 Decoder configuration information

// TODO check error

// reserved = '111111'b

// reserved = '111'b

// reserved = '111'b

// 5 + 5 + 1 + 2

// 13 + 1 + 2

func parseSpsBasic(br *nazabits.BitReader, sps *Sps) error { _ = "STUB: not implemented"; return nil }

//nalType SPS should be 0x67

func parseSpsGamma(br *nazabits.BitReader, sps *Sps) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// chroma_format_idc

// separate_colour_plane_flag

// qpprime_y_zero_transform_bypass_flag

// seq_scaling_matrix_present_flag

// seq_scaling_list_present_flag

// log2_max_frame_num_minus4

// log2_max_pic_order_cnt_lsb_minus4

// delta_pic_order_always_zero
// offset_for_non_ref_pic
// offset_for_top_to_bottom_field

// num_ref_frames_in_pic_order_cnt_cycle

// offset_for_ref_frame

// max_num_ref_frames
// gaps_in_frame_num_value_allowed_flag
// pic_width_in_mbs_minus1
// pic_height_in_map_units_minus1

// mb_adaptive_frame_field_flag

// direct_8x8_inference_flag

// frame_cropping_flag
// frame_crop_left_offset
// frame_crop_right_offset
// frame_crop_top_offset
// frame_crop_bottom_offset

// vui_parameters_present_flag

// aspect_ratio_info_present_flag

// aspect_ratio_idc
