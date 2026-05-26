// Copyright 2023, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package base

import "net/http"

var (
	CorsHeaders = "Access-Control-Allow-Credentials: true\r\n" +
		"Access-Control-Allow-Origin: *\r\n" +
		"Access-Control-Allow-Headers: Content-Type\r\n"
)

func AddCorsHeaders2HlsIfNeeded(w http.ResponseWriter) {
	_ = "STUB: not implemented"
	// TODO(chef): [opt] 为其他协议也增加配置项 202308
	return
}

func AddCorsHeaders(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

//resp.Header().Add("Access-Control-Allow-Origin", "*")
//resp.Header().Add("Access-Control-Allow-Credentials", "true")
//resp.Header().Add("Access-Control-Allow-Methods", "*")
//resp.Header().Add("Access-Control-Allow-Headers", "Content-Type,Access-Token")
//resp.Header().Add("Access-Control-Allow-Expose-Headers", "*")
