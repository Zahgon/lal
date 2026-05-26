// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package remux

// TODO(chef) 将gop_cache.go和gop_cache_mpegts.go的待完成项统一记录在这里
// - GopCache 和 GopCacheMpegts 尽量统一
// - 是否有必要单独存储帧，也即一个Gop的多个帧是一块内存，还是多块内存，从性能，功能，可读考虑
// - GopCache中非gop功能（包括meta和header的缓存）考虑移动到其他地方

type GopCacheMpegts struct {
	uniqueKey string
	gopNum    int

	gopRing              []GopMpegts
	gopRingFirst         int
	gopRingLast          int
	gopSize              int
	singleGopMaxFrameNum int
}

func NewGopCacheMpegts(uniqueKey string, gopNum int, singleGopMaxFrameNum int) *GopCacheMpegts {
	_ = "STUB: not implemented"
	return nil
}

// Feed
//
// @param b: 内部持有该内存块
func (gc *GopCacheMpegts) Feed(b []byte, boundary bool) { _ = "STUB: not implemented"; return }

// GetGopCount 获取GOP数量，注意，最后一个可能是不完整的
func (gc *GopCacheMpegts) GetGopCount() int { _ = "STUB: not implemented"; return 0 }

func (gc *GopCacheMpegts) GetGopDataAt(pos int) [][]byte { _ = "STUB: not implemented"; return nil }

func (gc *GopCacheMpegts) Clear() { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

// feedLastGop
//
// 往最后一个GOP元素追加一个msg
// 注意，如果GopCache为空，则不缓存msg
func (gc *GopCacheMpegts) feedLastGop(b []byte) { _ = "STUB: not implemented"; return }

// feedNewGop
//
// 生成一个最新的GOP元素，并往里追加一个msg
func (gc *GopCacheMpegts) feedNewGop(b []byte) { _ = "STUB: not implemented"; return }

func (gc *GopCacheMpegts) isGopRingFull() bool { _ = "STUB: not implemented"; return false }

func (gc *GopCacheMpegts) isGopRingEmpty() bool { _ = "STUB: not implemented"; return false }

// ---------------------------------------------------------------------------------------------------------------------

// GopMpegts
//
// 单个Gop，包含多帧数据
type GopMpegts struct {
	data [][]byte
}

// Feed
//
// @param b: 内部持有`b`内存块
func (g *GopMpegts) Feed(b []byte) { _ = "STUB: not implemented"; return }

func (g *GopMpegts) Clear() { _ = "STUB: not implemented"; return }

func (g *GopMpegts) len() int { _ = "STUB: not implemented"; return 0 }
