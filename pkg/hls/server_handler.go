// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package hls

import (
	"net/http"
	"sync"
	"time"

	"github.com/q191201771/lal/pkg/base"
)

type IHlsServerHandlerObserver interface {
	OnNewHlsSubSession(session *SubSession) error
	OnDelHlsSubSession(session *SubSession)
}

type ServerHandler struct {
	outPath           string
	observer          IHlsServerHandlerObserver
	urlPattern        string
	sessionMap        map[string]*SubSession
	mutex             sync.Mutex
	subSessionTimeout time.Duration
	subSessionHashKey string
}

func NewServerHandler(outPath, urlPattern, subSessionHashKey string, subSessionTimeoutMs int, observer IHlsServerHandlerObserver) *ServerHandler {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServerHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *ServerHandler) ServeHTTPWithUrlCtx(resp http.ResponseWriter, req *http.Request, urlCtx base.UrlContext) {
	_ = "STUB: not implemented"
	//Log.Debugf("%+v", req)
	return
}

// TODO chef:
// - check appname in URI path

// 如果开启了hls sub session功能

// 创建session对象，并让m3u8跳转到携带session_id的url请求

//Log.Debugf("%+v", ri)

// 给ts文件都携带上session_id字段

// getSubSession 获取 SubSession，如果不存在，返回nil
func (s *ServerHandler) getSubSession(sessionIdHash string) *SubSession {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServerHandler) createSubSession(req *http.Request, urlCtx base.UrlContext) (*SubSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// keepSessionAlive 标记延长session存活时间，如果session不存在，返回 base.ErrHlsSessionNotFound
func (s *ServerHandler) keepSessionAlive(sessionIdHash string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServerHandler) clearExpireSession() { _ = "STUB: not implemented"; return }

func (s *ServerHandler) CloseSubSessionIfExist(req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *ServerHandler) isSubSessionModeEnable() bool { _ = "STUB: not implemented"; return false }

func (s *ServerHandler) runLoop() {
	_ = "STUB: not implemented"
	// TODO(chef): [refactor] 也许可以弄到group中管理超时，和其他协议的session管理方式保持一致 202211
	return
}

// m3u8文件用这个也行
//resp.Header().Add("Content-Type", "application/vnd.apple.mpegurl")
