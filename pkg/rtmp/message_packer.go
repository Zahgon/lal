// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtmp

import (
	"io"
)

const (
	peerBandwidthLimitTypeHard    = uint8(0)
	peerBandwidthLimitTypeSoft    = uint8(1)
	peerBandwidthLimitTypeDynamic = uint8(2)
)

// MessagePacker 打包并发送 rtmp 信令
type MessagePacker struct {
	b *Buffer
}

func NewMessagePacker() *MessagePacker { _ = "STUB: not implemented"; return nil }

// 注意，这个函数只会打包一个chunk头，所以调用方应自己保证在`bodyLen`小于chunk size时使用
func writeSingleChunkHeader(out []byte, csid int, bodyLen int, typeid uint8, streamid int) {
	_ = "STUB: not implemented"
	// 目前这个函数只供发送信令时调用，信令的 csid 都是小于等于 63 的，如果传入的 csid 大于 63，直接 panic
	return
}

// 0 0 0 是时间戳

func (packer *MessagePacker) ChunkAndWrite(writer io.Writer, csid int, typeid uint8, streamid int) error {
	_ = "STUB: not implemented"
	return nil
}

// 如果一个chunk就够放（大部分信令都是这种情况），我们直接在buffer前面预留的空间写入chunk header内容，避免造成拷贝

func (packer *MessagePacker) writeProtocolControlMessage(writer io.Writer, typeid uint8, val int) error {
	_ = "STUB: not implemented"
	return nil

	// 4
}

func (packer *MessagePacker) writeChunkSize(writer io.Writer, val int) error {
	_ = "STUB: not implemented"
	return nil
}

func (packer *MessagePacker) writeWinAckSize(writer io.Writer, val int) error {
	_ = "STUB: not implemented"
	return nil
}

func (packer *MessagePacker) writePeerBandwidth(writer io.Writer, val int, limitType uint8) error {
	_ = "STUB: not implemented"
	return nil

	// 5
}

// @param isPush: 推流为true，拉流为false
func (packer *MessagePacker) writeConnect(writer io.Writer, appName, tcUrl string, isPush bool) error {
	_ = "STUB: not implemented"
	return nil
}

// fpad True if proxy is being used.

// @param objectEncoding 设置0或者3，表示是Amf0或AMF3，上层可根据connect信令中的objectEncoding值设置该值
func (packer *MessagePacker) writeConnectResult(writer io.Writer, tid int, objectEncoding int) error {
	_ = "STUB: not implemented"
	return nil
}

func (packer *MessagePacker) writeCreateStream(writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil

	// 25 = 15 + 9 + 1
}

func (packer *MessagePacker) writeCreateStreamResult(writer io.Writer, tid int) error {
	_ = "STUB: not implemented"
	return nil

	// 29
}

func (packer *MessagePacker) writePlay(writer io.Writer, streamName string, streamid int) error {
	_ = "STUB: not implemented"
	return nil
}

func (packer *MessagePacker) writePublish(writer io.Writer, appName string, streamName string, streamid int) error {
	_ = "STUB: not implemented"
	return nil
}

// <spec-rtmp_specification_1.0.pdf>
// 7.2.2.6.  publish
//
// Type of publishing.
// Set to "live": Live data is published without recording it in a file.
//

func (packer *MessagePacker) writeOnStatusPublish(writer io.Writer, streamid int) error {
	_ = "STUB: not implemented"
	return nil

	// 105
}

func (packer *MessagePacker) writeOnStatusPlay(writer io.Writer, streamid int) error {
	_ = "STUB: not implemented"
	return nil

	// 96
}

func (packer *MessagePacker) writeStreamIsRecorded(writer io.Writer, streamid uint32) error {
	_ = "STUB: not implemented"
	return nil

	// 6
}

func (packer *MessagePacker) writeStreamBegin(writer io.Writer, streamid uint32) error {
	_ = "STUB: not implemented"
	return nil

	// 6
}

func (packer *MessagePacker) writePingRequest(writer io.Writer, timestamp uint32) error {
	_ = "STUB: not implemented"
	return nil

	// 6
}

func (packer *MessagePacker) writeAcknowledgement(writer io.Writer, seqNum uint32) error {
	_ = "STUB: not implemented"
	return nil

	// 5
}

func (packer *MessagePacker) writePingResponse(writer io.Writer, timestamp uint32) error {
	_ = "STUB: not implemented"
	return nil

	// 6
}

// ---------------------------------------------------------------------------------------------------------------------

// TODO(chef): 整理所有的buffer

type Buffer struct {
	core     []byte
	readPos  int
	writePos int
}

func NewBuffer(n int) *Buffer { _ = "STUB: not implemented"; return nil }

func (b *Buffer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *Buffer) Len() int { _ = "STUB: not implemented"; return 0 }

func (b *Buffer) Reset() { _ = "STUB: not implemented"; return }

func (b *Buffer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (b *Buffer) WriteByte(c byte) error { _ = "STUB: not implemented"; return nil }

func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *Buffer) ModWritePos(pos int) { _ = "STUB: not implemented"; return }

func (b *Buffer) grow(n int) { _ = "STUB: not implemented"; return }

// TODO(chef): 可以先尝试是否能挪出空闲位置
