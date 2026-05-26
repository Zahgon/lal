// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/remux"
	"github.com/q191201771/naza/pkg/nazaatomic"
)

type CustomizePubSessionOption struct {
	DebugDumpPacket string
}

type ModCustomizePubSessionOptionFn func(option *CustomizePubSessionOption)

type CustomizePubSessionContext struct {
	uniqueKey string

	streamName string
	remuxer    *remux.AvPacket2RtmpRemuxer
	onRtmpMsg  func(msg base.RtmpMsg)
	option     CustomizePubSessionOption
	dumpFile   *base.DumpFile

	disposeFlag nazaatomic.Bool
}

func NewCustomizePubSessionContext(streamName string) *CustomizePubSessionContext {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *CustomizePubSessionContext) WithOnRtmpMsg(onRtmpMsg func(msg base.RtmpMsg)) *CustomizePubSessionContext {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *CustomizePubSessionContext) WithCustomizePubSessionContextOption(modFn func(option *CustomizePubSessionOption)) *CustomizePubSessionContext {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *CustomizePubSessionContext) UniqueKey() string { _ = "STUB: not implemented"; return "" }

func (ctx *CustomizePubSessionContext) StreamName() string { _ = "STUB: not implemented"; return "" }

func (ctx *CustomizePubSessionContext) Dispose() { _ = "STUB: not implemented"; return }

// -----implement of base.IAvPacketStream ------------------------------------------------------------------------------

func (ctx *CustomizePubSessionContext) WithOption(modOption func(option *base.AvPacketStreamOption)) {
	_ = "STUB: not implemented"
	return
}

func (ctx *CustomizePubSessionContext) FeedAudioSpecificConfig(asc []byte) error {
	_ = "STUB: not implemented"
	return nil
}

//nazalog.Debugf("[%s] FeedAudioSpecificConfig. asc=%s", ctx.uniqueKey, hex.Dump(asc))

func (ctx *CustomizePubSessionContext) FeedAvPacket(packet base.AvPacket) error {
	_ = "STUB: not implemented"
	return nil
}

//nazalog.Debugf("[%s] FeedAvPacket. packet=%s", ctx.uniqueKey, packet.DebugString())

func (ctx *CustomizePubSessionContext) FeedRtmpMsg(msg base.RtmpMsg) error {
	_ = "STUB: not implemented"
	return nil
}
