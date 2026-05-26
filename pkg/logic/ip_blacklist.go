// Copyright 2024, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"sync"
)

type IpBlacklist struct {
	mu  sync.Mutex
	ips map[string]int64 // TODO(chef): 优化性能 202405
}

func (l *IpBlacklist) Add(ip string, durationSec int) { _ = "STUB: not implemented"; return }

func (l *IpBlacklist) Has(ip string) bool { _ = "STUB: not implemented"; return false }

func (l *IpBlacklist) eraseStale() { _ = "STUB: not implemented"; return }
