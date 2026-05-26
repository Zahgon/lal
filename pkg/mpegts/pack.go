// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package mpegts

// Frame 帧数据，用于打包成mpegts格式的数据
type Frame struct {
	Pts uint64 // =(毫秒 * 90)
	Cts uint32
	Dts uint64
	Cc  uint8 // continuity_counter of TS Header

	// PID of PES Header
	// 音频 mpegts.PidAudio
	// 视频 mpegts.PidVideo
	Pid uint16

	// stream_id of PES Header
	// 音频 mpegts.StreamIdAudio
	// 视频 mpegts.StreamIdVideo
	Sid uint8

	// 音频 全部为false
	// 视频 关键帧为true，非关键帧为false
	Key bool

	// 音频AAC 格式为2字节ADTS头加raw frame
	// 视频AVC 格式为Annexb
	Raw []byte
}

// Pack annexb格式的流转换为mpegts流
//
// 注意，内部会增加 Frame.Cc 的值.
//
// @return: 内存块为独立申请，调度结束后，内部不再持有
func (frame *Frame) Pack() []byte { _ = "STUB: not implemented"; return nil }

// 预分配一块足够大的内存

// TODO(chef): perf 复用这块buffer

// 当前输入帧的处理位置
// 当前输入帧大小
// 是否为帧的首个packet的标准
// 当前输出packet相对于整个输出内存块的位置

// TODO(chef): CHEFNOTICEME 正常来说，预分配的内存应该是足够用了，我们加个扩容逻辑保证绝对正确性，并且加个日志观察一段时间

// 当前输出packet
// 当前输出packet的写入位置

// 每个packet都需要添加TS Header
// -----TS Header----------------
// sync_byte
// transport_error_indicator    0
// payload_unit_start_indicator
// transport_priority           0
// PID
// transport_scrambling_control 0
// adaptation_field_control
// continuity_counter
// ------------------------------
// sync_byte

// payload_unit_start_indicator

//PID高5位
//PID低8位

// adaptation_field_control 先设置成无Adaptation
// continuity_counter

// 关键帧的首个packet需要添加Adaptation
// -----Adaptation-----------------------
// adaptation_field_length
// discontinuity_indicator              0
// random_access_indicator              1
// elementary_stream_priority_indicator 0
// PCR_flag                             1
// OPCR_flag                            0
// splicing_point_flag                  0
// transport_private_data_flag          0
// adaptation_field_extension_flag      0
// program_clock_reference_base
// reserved
// program_clock_reference_extension
// --------------------------------------
// adaptation_field_control 设置Adaptation
// adaptation_field_length
// random_access_indicator + PCR_flag

// using 6 byte

// 帧的首个packet需要添加PES Header
// -----PES Header------------
// packet_start_code_prefix
// stream_id
// PES_packet_length
// '10'
// PES_scrambling_control    0
// PES_priority              0
// data_alignment_indicator  0
// copyright                 0
// original_or_copy          0
// PTS_DTS_flags
// ESCR_flag                 0
// ES_rate_flag              0
// DSM_trick_mode_flag       0
// additional_copy_info_flag 0
// PES_CRC_flag              0
// PES_extension_flag        0
// PES_header_data_length
// ---------------------------
// packet_start_code_prefix 24-bits
//
//
// stream_id

// 计算PES Header中一些字段的值
// PTS相关

// DTS相关

// PES Header剩余3字节 + PTS/PTS长度 + 整个帧的长度

// PES_packet_length
//
// 除了reserve的'10'，其他字段都是0
// PTS/DTS flag
// PES_header_data_length: PTS+DTS数据长度

// 写入PTS的值

// 写入DTS的值

// 把帧的内容切割放入packet中
// 当前TS packet，可写入大小
// 整个帧剩余待打包大小

// 当前packet写不完这个帧，或者刚好够写完

// 当前packet可以写完这个帧，并且还有空闲空间
// 此时，真实数据挪最后，中间用0xFF填充到Adaptation中
// 注意，此时有两种情况
// 1. 原本有Adaptation
// 2. 原本没有Adaptation

// 当前TS packet的剩余空闲空间

// has Adaptation

// TS Header + Adaptation

// 比如有PES Header

// adaptation_field_length

// no Adaptation

// adaptation_field_length

// TODO chef 这里是参考nginx rtmp module的实现，为什么这个字节写0而不是0xFF

// 真实数据放在packet尾部

func (frame *Frame) DebugString() string { _ = "STUB: not implemented"; return "" }

// ----- private -------------------------------------------------------------------------------------------------------

func packPcr(out []byte, pcr uint64) { _ = "STUB: not implemented"; return }

//pcrLow := pcr % 300
//pcrHigh := pcr / 300
//out[0] = uint8(pcrHigh >> 25)
//out[1] = uint8(pcrHigh >> 17)
//out[2] = uint8(pcrHigh >> 9)
//out[3] = uint8(pcrHigh >> 1)
//out[4] = uint8(pcrHigh<<7) | uint8(pcrLow>>8) | 0x7e
//out[5] = uint8(pcrLow)

// 注意，除PTS外，DTS也使用这个函数打包
func packPts(out []byte, fb uint8, pts uint64) { _ = "STUB: not implemented"; return }
