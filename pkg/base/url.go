// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package base

import (
	"net/http"
	"net/url"
)

// 见单元测试

// TODO chef: 考虑部分内容移入naza中

const (
	DefaultRtmpPort  = 1935
	DefaultHttpPort  = 80
	DefaultHttpsPort = 443
	DefaultRtspPort  = 554
	DefaultRtmpsPort = 443
	DefaultRtspsPort = 322
)

type UrlPathContext struct {
	PathWithRawQuery    string
	Path                string
	PathWithoutLastItem string // 注意，没有前面的'/'，也没有后面的'/'
	LastItemOfPath      string // 注意，没有前面的'/'
	RawQuery            string
}

type UrlContext struct {
	Url string

	Scheme       string
	Username     string
	Password     string
	StdHost      string // host or host:port
	HostWithPort string // 当原始url中不包含port时，填充scheme对应的默认port
	Host         string // 不包含port
	Port         int    // 当原始url中不包含port时，填充scheme对应的默认port

	//UrlPathContext
	PathWithRawQuery    string // 注意，有前面的'/'
	Path                string // 注意，有前面的'/'
	PathWithoutLastItem string // 注意，没有前面的'/'，也没有后面的'/'
	LastItemOfPath      string // 注意，没有前面的'/'
	RawQuery            string // 参数，注意，没有前面的'?'

	RawUrlWithoutUserInfo string

	filenameWithoutType string
	fileType            string
}

func (u *UrlContext) GetFilenameWithoutType() string { _ = "STUB: not implemented"; return "" }

func (u *UrlContext) GetFileType() string { _ = "STUB: not implemented"; return "" }

func (u *UrlContext) calcFilenameAndTypeIfNeeded() { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

// ParseUrl
//
// @param defaultPort:
// 注意，如果rawUrl中显示指定了端口，则该参数不生效。
// 注意，如果设置为-1，内部依然会对常见协议(http, https, rtmp, rtsp)设置官方默认端口。
func ParseUrl(rawUrl string, defaultPort int) (ctx UrlContext, err error) {
	_ = "STUB: not implemented"
	return *new(UrlContext), nil
}

// 如果不存在，则设置默认的

// TODO(chef): 测试大小写的情况

// url中端口不存在

// 端口存在

// ---------------------------------------------------------------------------------------------------------------------

func ParseRtmpUrl(rawUrl string) (ctx UrlContext, err error) {
	_ = "STUB: not implemented"
	return *new(UrlContext), nil
}

// 处理特殊case，具体见 testParseRtmpUrlCase1
// 注意，使用ffmpeg推流时，会把`rtmp://127.0.0.1/test110`中的test110作为appName(streamName则为空)
// 这种其实已不算十分合法的rtmp url了
// 我们这里也处理一下，和ffmpeg保持一致

// 处理特殊case, 具体见 testParseRtmpUrlCase2
//
// PathWithRawQuery:/vyun?vhost=thirdVhost?token=88F4/lss_7
//
// Path:/vyun-----------------------------------------------> /vyun?vhost=thirdVhost?token=88F4/lss_7
// PathWithoutLastItem:vyun---------------------------------> vyun?vhost=thirdVhost?token=88F4
// LastItemOfPath:------------------------------------------> lss_7
// RawQuery:vhost=thirdVhost?token=88F4/lss_7---------------> 空
//

func ParseRtspUrl(rawUrl string) (ctx UrlContext, err error) {
	_ = "STUB: not implemented"
	return *new(UrlContext), nil
}

// 注意，存在一种情况，使用rtsp pull session，直接拉取没有url path的流，所以不检查ctx.Path

func ParseHttpflvUrl(rawUrl string) (ctx UrlContext, err error) {
	_ = "STUB: not implemented"
	return *new(UrlContext), nil
}

// ---------------------------------------------------------------------------------------------------------------------

// ParseHttpRequest
//
// @return 完整url
func ParseHttpRequest(req *http.Request) string {
	_ = "STUB: not implemented"
	// TODO(chef): [refactor] scheme是否能从从req.URL.Scheme获取
	return ""
}

// ----- private -------------------------------------------------------------------------------------------------------

func parseUrlPath(stdUrl *url.URL) (ctx UrlPathContext, err error) {
	_ = "STUB: not implemented"
	return *new(UrlPathContext), nil
}

func parseHttpUrl(rawUrl string, filetype string) (ctx UrlContext, err error) {
	_ = "STUB: not implemented"
	return *new(UrlContext), nil
}
