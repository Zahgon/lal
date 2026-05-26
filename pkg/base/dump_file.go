// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package base

import (
	"os"
)

// TODO(chef): [refactor] move to naza 202208

//
// lal中的支持情况列表：
//
// | 支持情况 | 协议          | 类型 | 应用           | 开关手段     | 方式           | 测试(dump, parse) |
// | 已支持  | ps            | pub | lalserver      | http-api参数 | hook到logic中 | 00               |
// | 已支持  | rtsp          | pull | lalserver     | http-api参数 | 回调到logic中 | 00                |
// | 已支持  | rtsp          | pull | demo/pullrtsp | 运行参数     |  回调到上层逻辑 | 11               |
// | 已支持  | customize pub | pub  | lalserver     | 参数         | 接口提供选项   | 00               |
// | 未支持  | rtmp          |      |               |             |              |                   |

const (
	DumpTypeDefault                             uint32 = 0
	DumpTypePsRtpData                           uint32 = 1  // 1
	DumpTypeRtspRtpData                         uint32 = 17 // 1+16
	DumpTypeRtspSdpData                         uint32 = 18
	DumpTypeCustomizePubData                    uint32 = 33 // 1+16*2
	DumpTypeCustomizePubAudioSpecificConfigData uint32 = 34
	DumpTypeInnerFileHeaderData                 uint32 = 49 // 1+16*3
)

func (d *DumpFile) WriteAvPacket(packet AvPacket, typ uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

const (
	writeVer uint32 = 3
)

type DumpFile struct {
	file *os.File
}

type DumpFileMessage struct {
	Ver       uint32 // 制造数据时的代码版本
	Typ       uint32
	Len       uint32 // Body 的长度
	Timestamp uint64 // 写入时的时间戳
	Reserve   uint32
	Body      []byte
}

func NewDumpFile() *DumpFile { _ = "STUB: not implemented"; return nil }

func (d *DumpFile) OpenToWrite(filename string) (err error) { _ = "STUB: not implemented"; return nil }

func (d *DumpFile) OpenToRead(filename string) (err error) { _ = "STUB: not implemented"; return nil }

func (d *DumpFile) WriteWithType(b []byte, typ uint32) error { _ = "STUB: not implemented"; return nil }

func (d *DumpFile) ReadOneMessage() (m DumpFileMessage, err error) {
	_ = "STUB: not implemented"
	return *new(DumpFileMessage), nil
}

func (d *DumpFile) Close() error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------

func (m *DumpFileMessage) DebugString() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------

func (d *DumpFile) pack(b []byte, typ uint32) []byte {
	_ = "STUB: not implemented"
	// TODO(chef): [perf] 优化这块内存 202211
	return nil
}

// Ver

// Typ

// Len

// Timestamp
