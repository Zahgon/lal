// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package httpflv

import (
	"github.com/q191201771/naza/pkg/mock"
)

var Clock = mock.NewStdClock()

// FlvFilePumpOption
//
// 读取flv文件，将tag按时间戳间隔缓慢（类似于ffmpeg的-re）返回
type FlvFilePumpOption struct {
	IsRecursive bool // 如果为true，则循环返回文件内容（类似于ffmpeg的-stream_loop -1）
}

var defaultFlvFilePumpOption = FlvFilePumpOption{
	IsRecursive: false,
}

type FlvFilePump struct {
	option FlvFilePumpOption
}

type ModFlvFilePumpOption func(option *FlvFilePumpOption)

func NewFlvFilePump(modOptions ...ModFlvFilePumpOption) *FlvFilePump {
	_ = "STUB: not implemented"
	return nil
}

type OnPumpFlvTag func(tag Tag) bool

// Pump
//
// @param onFlvTag 如果回调中返回false，则停止Pump
func (f *FlvFilePump) Pump(filename string, onFlvTag OnPumpFlvTag) error {
	_ = "STUB: not implemented"
	// 一次性将文件所有内容读入内存，后续不再读取文件
	return nil
}

// PumpWithTags @return error 暂时只做预留，目前只会返回nil
func (f *FlvFilePump) PumpWithTags(tags []Tag, onFlvTag OnPumpFlvTag) error {
	_ = "STUB: not implemented"
	return nil
	// 整体的基础时间戳。每轮最后更新
}

// 每一轮的第一个tag时间戳

// 上一个tag的时间戳

// 第一轮的第一个tag的时间戳
// 第一轮的第一个tag的物理时间

// 循环一次，代表遍历文件一次

// 遍历所有tag数据

// metadata只在第一轮发送一次

// 修改时间戳
// 使得不同轮依然线性增长

// 本轮第一个tag

// 修改时间戳
// 如果时间戳比前一个tag的还小，可能发生了跳跃，我们直接设置为上一包的值+1，然后不sleep直接发送

// 当前时间戳与第一轮的第一个tag的时间戳差值

// 当前物理时间与第一轮的第一个tag的物理时间差值

// 如果还没到物理时间差值，就sleep

// 第一轮的第一个tag，记录下来
