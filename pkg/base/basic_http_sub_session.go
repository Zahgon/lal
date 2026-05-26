// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package base

import (
	"net"

	"github.com/q191201771/naza/pkg/connection"
)

type BasicHttpSubSession struct {
	BasicHttpSubSessionOption

	suffix      string
	conn        connection.Connection
	sessionStat BasicSessionStat
}

type BasicHttpSubSessionOption struct {
	Conn          net.Conn
	ConnModOption connection.ModOption
	SessionType   SessionType
	UrlCtx        UrlContext
	IsWebSocket   bool
	WebSocketKey  string
}

func NewBasicHttpSubSession(option BasicHttpSubSessionOption) *BasicHttpSubSession {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
// IServerSessionLifecycle interface
// ---------------------------------------------------------------------------------------------------------------------

func (session *BasicHttpSubSession) RunLoop() error { _ = "STUB: not implemented"; return nil }

func (session *BasicHttpSubSession) Dispose() error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------

func (session *BasicHttpSubSession) WriteHttpResponseHeader(b []byte) {
	_ = "STUB: not implemented"
	return
}

func (session *BasicHttpSubSession) Write(b []byte) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------
// IObject interface
// ---------------------------------------------------------------------------------------------------------------------

func (session *BasicHttpSubSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionUrlContext interface
// ---------------------------------------------------------------------------------------------------------------------

func (session *BasicHttpSubSession) Url() string { _ = "STUB: not implemented"; return "" }

func (session *BasicHttpSubSession) AppName() string { _ = "STUB: not implemented"; return "" }

func (session *BasicHttpSubSession) StreamName() string { _ = "STUB: not implemented"; return "" }

func (session *BasicHttpSubSession) RawQuery() string { _ = "STUB: not implemented"; return "" }

// ----- ISessionStat --------------------------------------------------------------------------------------------------

func (session *BasicHttpSubSession) GetStat() StatSession {
	_ = "STUB: not implemented"
	return *new(StatSession)
}

func (session *BasicHttpSubSession) UpdateStat(intervalSec uint32) {
	_ = "STUB: not implemented"
	return
}

func (session *BasicHttpSubSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

// ---------------------------------------------------------------------------------------------------------------------

func (session *BasicHttpSubSession) write(b []byte) {
	_ = "STUB: not implemented"
	// TODO(chef) handle write error
	return
}
