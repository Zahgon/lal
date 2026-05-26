// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package mpegts

// Pmt
//
// ----------------------------------------
// Program Map Table
// <iso13818-1.pdf> <2.4.4.8> <page 64/174>
// table_id                 [8b]  *
// section_syntax_indicator [1b]
// 0                        [1b]
// reserved                 [2b]
// section_length           [12b] **
// program_number           [16b] **
// reserved                 [2b]
// version_number           [5b]
// current_next_indicator   [1b]  *
// section_number           [8b]  *
// last_section_number      [8b]  *
// reserved                 [3b]
// PCR_PID                  [13b] **
// reserved                 [4b]
// program_info_length      [12b] **
// -----loop-----
// stream_type              [8b]  *
// reserved                 [3b]
// elementary_PID           [13b] **
// reserved                 [4b]
// ES_info_length_length    [12b] **
// --------------
// CRC32                    [32b] ****
// ----------------------------------------
type Pmt struct {
	tid             uint8
	ssi             uint8
	sl              uint16
	pn              uint16
	vn              uint8
	cni             uint8
	sn              uint8
	lsn             uint8
	pp              uint16
	pil             uint16
	ProgramElements []PmtProgramElement
	crc32           uint32
}

type PmtProgramElement struct {
	StreamType  uint8
	Pid         uint16
	Length      uint16
	Descriptors []Descriptor
}

func ParsePmt(b []byte) (pmt Pmt) { _ = "STUB: not implemented"; return *new(Pmt) }

func (pmt *Pmt) SearchPid(pid uint16) *PmtProgramElement { _ = "STUB: not implemented"; return nil }

func PackPmt(videoCodecId, audioCodecId int) []byte { _ = "STUB: not implemented"; return nil }
