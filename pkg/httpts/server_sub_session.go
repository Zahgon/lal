// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package httpts

import (
	"net"

	"github.com/q191201771/lal/pkg/base"
)

type SubSession struct {
	core               *base.BasicHttpSubSession
	IsFresh            bool
	ShouldWaitBoundary bool
}

func NewSubSession(conn net.Conn, urlCtx base.UrlContext, isWebSocket bool, websocketKey string) *SubSession {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
// IServerSessionLifecycle interface
// ---------------------------------------------------------------------------------------------------------------------

func (session *SubSession) RunLoop() error { _ = "STUB: not implemented"; return nil }

func (session *SubSession) Dispose() error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------

func (session *SubSession) WriteHttpResponseHeader() { _ = "STUB: not implemented"; return }

func (session *SubSession) Write(b []byte) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------
// IObject interface
// ---------------------------------------------------------------------------------------------------------------------

func (session *SubSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionUrlContext interface
// ---------------------------------------------------------------------------------------------------------------------

func (session *SubSession) Url() string { _ = "STUB: not implemented"; return "" }

func (session *SubSession) AppName() string { _ = "STUB: not implemented"; return "" }

func (session *SubSession) StreamName() string { _ = "STUB: not implemented"; return "" }

func (session *SubSession) RawQuery() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------
// ISessionStat interface
// ---------------------------------------------------------------------------------------------------------------------

func (session *SubSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

func (session *SubSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *new(base.StatSession)
}

func (session *SubSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}
