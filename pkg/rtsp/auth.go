// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

// TODO chef: 考虑部分内容移入naza中

const (
	AuthTypeDigest = "Digest"
	AuthTypeBasic  = "Basic"
	AuthAlgorithm  = "MD5"
)

type Auth struct {
	Username string
	Password string

	Typ       string
	Realm     string
	Nonce     string
	Algorithm string
	Uri       string
	Response  string
	Opaque    string // 暂时没用
	Stale     string // 暂时没用
}

// ParseAuthorization 解析字段，server side使用
func (a *Auth) ParseAuthorization(authStr string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// FeedWwwAuthenticate 使用第一轮回复，client side使用
func (a *Auth) FeedWwwAuthenticate(auths []string, username, password string) {
	_ = "STUB: not implemented"
	return
}

//目前只处理第一个

// MakeAuthorization 生成第二轮请求，client side使用
//
// 如果没有调用`FeedWwwAuthenticate`初始化过，则直接返回空字符串
func (a *Auth) MakeAuthorization(method, uri string) string { _ = "STUB: not implemented"; return "" }

// MakeAuthenticate 生成第一轮的回复，server side使用
func (a *Auth) MakeAuthenticate(method string) string { _ = "STUB: not implemented"; return "" }

// CheckAuthorization 验证第二轮请求，server side使用
func (a *Auth) CheckAuthorization(method, username, password string) bool {
	_ = "STUB: not implemented"
	return false
}

// The "response" field is computed as:
// md5(md5(<username>:<realm>:<password>):<nonce>:md5(<cmd>:<url>))

// ---------------------------------------------------------------------------------------------------------------------

func (a *Auth) getV(s string, pre string) string { _ = "STUB: not implemented"; return "" }

func (a *Auth) nonce() string { _ = "STUB: not implemented"; return "" }
