// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"github.com/q191201771/lal/pkg/rtsp"

	"github.com/q191201771/naza/pkg/nazalog"

	"github.com/q191201771/lal/pkg/base"

	"github.com/q191201771/lal/pkg/logic"
)

// 注意，使用这个demo时，请确保这三个文件存在，文件下载地址 https://github.com/q191201771/lalext/tree/master/avfile
const (
	h264filename = "/tmp/test.h264"
	aacfilename  = "/tmp/test.aac"
	flvfilename  = "/tmp/test.flv"
)

// 文档见 <lalserver二次开发 - pub接入自定义流>
// https://pengrl.com/lal/#/customize_pub
//

// MySession 演示业务方实现 logic.ICustomizeHookSessionContext 接口，从而hook所有输入到lalserver中的流以及流中的数据。
type MySession struct {
	uniqueKey  string
	streamName string
}

func (i *MySession) OnMsg(msg base.RtmpMsg) {
	_ = "STUB: not implemented"
	// 业务方可以在这里对流做处理
	return
}

func (i *MySession) OnStop() { _ = "STUB: not implemented"; return }

func main() {
	defer nazalog.Sync()

	rtsp.BaseInSessionTimestampFilterFlag = false

	confFilename := parseFlag()
	lals := logic.NewLalServer(func(option *logic.Option) {
		option.ConfFilename = confFilename
	})

	// 在常规lalserver基础上增加这行，用于演示hook lalserver中的流
	lals.WithOnHookSession(func(uniqueKey string, streamName string) logic.ICustomizeHookSessionContext {
		// 有新的流了，创建业务层的对象，用于hook这个流
		return &MySession{
			uniqueKey:  uniqueKey,
			streamName: streamName,
		}
	})

	// 在常规lalserver基础上增加这两个例子，用于演示向lalserver输入自定义流
	go showHowToCustomizePub(lals)
	go showHowToFlvCustomizePub(lals)

	err := lals.RunLoop()
	nazalog.Infof("lal server loop done. err=%+v", err)
}

func parseFlag() string { _ = "STUB: not implemented"; return "" }

func showHowToFlvCustomizePub(lals logic.ILalServer) { _ = "STUB: not implemented"; return }

func showHowToCustomizePub(lals logic.ILalServer) { _ = "STUB: not implemented"; return }

// 从音频和视频各自的ES流文件中读取出所有数据
// 然后将它们按时间戳排序，合并到一个AvPacket数组中

// 1. 向lalserver中加入自定义的pub session

// 2. 配置session

// 3. 填入aac的audio specific config信息

// 4. 按时间戳间隔匀速发送音频和视频

//nazalog.Debugf("%d: %s, %d, %d", i, packets[i].DebugString(), diffTs, diffReal)

// 5. 所有数据发送关闭后，将pub session从lal server移除

// readAudioPacketsFromFile 从aac es流文件读取所有音频包
func readAudioPacketsFromFile(filename string) (audioContent []byte, audioPackets []base.AvPacket) {
	_ = "STUB: not implemented"
	return nil, nil
}

// (frequence * bytePerSample * channel) / (packetSize * channel)

// readVideoPacketsFromFile 从h264 es流文件读取所有视频包
func readVideoPacketsFromFile(filename string) (videoContent []byte, videoPackets []base.AvPacket) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 将nal数据转换为lalserver要求的格式输入

// noop

// 1秒 / fps

// mergePackets 将音频队列和视频队列按时间戳有序合并为一个队列
func mergePackets(audioPackets, videoPackets []base.AvPacket) (packets []base.AvPacket) {
	_ = "STUB: not implemented"
	return nil

	// audio数组为空，将video的剩余数据取出，然后merge结束
}

//

// 音频和视频都有数据，取时间戳小的
