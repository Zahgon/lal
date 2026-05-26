// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtmp

import (
	"net"
)

type IServerObserver interface {
	OnRtmpConnect(session *ServerSession, opa ObjectPairArray)

	// OnNewRtmpPubSession
	//
	// 上层代码应该在这个事件回调中注册音视频数据的监听
	//
	// @return 上层如果想关闭这个session，则回调中返回不为nil的error值
	//
	OnNewRtmpPubSession(session *ServerSession) error

	// OnDelRtmpPubSession
	//
	// 注意，如果session是上层通过 OnNewRtmpPubSession 回调的返回值关闭的，则该session不再触发这个逻辑
	//
	OnDelRtmpPubSession(session *ServerSession)

	OnNewRtmpSubSession(session *ServerSession) error
	OnDelRtmpSubSession(session *ServerSession)
}

type Server struct {
	addr     string
	observer IServerObserver
	ln       net.Listener
}

func NewServer(addr string, observer IServerObserver) *Server {
	_ = "STUB: not implemented"
	return nil
}

func (server *Server) Listen() (err error) { _ = "STUB: not implemented"; return nil }

func (server *Server) ListenWithTLS(certFile, keyFile string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (server *Server) RunLoop() error { _ = "STUB: not implemented"; return nil }

func (server *Server) Dispose() { _ = "STUB: not implemented"; return }

func (server *Server) handleTcpConnect(conn net.Conn) { _ = "STUB: not implemented"; return }

// ----- IServerSessionObserver ------------------------------------------------------------------------------------

func (server *Server) OnRtmpConnect(session *ServerSession, opa ObjectPairArray) {
	_ = "STUB: not implemented"
	return
}

func (server *Server) OnNewRtmpPubSession(session *ServerSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (server *Server) OnNewRtmpSubSession(session *ServerSession) error {
	_ = "STUB: not implemented"
	return nil
}
