// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/q191201771/lal/pkg/base"

	"github.com/q191201771/naza/pkg/nazalog"
)

// RTMP拉流客户端，从远端服务器拉取RTMP流，存储为本地FLV文件
//
// 另外，作为一个RTMP拉流压测工具，已经支持：
// 1. 对一路流拉取n份
// 2. 拉取n路流
//
// Usage of ./bin/pullrtmp:
//   -i string
//     	specify pull rtmp url
//   -n int
//     	num of pull connection (default 1)
//   -o string
//     	specify output flv file
// Example:
//   ./bin/pullrtmp -i rtmp://127.0.0.1:1935/live/test -o out.flv
//   ./bin/pullrtmp -i rtmp://127.0.0.1:1935/live/test -n 1000
//   ./bin/pullrtmp -i rtmp://127.0.0.1:1935/live/test_{i} -n 1000

var aliveSessionCount int32

func main() {
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.AssertBehavior = nazalog.AssertFatal
	})
	defer nazalog.Sync()
	base.LogoutStartInfo()

	urlTmpl, fileNameTmpl, num := parseFlag()
	nazalog.Infof("parse flag succ. urlTmpl=%s, fileNameTmpl=%s, num=%d", urlTmpl, fileNameTmpl, num)
	urls, filenames := collect(urlTmpl, fileNameTmpl, num)

	go func() {
		for {
			nazalog.Debugf("alive session:%d", atomic.LoadInt32(&aliveSessionCount))
			time.Sleep(1 * time.Second)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(index int) {
			pull(urls[index], filenames[index])
			wg.Done()
			atomic.AddInt32(&aliveSessionCount, -1)
		}(i)
	}
	wg.Wait()
	time.Sleep(1 * time.Second)
	nazalog.Info("< main.")
}

func pull(url string, filename string) { _ = "STUB: not implemented"; return }

// 临时测试一下主动关闭client session
//go func() {
//	time.Sleep(5 * time.Second)
//	err := session.Dispose()
//	nazalog.Debugf("< session Dispose. err=%+v", err)
//}()

func collect(urlTmpl string, fileNameTmpl string, num int) (urls []string, filenames []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFlag() (urlTmpl string, fileNameTmpl string, num int) {
	_ = "STUB: not implemented"
	return "", "", 0
}
