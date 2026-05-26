// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

import (
	"net"
)

type IServerObserver interface {
	// OnNewRtspSessionConnect @brief 使得上层有能力管理未进化到Pub、Sub阶段的Session
	OnNewRtspSessionConnect(session *ServerCommandSession)

	// OnDelRtspSession @brief 注意，对于已经进化到了Pub、Sub阶段的Session，该回调依然会被调用
	OnDelRtspSession(session *ServerCommandSession)

	///////////////////////////////////////////////////////////////////////////

	// OnNewRtspPubSession
	//
	// @brief  Announce阶段回调
	// @return 如果返回非nil，则表示上层要强制关闭这个推流请求
	//
	OnNewRtspPubSession(session *PubSession) error

	OnDelRtspPubSession(session *PubSession)

	///////////////////////////////////////////////////////////////////////////

	// OnNewRtspSubSessionDescribe
	//
	// @return 如果返回false，则表示上层要强制关闭这个拉流请求
	// @return sdp
	//
	OnNewRtspSubSessionDescribe(session *SubSession) (ok bool, sdp []byte)

	// OnNewRtspSubSessionPlay
	//
	// @brief Play阶段回调
	// @return ok  如果返回非nil，则表示上层要强制关闭这个拉流请求
	//
	OnNewRtspSubSessionPlay(session *SubSession) error

	OnDelRtspSubSession(session *SubSession)
}

type ServerAuthConfig struct {
	AuthEnable bool   `json:"auth_enable"`
	AuthMethod int    `json:"auth_method"`
	UserName   string `json:"username"`
	PassWord   string `json:"password"`
}

type Server struct {
	addr     string
	observer IServerObserver

	ln   net.Listener
	auth ServerAuthConfig
}

func NewServer(addr string, observer IServerObserver, auth ServerAuthConfig) *Server {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) Listen() (err error) { _ = "STUB: not implemented"; return nil }

func (s *Server) ListenWithTLS(certFile, keyFile string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) RunLoop() error { _ = "STUB: not implemented"; return nil }

func (s *Server) Dispose() { _ = "STUB: not implemented"; return }

// ----- ServerCommandSessionObserver ----------------------------------------------------------------------------------

func (s *Server) OnNewRtspPubSession(session *PubSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) OnNewRtspSubSessionDescribe(session *SubSession) (ok bool, sdp []byte) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Server) OnNewRtspSubSessionPlay(session *SubSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) OnDelRtspPubSession(session *PubSession) { _ = "STUB: not implemented"; return }

func (s *Server) OnDelRtspSubSession(session *SubSession) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

func (s *Server) handleTcpConnect(conn net.Conn) { _ = "STUB: not implemented"; return }
