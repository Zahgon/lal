// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package mpegts

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
	pts        uint64
	dts        uint64
}

func ParsePes(b []byte) (pes Pes, length int) { _ = "STUB: not implemented"; return *new(Pes), 0 }

// 处理得不是特别标准

//pes.pts = (pes.pts - delay) / 90
//pes.dts = (pes.dts - delay) / 90

// read pts or dts
func readPts(b []byte) (fb uint8, pts uint64) { _ = "STUB: not implemented"; return 0, 0 }
