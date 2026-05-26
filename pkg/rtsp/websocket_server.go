// Copyright 2024, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

import (
	"net"
	"net/http"
)

type WebsocketServer struct {
	addr     string
	observer IServerObserver

	ln         net.Listener
	auth       ServerAuthConfig
	httpServer http.Server
}

func NewWebsocketServer(addr string, observer IServerObserver, auth ServerAuthConfig) *WebsocketServer {
	_ = "STUB: not implemented"
	return nil
}

func (s *WebsocketServer) Listen() (err error) { _ = "STUB: not implemented"; return nil }

func (s *WebsocketServer) HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// 火狐浏览器 Connection = [keep-alive, Upgrade]

func (s *WebsocketServer) Dispose() { _ = "STUB: not implemented"; return }

// ----- ServerCommandSessionObserver ----------------------------------------------------------------------------------

func (s *WebsocketServer) OnNewRtspPubSession(session *PubSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *WebsocketServer) OnNewRtspSubSessionDescribe(session *SubSession) (ok bool, sdp []byte) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *WebsocketServer) OnNewRtspSubSessionPlay(session *SubSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *WebsocketServer) OnDelRtspPubSession(session *PubSession) {
	_ = "STUB: not implemented"
	return
}

func (s *WebsocketServer) OnDelRtspSubSession(session *SubSession) {
	_ = "STUB: not implemented"
	return
}
