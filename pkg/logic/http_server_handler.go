// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"net/http"

	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/lal/pkg/httpts"
)

type IHttpServerHandlerObserver interface {
	// OnNewHttpflvSubSession
	//
	// 通知上层有新的拉流者
	//
	// @return nil则允许拉流，不为nil则关闭连接
	//
	OnNewHttpflvSubSession(session *httpflv.SubSession) error
	OnDelHttpflvSubSession(session *httpflv.SubSession)

	OnNewHttptsSubSession(session *httpts.SubSession) error
	OnDelHttptsSubSession(session *httpts.SubSession)
}

type HttpServerHandler struct {
	observer IHttpServerHandlerObserver
}

func NewHttpServerHandler(observer IHttpServerHandlerObserver) *HttpServerHandler {
	_ = "STUB: not implemented"
	return nil
}

func (h *HttpServerHandler) ServeSubSession(writer http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// 火狐浏览器 Connection = [keep-alive, Upgrade]
