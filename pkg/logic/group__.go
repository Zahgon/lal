// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"sync"

	"github.com/q191201771/lal/pkg/gb28181"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/hls"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/lal/pkg/httpts"
	"github.com/q191201771/lal/pkg/mpegts"
	"github.com/q191201771/lal/pkg/remux"
	"github.com/q191201771/lal/pkg/rtmp"
	"github.com/q191201771/lal/pkg/rtsp"
	"github.com/q191201771/lal/pkg/sdp"
)

// ---------------------------------------------------------------------------------------------------------------------
// 输入流需要做的事情
// TODO(chef): [refactor] 考虑抽象出通用接口 202208
//
// checklist表格
// | .                                           | rtmp pub | ps pub |
// | 添加到group中                                | Y        | Y      |
// | 到输出流的转换路径关系                         | Y        | Y      |
// | 删除                                        | Y        | Y      |
// | group.hasPubSession()                       | Y        | Y      |
// | group.disposeInactiveSessions()检查超时并清理 | Y        | Y      |
// | group.Dispose()时销毁                        | Y        | Y      |
// | group.GetStat()时获取信息                     | Y        | Y      |
// | group.KickSession()时踢出                    | Y        | Y      |
// | group.updateAllSessionStat()更新信息         | Y        | Y      |
// | group.inSessionUniqueKey()                  | Y        | Y      |

// TODO(chef): [refactor] 整理sub类型流接入需要做的事情的文档 202211

// ---------------------------------------------------------------------------------------------------------------------
// 输入流到输出流的转换路径关系（一共6种输入）：
//
// rtmpPullSession.WithOnReadRtmpAvMsg  ->
// rtmpPubSession.SetPubSessionObserver ->
//    customizePubSession.WithOnRtmpMsg -> OnReadRtmpAvMsg(enter Lock) -> [dummyAudioFilter] -> broadcastByRtmpMsg -> rtmp, http-flv
//                                                                                                                 -> rtmp2RtspRemuxer -> rtsp
//                                                                                                                 -> rtmp2MpegtsRemuxer -> ts, hls
//
// ---------------------------------------------------------------------------------------------------------------------
// rtspPullSession ->
//  rtspPubSession -> OnRtpPacket(enter Lock) -> rtsp
//                 -> OnAvPacket(enter Lock) -> rtsp2RtmpRemuxer -> onRtmpMsgFromRemux -> [dummyAudioFilter] -> broadcastByRtmpMsg -> rtmp, http-flv
//                                                                                                           -> rtmp2MpegtsRemuxer -> ts, hls
//
// ---------------------------------------------------------------------------------------------------------------------
// psPubSession -> OnAvPacketFromPsPubSession(enter Lock) -> rtsp2RtmpRemuxer -> onRtmpMsgFromRemux -> [dummyAudioFilter] -> broadcastByRtmpMsg -> ...
//                                                                                                                                              -> ...
//                                                                                                                                              -> ...

type GroupOption struct {
	onHookSession func(uniqueKey string, streamName string) ICustomizeHookSessionContext
}

type IGroupObserver interface {
	CleanupHlsIfNeeded(appName string, streamName string, path string)
	OnHlsMakeTs(info base.HlsMakeTsInfo)
	OnRelayPullStart(info base.PullStartInfo) // TODO(chef): refactor me
	OnRelayPullStop(info base.PullStopInfo)
}

type Group struct {
	UniqueKey  string // const after init
	appName    string // const after init
	streamName string // const after init TODO chef: 和stat里的字段重复，可以删除掉
	config     *Config

	option                      GroupOption
	customizeHookSessionContext ICustomizeHookSessionContext

	observer IGroupObserver

	exitChan chan struct{}

	mutex sync.Mutex
	// pub
	rtmpPubSession      *rtmp.ServerSession
	rtspPubSession      *rtsp.PubSession
	customizePubSession *CustomizePubSessionContext
	psPubSession        *gb28181.PubSession
	rtsp2RtmpRemuxer    *remux.AvPacket2RtmpRemuxer // TODO(chef): [refactor] 重命名为avPacket2RtmpRemuxer，因为除了rtsp，customize pub和gb28181 pub都是 202208
	rtmp2RtspRemuxer    *remux.Rtmp2RtspRemuxer
	rtmp2MpegtsRemuxer  *remux.Rtmp2MpegtsRemuxer
	// pull
	pullProxy *pullProxy
	// rtmp pub使用 TODO(chef): [doc] 更新这个注释，是共同使用 202210
	dummyAudioFilter *remux.DummyAudioFilter
	// ps pub使用
	psPubTimeoutSec            uint32 // 超时时间
	psPubPrevInactiveCheckTick int64  // 上次检查时间
	// rtmp sub使用
	rtmpGopCache *remux.GopCache
	// httpflv sub使用
	httpflvGopCache *remux.GopCache
	// httpts sub使用
	httptsGopCache *remux.GopCacheMpegts
	// rtsp使用
	sdpCtx *sdp.LogicContext
	// mpegts使用
	patpmt []byte
	// sub
	rtmpSubSessionSet    map[*rtmp.ServerSession]struct{}
	httpflvSubSessionSet map[*httpflv.SubSession]struct{}
	httptsSubSessionSet  map[*httpts.SubSession]struct{}
	rtspSubSessionSet    map[*rtsp.SubSession]struct{} // 注意，使用这个容器时，一定要注意 session 的 Stage 属性
	hlsSubSessionSet     map[*hls.SubSession]struct{}
	// push
	pushEnable    bool
	url2PushProxy map[string]*pushProxy
	// hls
	hlsMuxer *hls.Muxer
	// record
	recordFlv    *httpflv.FlvFileWriter
	recordMpegts *mpegts.FileWriter
	// rtmp sub使用
	rtmpMergeWriter *base.MergeWriter // TODO(chef): 后面可以在业务层加一个定时Flush
	//
	stat base.StatGroup
	//
	inVideoFpsRecords base.PeriodRecord
	//
	hlsCalcSessionStatIntervalSec uint32
	//
	psPubDumpFile    *base.DumpFile
	rtspPullDumpFile *base.DumpFile
}

func NewGroup(appName string, streamName string, config *Config, option GroupOption, observer IGroupObserver) *Group {
	_ = "STUB: not implemented"
	return nil
}

// equals to (ms/1000) * 10

func (group *Group) RunLoop() {
	_ = "STUB: not implemented"

	// Tick 定时器
	//
	// @param tickCount 当前时间，单位秒。注意，不一定是Unix时间戳，可以是从0开始+1秒递增的时间
	return
}

func (group *Group) Tick(tickCount uint32) { _ = "STUB: not implemented"; return }

// 定时关闭没有数据的session

// 定时计算session bitrate

// because hls make multiple separate http request to get stream content and gap between request base on hls segment duration
// if we update every 5s can cause bitrateKbit equal to 0 if within 5s do not have any ts http request is make

// Dispose ...
func (group *Group) Dispose() { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

func (group *Group) StringifyDebugStats(maxsub int) string { _ = "STUB: not implemented"; return "" }

func (group *Group) GetStat(maxsub int) base.StatGroup {
	_ = "STUB: not implemented"
	// TODO(chef): [refactor] param maxsub
	return *new(base.StatGroup)
}

//Log.Debugf("GetStat. len(group.hlsSubSessionSet)=%d", len(group.hlsSubSessionSet))

func (group *Group) KickSession(sessionId string) bool { _ = "STUB: not implemented"; return false }

// TODO chef: 考虑数据结构改成sessionIdzuokey的map

func (group *Group) IsInactive() bool { _ = "STUB: not implemented"; return false }

func (group *Group) HasInSession() bool { _ = "STUB: not implemented"; return false }

func (group *Group) HasOutSession() bool { _ = "STUB: not implemented"; return false }

func (group *Group) OutSessionNum() int {
	_ = "STUB: not implemented"
	// TODO(chef): 没有包含hls的播放者
	return 0
}

// TODO(chef): [refactor] 考虑只判断session是否为nil 202205

// ---------------------------------------------------------------------------------------------------------------------

// disposeInactiveSessions 关闭不活跃的session
func (group *Group) disposeInactiveSessions(tickCount uint32) { _ = "STUB: not implemented"; return }

// noop
// 没有超时逻辑

// 以下都是以 CheckSessionAliveIntervalSec 为间隔的清理逻辑

// updateAllSessionStat 更新所有session的状态
func (group *Group) updateAllSessionStat() { _ = "STUB: not implemented"; return }

func (group *Group) hasPubSession() bool { _ = "STUB: not implemented"; return false }

func (group *Group) hasSubSession() bool { _ = "STUB: not implemented"; return false }

func (group *Group) hasPushSession() bool { _ = "STUB: not implemented"; return false }

func (group *Group) hasInSession() bool { _ = "STUB: not implemented"; return false }

// hasOutSession 是否还有out往外发送音视频数据的session
func (group *Group) hasOutSession() bool { _ = "STUB: not implemented"; return false }

// isTotalEmpty 当前group是否完全没有流了
func (group *Group) isTotalEmpty() bool { _ = "STUB: not implemented"; return false }

func (group *Group) inSessionUniqueKey() string { _ = "STUB: not implemented"; return "" }

func (group *Group) shouldStartRtspRemuxer() bool { _ = "STUB: not implemented"; return false }

func (group *Group) shouldStartMpegtsRemuxer() bool { _ = "STUB: not implemented"; return false }

func (group *Group) OnHlsMakeTs(info base.HlsMakeTsInfo) { _ = "STUB: not implemented"; return }
