// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package hls

import (
	"net/http"
	"time"

	"github.com/q191201771/naza/pkg/nazaatomic"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/naza/pkg/connection"
)

type SubSession struct {
	LastRequestTime time.Time
	urlCtx          base.UrlContext
	hlsUrlPattern   string
	appName         string
	timeout         time.Duration
	sessionIdHash   string // Because session.UniqueKey() too easy to guess so that we need to hash it with a key to prevent client guess session id

	stat     base.StatSession
	prevStat connection.Stat
	currStat connection.StatAtomic

	disposedFlag nazaatomic.Bool
}

func (s *SubSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

func NewSubSession(req *http.Request, urlCtx base.UrlContext, hlsUrlPattern, sessionHashKey string, timeout time.Duration) *SubSession {
	_ = "STUB: not implemented"
	return nil
}

// stat

// TODO(chef): [refactor] 也许后续可以弄短点，比如前8位或16位 202211

func (s *SubSession) Url() string { _ = "STUB: not implemented"; return "" }

func (s *SubSession) AppName() string { _ = "STUB: not implemented"; return "" }

func (s *SubSession) StreamName() string { _ = "STUB: not implemented"; return "" }

func (s *SubSession) RawQuery() string { _ = "STUB: not implemented"; return "" }

func (s *SubSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

func (s *SubSession) AddWroteBytesSum(wbs uint64) { _ = "STUB: not implemented"; return }

func (s *SubSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *new(base.StatSession)
}

func (s *SubSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (s *SubSession) IsExpired() bool { _ = "STUB: not implemented"; return false }

func (s *SubSession) IsDisposed() bool { _ = "STUB: not implemented"; return false }

func (s *SubSession) KeepAlive() { _ = "STUB: not implemented"; return }

func (s *SubSession) Dispose() { _ = "STUB: not implemented"; return }

func GetAppNameFromUrlCtx(urlCtx base.UrlContext, hlsUrlPattern string) string {
	_ = "STUB: not implemented"
	return ""
}

func GetStreamNameFromUrlCtx(urlCtx base.UrlContext) string { _ = "STUB: not implemented"; return "" }
