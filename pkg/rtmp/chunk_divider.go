// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtmp

// 将message切割成chunk

import (
	"net"

	"github.com/q191201771/lal/pkg/base"
)

type ChunkDivider struct {
	localChunkSize int
}

var defaultChunkDivider = ChunkDivider{
	localChunkSize: LocalChunkSize,
}

// Message2Chunks
//
// @return 返回的内存块由内部申请，不依赖参数<message>内存块
func Message2Chunks(message []byte, header *base.RtmpHeader) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Message2ChunksV
//
// @param message: 待打包的message支持放在多个字节切片中
func Message2ChunksV(message net.Buffers, header *base.RtmpHeader) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Message2Chunks
//
// TODO chef: [opt] 新的 message 的第一个 chunk 始终使用 fmt0 格式，没有参考前一个 message
func (d *ChunkDivider) Message2Chunks(message []byte, header *base.RtmpHeader) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (d *ChunkDivider) Message2ChunksV(message net.Buffers, header *base.RtmpHeader) []byte {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// @param 返回头的大小
func calcHeader(header *base.RtmpHeader, prevHeader *base.RtmpHeader, out []byte) int {
	_ = "STUB: not implemented"

	// 计算fmt和timestamp
	return 0
}

// 将数据打包成rtmp chunk发送给vlc，时间戳超过3字节最大范围时，
// vlc认为fmt0和fmt3两种格式，都需要携带扩展时间戳字段，并且该时间戳字段必须使用绝对时间戳。

// 设置fmt

// 设置csid

// value 0

// 设置timestamp msgLen msgTypeId msgStreamId

// 设置扩展时间戳

func message2Chunks(message []byte, header *base.RtmpHeader, prevHeader *base.RtmpHeader, chunkSize int) []byte {
	_ = "STUB: not implemented"
	//if header.Csid < minCsid || header.Csid > maxCsid {
	//	return nil, ErrRtmp
	//}
	return nil
}

// 注意，这里我们要尽量缩小预分配内存的大小

// 计算chunk数量，最后一个chunk的大小

// NOTICE 和srs交互时，发现srs要求message中的非第一个chunk不能使用fmt0
// 将message切割成chunk放入chunk body中

// copyBufferFromBuffers
//
// TODO(chef): [refactor] move to naza 202206
// TODO(chef): [perf] impl me
func copyBufferFromBuffers(out []byte, bs net.Buffers, pos int, length int) {
	_ = "STUB: not implemented"
	return
}

func message2ChunksV(message net.Buffers, header *base.RtmpHeader, prevHeader *base.RtmpHeader, chunkSize int) []byte {
	_ = "STUB: not implemented"
	return nil
}

// 计算chunk数量，最后一个chunk的大小

// NOTICE 和srs交互时，发现srs要求message中的非第一个chunk不能使用fmt0
// 将message切割成chunk放入chunk body中

//copy(out[index:], message[i*chunkSize:i*chunkSize+chunkSize])

//copy(out[index:], message[i*chunkSize:i*chunkSize+lastChunkSize])
