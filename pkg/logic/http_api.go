// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	_ "embed"
	"net"
	"net/http"

	"github.com/q191201771/naza/pkg/nazajson"
)

//go:embed http_an__lal.html
var webUITpl string

type HttpApiServer struct {
	addr string
	sm   *ServerManager

	ln net.Listener
}

func NewHttpApiServer(addr string, sm *ServerManager) *HttpApiServer {
	_ = "STUB: not implemented"
	return nil
}

func (h *HttpApiServer) Listen() (err error) { _ = "STUB: not implemented"; return nil }

func (h *HttpApiServer) RunLoop() error { _ = "STUB: not implemented"; return nil }

// 所有没有注册路由的走下面这个处理函数

// TODO chef: dispose

// ---------------------------------------------------------------------------------------------------------------------

func (h *HttpApiServer) statLalInfoHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *HttpApiServer) statAllGroupHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *HttpApiServer) statGroupHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// ---------------------------------------------------------------------------------------------------------------------

func (h *HttpApiServer) ctrlStartRelayPullHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *HttpApiServer) ctrlStopRelayPullHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *HttpApiServer) ctrlKickSessionHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *HttpApiServer) ctrlStartRtpPubHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// 不存在时默认0值的，不需要手动写了
//if !j.Exist("port") {
//	info.Port = 0
//}
//if !j.Exist("is_tcp_flag") {
//	info.IsTcpFlag = 0
//}

func (h *HttpApiServer) ctrlAddIpBlacklistHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *HttpApiServer) webUIHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *HttpApiServer) notFoundHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

//w.WriteHeader(http.StatusNotFound)

// ---------------------------------------------------------------------------------------------------------------------

func feedback(v interface{}, w http.ResponseWriter) { _ = "STUB: not implemented"; return }

// unmarshalRequestJsonBody
//
// TODO(chef): [refactor] 搬到naza中 202205
func unmarshalRequestJsonBody(r *http.Request, info interface{}, keyFieldList ...string) (nazajson.Json, error) {
	_ = "STUB: not implemented"
	return *new(nazajson.Json), nil
}
