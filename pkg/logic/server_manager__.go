// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"net/http"
	_ "net/http/pprof"
	"sync"

	"github.com/q191201771/naza/pkg/taskpool"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/hls"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/lal/pkg/httpts"
	"github.com/q191201771/lal/pkg/rtmp"
	"github.com/q191201771/lal/pkg/rtsp"
	//"github.com/felixge/fgprof"
)

type ServerManager struct {
	option          Option
	serverStartTime string
	config          *Config

	httpServerManager *base.HttpServerManager
	httpServerHandler *HttpServerHandler
	hlsServerHandler  *hls.ServerHandler

	rtmpServer    *rtmp.Server
	rtmpsServer   *rtmp.Server
	rtspServer    *rtsp.Server
	rtspsServer   *rtsp.Server
	httpApiServer *HttpApiServer
	pprofServer   *http.Server
	wsrtspServer  *rtsp.WebsocketServer
	exitChan      chan struct{}

	mutex        sync.Mutex
	groupManager IGroupManager

	onHookSession func(uniqueKey string, streamName string) ICustomizeHookSessionContext

	notifyHandlerThread taskpool.Pool

	ipBlacklist IpBlacklist
}

func NewServerManager(modOption ...ModOption) *ServerManager { _ = "STUB: not implemented"; return nil }

// ----- implement ILalServer interface --------------------------------------------------------------------------------

func (sm *ServerManager) RunLoop() error {
	_ = "STUB: not implemented"
	// TODO(chef): 作为阻塞函数，外部只能获取失败或结束的信息，没法获取到启动成功的信息
	return nil
}

//Log.Warn("start fgprof.")
//http.DefaultServeMux.Handle("/debug/fgprof", fgprof.Handler())

// rtmps启动失败影响降级：当rtmps启动时我们并不返回错误，保证不因为rtmps影响其他服务

// rtsps启动失败影响降级：当rtsps启动时我们并不返回错误，保证不因为rtsps影响其他服务

// 关闭空闲的group

// 定时打印一些group相关的debug日志

// 定时通过http notify发送group相关的信息

// never reach here

func (sm *ServerManager) Dispose() { _ = "STUB: not implemented"; return }

//if sm.hlsServer != nil {
//	sm.hlsServer.Dispose()
//}

// ---------------------------------------------------------------------------------------------------------------------

func (sm *ServerManager) AddCustomizePubSession(streamName string) (ICustomizePubSessionContext, error) {
	_ = "STUB: not implemented"
	return *new(ICustomizePubSessionContext), nil
}

func (sm *ServerManager) DelCustomizePubSession(sessionCtx ICustomizePubSessionContext) {
	_ = "STUB: not implemented"
	return
}

func (sm *ServerManager) WithOnHookSession(onHookSession func(uniqueKey string, streamName string) ICustomizeHookSessionContext) {
	_ = "STUB: not implemented"
	return
}

// ----- implement rtmp.IServerObserver interface -----------------------------------------------------------------------

func (sm *ServerManager) OnRtmpConnect(session *rtmp.ServerSession, opa rtmp.ObjectPairArray) {
	_ = "STUB: not implemented"
	return
}

func (sm *ServerManager) OnNewRtmpPubSession(session *rtmp.ServerSession) error {
	_ = "STUB: not implemented"
	return nil
}

// 先做simple auth鉴权

func (sm *ServerManager) OnDelRtmpPubSession(session *rtmp.ServerSession) {
	_ = "STUB: not implemented"
	return
}

func (sm *ServerManager) OnNewRtmpSubSession(session *rtmp.ServerSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (sm *ServerManager) OnDelRtmpSubSession(session *rtmp.ServerSession) {
	_ = "STUB: not implemented"
	return
}

// ----- implement IHttpServerHandlerObserver interface -----------------------------------------------------------------

func (sm *ServerManager) OnNewHttpflvSubSession(session *httpflv.SubSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (sm *ServerManager) OnDelHttpflvSubSession(session *httpflv.SubSession) {
	_ = "STUB: not implemented"
	return
}

func (sm *ServerManager) OnNewHttptsSubSession(session *httpts.SubSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (sm *ServerManager) OnDelHttptsSubSession(session *httpts.SubSession) {
	_ = "STUB: not implemented"
	return
}

// ----- implement rtsp.IServerObserver interface -----------------------------------------------------------------------

func (sm *ServerManager) OnNewRtspSessionConnect(session *rtsp.ServerCommandSession) {
	_ = "STUB: not implemented"
	// TODO chef: impl me
	return
}

func (sm *ServerManager) OnDelRtspSession(session *rtsp.ServerCommandSession) {
	_ = "STUB: not implemented"
	// TODO chef: impl me
	return
}

func (sm *ServerManager) OnNewRtspPubSession(session *rtsp.PubSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (sm *ServerManager) OnDelRtspPubSession(session *rtsp.PubSession) {
	_ = "STUB: not implemented"
	return
}

func (sm *ServerManager) OnNewRtspSubSessionDescribe(session *rtsp.SubSession) (ok bool, sdp []byte) {
	_ = "STUB: not implemented"
	return false, nil
}

func (sm *ServerManager) OnNewRtspSubSessionPlay(session *rtsp.SubSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (sm *ServerManager) OnDelRtspSubSession(session *rtsp.SubSession) {
	_ = "STUB: not implemented"
	return
}

func (sm *ServerManager) OnNewHlsSubSession(session *hls.SubSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (sm *ServerManager) OnDelHlsSubSession(session *hls.SubSession) {
	_ = "STUB: not implemented"
	return
}

// ----- implement IGroupCreator interface -----------------------------------------------------------------------------

func (sm *ServerManager) CreateGroup(appName string, streamName string) *Group {
	_ = "STUB: not implemented"
	return nil
}

// ----- implement IGroupObserver interface -----------------------------------------------------------------------------

func (sm *ServerManager) CleanupHlsIfNeeded(appName string, streamName string, path string) {
	_ = "STUB: not implemented"
	return
}

func (sm *ServerManager) OnRelayPullStart(info base.PullStartInfo) {
	_ = "STUB: not implemented"
	return
}

func (sm *ServerManager) OnRelayPullStop(info base.PullStopInfo) { _ = "STUB: not implemented"; return }

func (sm *ServerManager) OnHlsMakeTs(info base.HlsMakeTsInfo) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

func (sm *ServerManager) Config() *Config { _ = "STUB: not implemented"; return nil }

func (sm *ServerManager) GetGroup(appName string, streamName string) *Group {
	_ = "STUB: not implemented"
	return nil
}

// ----- private method ------------------------------------------------------------------------------------------------

// 注意，函数内部不加锁，由调用方保证加锁进入
func (sm *ServerManager) getOrCreateGroup(appName string, streamName string) *Group {
	_ = "STUB: not implemented"
	return nil
}

func (sm *ServerManager) getGroup(appName string, streamName string) *Group {
	_ = "STUB: not implemented"
	return nil
}

func (sm *ServerManager) serveHls(writer http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// TODO(chef): [refactor] 需要整理，这里使用 hls.PathStrategy 不太好 202207

//Log.Warnf("found %s in ip blacklist, so do not serve this request.", remoteIp)
