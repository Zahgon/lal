// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package datamanager

import (
	"sync"
)

type DataManagerMemory struct {
	serverTimeoutSec    int
	mutex               sync.Mutex
	serverId2pubStreams map[string]map[string]struct{}
	serverId2AliveTs    map[string]int64
}

func NewDataManagerMemory(serverTimeoutSec int) *DataManagerMemory {
	_ = "STUB: not implemented"
	return nil
}

// TODO chef: release goroutine

// 清除长时间没有update报活的节点

// 定时打印数据日志

func (d *DataManagerMemory) AddPub(streamName, serverId string) { _ = "STUB: not implemented"; return }

func (d *DataManagerMemory) DelPub(streamName, serverId string) { _ = "STUB: not implemented"; return }

func (d *DataManagerMemory) QueryPub(streamName string) (serverId string, exist bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (d *DataManagerMemory) UpdatePub(serverId string, streamNameList []string) {
	_ = "STUB: not implemented"
	// 3. server超时，去掉所有上面所有的pub
	return
}

// 更新serverId对应的stream列表

// only for log

func (d *DataManagerMemory) queryPub(streamName string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (d *DataManagerMemory) markAlive(serverId string) { _ = "STUB: not implemented"; return }
