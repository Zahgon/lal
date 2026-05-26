// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package innertest

import (
	"testing"

	"github.com/q191201771/lal/pkg/rtprtcp"
	"github.com/q191201771/lal/pkg/rtsp"
	"github.com/q191201771/lal/pkg/sdp"

	"github.com/q191201771/lal/pkg/base"

	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/lal/pkg/rtmp"
	"github.com/q191201771/naza/pkg/nazaatomic"
)

// 开启了一个lalserver
// rtmp pub              读取flv文件，使用rtmp协议推送至服务端
// rtmp sub, httpflv sub 分别用rtmp协议以及httpflv协议从服务端拉流，再将拉取的流保存为flv文件
// 对比三份flv文件，看是否完全一致
// hls                   并检查hls生成的m3u8和ts文件，是否和之前的完全一致

// TODO chef:
// - 加上relay push
// - 加上relay pull

var (
	t *testing.T

	mode         int // 0 正常 1 输入只有音频 2 输入只有视频
	confFilename = "../../testdata/lalserver.conf.json"
	rFlvFileName = "../../testdata/test.flv"

	pushUrl        string
	httpflvPullUrl string
	httptsPullUrl  string
	rtmpPullUrl    string
	rtspPullUrl    string

	wRtmpPullFileName     string
	wFlvPullFileName      string
	wPlaylistM3u8FileName string
	wRecordM3u8FileName   string
	wHlsTsFilePath        string
	wTsPullFileName       string

	fileTagCount          int
	httpflvPullTagCount   nazaatomic.Uint32
	rtmpPullTagCount      nazaatomic.Uint32
	httptsSize            nazaatomic.Uint32
	rtspSdpCtx            sdp.LogicContext
	rtspPullAvPacketCount nazaatomic.Uint32

	httpFlvWriter httpflv.FlvFileWriter
	rtmpWriter    httpflv.FlvFileWriter

	pushSession        *rtmp.PushSession
	httpflvPullSession *httpflv.PullSession
	rtmpPullSession    *rtmp.PullSession
	rtspPullSession    *rtsp.PullSession
)

type RtspPullObserver struct {
}

func (r RtspPullObserver) OnSdp(sdpCtx sdp.LogicContext) { _ = "STUB: not implemented"; return }

func (r RtspPullObserver) OnRtpPacket(pkt rtprtcp.RtpPacket) { _ = "STUB: not implemented"; return }

func (r RtspPullObserver) OnAvPacket(pkt base.AvPacket) { _ = "STUB: not implemented"; return }

func Entry(tt *testing.T) {
	_ = "STUB: not implemented"
	// 在MacOS只测试一次
	// 其他环境（比如github CI）上则每个package都执行，因为要生产测试覆盖率
	return
}

func entry() { _ = "STUB: not implemented"; return }

// 用于等待所有协程结束

// TODO(chef): [test] rtsp sub没有验证收到的数据，因为即使是先sub，它还有一个数据到来后，才能完成信令交互的逻辑 202206
// TODO(chef): [perf] [2021.12.25] rtmp推rtsp拉的性能。开启rtsp pull后，rtmp pull的总时长增加了

//option.WriteChanSize = 1024

//Log.Debugf("rtmp push: %d", fileTagCount.Load())

// 注意，先释放push，触发pub释放，从而刷新hls的结束时切片逻辑

// 由于windows没有信号，会导致编译错误，所以直接调用Dispose
//_ = syscall.Kill(syscall.Getpid(), syscall.SIGUSR1)

func compareFile() { _ = "STUB: not implemented"; return }

// 检查httpflv

// 检查rtmp

// 检查hls的m3u8文件

// 检查hls的ts文件

func getAllHttpApi(addr string) { _ = "STUB: not implemented"; return }

func getHttpts() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ---------------------------------------------------------------------------------------------------------------------

// TODO(chef): refactor 移入naza中

func httpGet(url string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func httpPost(url string, info interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ---------------------------------------------------------------------------------------------------------------------

var (
	goldenRtmpLenList = []int{2120047, 504722, 1615715}
	goldenRtmpMd5List = []string{
		"7d68f0e2ab85c1992f70740479c8d3db",
		"b889f690e07399c8c8353a3b1dba7efb",
		"b5a9759455039761b6d4dd3ed8e97634",
	}

	goldenHttpflvLenList = []int{2120047, 504722, 1615715}
	goldenHttpflvMd5List = []string{
		"7d68f0e2ab85c1992f70740479c8d3db",
		"b889f690e07399c8c8353a3b1dba7efb",
		"b5a9759455039761b6d4dd3ed8e97634",
	}

	goldenHlsTsNumList = []int{8, 10, 8}
	goldenHlsTsLenList = []int{2219152, 525648, 1696512}
	goldenHlsTsMd5List = []string{
		"5a05b84486382f3b8d2a0f9e75c67623",
		"f03c5ab24dddd8875f06cdea605cd87c",
		"ceec699eae6671507376c26ac1cdaba4",
	}

	goldenHttptsLenList = []int{2216332, 522264, 1693880}
	goldenHttptsMd5List = []string{
		"3ba4baa20df968196eac75d96f8041b5",
		"0ec316060f1aeeb1d283b30158c2eeb8",
		"e0d90dd5efd1119f1f66db8ac4cf5d48",
	}
)

var goldenPlaylistM3u8List = []string{
	`#EXTM3U
#EXT-X-VERSION:3
#EXT-X-ALLOW-CACHE:NO
#EXT-X-TARGETDURATION:5
#EXT-X-MEDIA-SEQUENCE:2

#EXTINF:3.333,
innertest-1642375465000-2.ts
#EXTINF:4.000,
innertest-1642375465000-3.ts
#EXTINF:4.867,
innertest-1642375465000-4.ts
#EXTINF:3.133,
innertest-1642375465000-5.ts
#EXTINF:4.000,
innertest-1642375465000-6.ts
#EXTINF:2.644,
innertest-1642375465000-7.ts
#EXT-X-ENDLIST
`,
	`#EXTM3U
#EXT-X-VERSION:3
#EXT-X-ALLOW-CACHE:NO
#EXT-X-TARGETDURATION:3
#EXT-X-MEDIA-SEQUENCE:4

#EXTINF:3.088,
innertest-1642375465000-4.ts
#EXTINF:3.088,
innertest-1642375465000-5.ts
#EXTINF:3.089,
innertest-1642375465000-6.ts
#EXTINF:3.088,
innertest-1642375465000-7.ts
#EXTINF:3.088,
innertest-1642375465000-8.ts
#EXTINF:2.113,
innertest-1642375465000-9.ts
#EXT-X-ENDLIST
`,
	`#EXTM3U
#EXT-X-VERSION:3
#EXT-X-ALLOW-CACHE:NO
#EXT-X-TARGETDURATION:5
#EXT-X-MEDIA-SEQUENCE:2

#EXTINF:3.333,
innertest-1642375465000-2.ts
#EXTINF:4.000,
innertest-1642375465000-3.ts
#EXTINF:4.867,
innertest-1642375465000-4.ts
#EXTINF:3.133,
innertest-1642375465000-5.ts
#EXTINF:4.000,
innertest-1642375465000-6.ts
#EXTINF:2.600,
innertest-1642375465000-7.ts
#EXT-X-ENDLIST
`,
}

var goldenRecordM3u8List = []string{
	`#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:5
#EXT-X-MEDIA-SEQUENCE:0

#EXT-X-DISCONTINUITY
#EXTINF:4.000,
innertest-1642375465000-0.ts
#EXTINF:4.000,
innertest-1642375465000-1.ts
#EXTINF:3.333,
innertest-1642375465000-2.ts
#EXTINF:4.000,
innertest-1642375465000-3.ts
#EXTINF:4.867,
innertest-1642375465000-4.ts
#EXTINF:3.133,
innertest-1642375465000-5.ts
#EXTINF:4.000,
innertest-1642375465000-6.ts
#EXTINF:2.644,
innertest-1642375465000-7.ts
#EXT-X-ENDLIST
`,
	`#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:3
#EXT-X-MEDIA-SEQUENCE:0

#EXT-X-DISCONTINUITY
#EXTINF:3.088,
innertest-1642375465000-0.ts
#EXTINF:3.088,
innertest-1642375465000-1.ts
#EXTINF:3.089,
innertest-1642375465000-2.ts
#EXTINF:3.088,
innertest-1642375465000-3.ts
#EXTINF:3.088,
innertest-1642375465000-4.ts
#EXTINF:3.088,
innertest-1642375465000-5.ts
#EXTINF:3.089,
innertest-1642375465000-6.ts
#EXTINF:3.088,
innertest-1642375465000-7.ts
#EXTINF:3.088,
innertest-1642375465000-8.ts
#EXTINF:2.113,
innertest-1642375465000-9.ts
#EXT-X-ENDLIST
`,
	`#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:5
#EXT-X-MEDIA-SEQUENCE:0

#EXT-X-DISCONTINUITY
#EXTINF:4.000,
innertest-1642375465000-0.ts
#EXTINF:4.000,
innertest-1642375465000-1.ts
#EXTINF:3.333,
innertest-1642375465000-2.ts
#EXTINF:4.000,
innertest-1642375465000-3.ts
#EXTINF:4.867,
innertest-1642375465000-4.ts
#EXTINF:3.133,
innertest-1642375465000-5.ts
#EXTINF:4.000,
innertest-1642375465000-6.ts
#EXTINF:2.600,
innertest-1642375465000-7.ts
#EXT-X-ENDLIST
`,
}

var goldenRtspSdpTmplList = []string{
	`v=0
o=- 0 0 IN IP4 127.0.0.1
s=No Name
c=IN IP4 127.0.0.1
t=0 0
a=tool:{atoolv}
m=video 0 RTP/AVP 96
a=rtpmap:96 H264/90000
a=fmtp:96 packetization-mode=1; sprop-parameter-sets=Z2QAFqyyAUBf8uAiAAADAAIAAAMAPB4sXJA=,aOvDyyLA; profile-level-id=640016
a=control:streamid=0
m=audio 0 RTP/AVP 97
b=AS:128
a=rtpmap:97 MPEG4-GENERIC/44100/2
a=fmtp:97 profile-level-id=1;mode=AAC-hbr;sizelength=13;indexlength=3;indexdeltalength=3; config=121056e500
a=control:streamid=1
`,
	`v=0
o=- 0 0 IN IP4 127.0.0.1
s=No Name
c=IN IP4 127.0.0.1
t=0 0
a=tool:{atoolv}
m=audio 0 RTP/AVP 97
b=AS:128
a=rtpmap:97 MPEG4-GENERIC/44100/2
a=fmtp:97 profile-level-id=1;mode=AAC-hbr;sizelength=13;indexlength=3;indexdeltalength=3; config=121056e500
a=control:streamid=0
`,
	`v=0
o=- 0 0 IN IP4 127.0.0.1
s=No Name
c=IN IP4 127.0.0.1
t=0 0
a=tool:{atoolv}
m=video 0 RTP/AVP 96
a=rtpmap:96 H264/90000
a=fmtp:96 packetization-mode=1; sprop-parameter-sets=Z2QAFqyyAUBf8uAiAAADAAIAAAMAPB4sXJA=,aOvDyyLA; profile-level-id=640016
a=control:streamid=0
`,
}
