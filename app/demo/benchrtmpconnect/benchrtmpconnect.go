// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/rtmp"
	"github.com/q191201771/naza/pkg/nazalog"
)

func main() {
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.AssertBehavior = nazalog.AssertFatal
		option.Level = nazalog.LevelLogNothing
	})
	defer nazalog.Sync()
	base.LogoutStartInfo()

	urlTmpl, num := parseFlag()
	urls := collect(urlTmpl, num)

	var mu sync.Mutex
	var succCosts []int64
	var failCosts []int64
	var wg sync.WaitGroup
	wg.Add(len(urls))

	go func() {
		for {
			mu.Lock()
			succ := len(succCosts)
			fail := len(failCosts)
			mu.Unlock()
			_, _ = fmt.Fprintf(os.Stderr, "task(num): total=%d, succ=%d, fail=%d\n", len(urls), succ, fail)
			time.Sleep(1 * time.Second)
		}
	}()

	totalB := time.Now()
	for _, url := range urls {
		go func(u string) {
			pullSession := rtmp.NewPullSession(func(option *rtmp.PullSessionOption) {
				option.PullTimeoutMs = 30000
				option.ReadAvTimeoutMs = 30000
				option.HandshakeComplexFlag = false
			})
			b := time.Now()
			err := pullSession.Start(u)
			e := time.Now()
			cost := e.Sub(b).Milliseconds()
			// 耗时不够1毫秒，我们将值取整到1毫秒，并打印更精确的实际耗时
			if cost == 0 {
				_, _ = fmt.Fprintf(os.Stderr, "round to 1 ms but actual is %s\n", e.Sub(b).String())
				cost = 1
			}

			mu.Lock()
			if err == nil {
				succCosts = append(succCosts, cost)
			} else {
				failCosts = append(failCosts, cost)
			}
			mu.Unlock()
			wg.Done()
		}(url)
	}
	wg.Wait()
	totalE := time.Now()
	totalCost := totalE.Sub(totalB).Milliseconds()
	if totalCost == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "round to 1 ms but actual is %s\n", totalE.Sub(totalB).String())
		totalCost = 1
	}
	min, max, avg := analyse(succCosts)
	_, _ = fmt.Fprintf(os.Stderr, "task(num): total=%d, succ=%d, fail=%d\n", len(urls), len(succCosts), len(failCosts))
	_, _ = fmt.Fprintf(os.Stderr, " cost(ms): total=%d, avg=%d, min=%d, max=%d\n", totalCost, avg, min, max)
}

func analyse(costs []int64) (min, max, avg int64) { _ = "STUB: not implemented"; return 0, 0, 0 }

func collect(urlTmpl string, num int) (urls []string) { _ = "STUB: not implemented"; return nil }

func parseFlag() (urlTmpl string, num int) { _ = "STUB: not implemented"; return "", 0 }
