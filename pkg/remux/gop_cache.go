// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package remux

import (
	"github.com/q191201771/lal/pkg/base"
)

// GopCache
//
// 提供两个功能:
//  1. 缓存Metadata, VideoSeqHeader, AacSeqHeader
//  2. 缓存音视频GOP数据
//
// 以下，只讨论GopCache的第2点功能
//
// 音频和视频都会缓存。
//
// GopCache也可能不缓存GOP数据，见NewGopCache函数的gopNum参数说明。
//
// 以下，我们只讨论gopNum > 0(也即gopSize > 1)的情况。
//
// GopCache为空时，只有输入了关键帧，才能开启GOP缓存，非关键帧以及音频数据不会被缓存。
// 因此，单音频的流是ok的，相当于不缓存任何数据。
//
// GopCache不为空时，输入关键帧触发生成新的GOP元素，其他情况则往最后一个GOP元素一直追加。
//
// first用于读取第一个GOP（可能不完整），last的前一个用于写入当前GOP。
//
// 最近不完整的GOP也会被缓存，见NewGopCache函数的gopNum参数说明。
//
// -----
// gopNum  = 1
// gopSize = 2
//
//	first     |   first       |       first   | 在后面两个状态间转换，就不画了
//	  |       |     |         |        |      |
//	  0   1   |     0   1	  |    0   1      |
//	  *   *   |     *   *	  |    *   *      |
//	  |       |         |	  |    |          |
//	last      |        last   |   last        |
//	          |               |               |
//	(empty)   |   (full)      |   (full)      |
//
// GetGopCount: 0         |   1           |   1           |
// -----
type GopCache struct {
	t         string
	uniqueKey string

	MetadataEnsureWithSetDataFrame    []byte
	MetadataEnsureWithoutSetDataFrame []byte
	VideoSeqHeader                    []byte
	AacSeqHeader                      []byte

	gopRing              []Gop
	gopRingFirst         int
	gopRingLast          int
	gopSize              int
	singleGopMaxFrameNum int
}

// NewGopCache
//
// @param gopNum: gop缓存大小。
//   - 如果为0，则不缓存音频数据，也即GOP缓存功能不生效。
//   - 如果>0，则缓存[0, gopNum]个GOP，最多缓存 gopNum 个GOP。注意，最后一个GOP可能是不完整的。
func NewGopCache(t string, uniqueKey string, gopNum int, singleGopMaxFrameNum int) *GopCache {
	_ = "STUB: not implemented"
	return nil
}

type LazyGet func() []byte

func (gc *GopCache) SetMetadata(w []byte, wo []byte) {
	_ = "STUB: not implemented"
	// TODO(chef): [refactor] 将metadata等缓存逻辑从GopCache中移除 202207
	return
}

// Feed
//
// @param lg: 内部可能持有lg返回的内存块
func (gc *GopCache) Feed(msg base.RtmpMsg, b []byte) bool {
	_ = "STUB: not implemented"
	// TODO(chef): [refactor] 重构lg两个参数这种方式 202207
	return false
}

// noop

// GetGopCount 获取GOP数量，注意，最后一个可能是不完整的
func (gc *GopCache) GetGopCount() int { _ = "STUB: not implemented"; return 0 }

func (gc *GopCache) GetGopDataAt(pos int) [][]byte { _ = "STUB: not implemented"; return nil }

func (gc *GopCache) Clear() { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

// feedLastGop
//
// 往最后一个GOP元素追加一个msg
// 注意，如果GopCache为空，则不缓存msg
func (gc *GopCache) feedLastGop(msg base.RtmpMsg, b []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// feedNewGop
//
// 生成一个最新的GOP元素，并往里追加一个msg
func (gc *GopCache) feedNewGop(msg base.RtmpMsg, b []byte) { _ = "STUB: not implemented"; return }

func (gc *GopCache) isGopRingFull() bool { _ = "STUB: not implemented"; return false }

func (gc *GopCache) isGopRingEmpty() bool { _ = "STUB: not implemented"; return false }

// ---------------------------------------------------------------------------------------------------------------------

type Gop struct {
	data [][]byte
}

// Feed
//
// @param b: 内部持有`b`内存块
func (g *Gop) Feed(msg base.RtmpMsg, b []byte) { _ = "STUB: not implemented"; return }

func (g *Gop) Clear() { _ = "STUB: not implemented"; return }

func (g *Gop) len() int { _ = "STUB: not implemented"; return 0 }
