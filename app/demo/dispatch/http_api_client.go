// Copyright 2024, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

func kickSession(serverId, streamName, sessionId string) { _ = "STUB: not implemented"; return }

func addIpBlacklist(serverId, ip string, durationSec int) { _ = "STUB: not implemented"; return }

func startRelayPull(reqId, reqApiAddr, pubRtmpAddr, appName, streamName string) {
	_ = "STUB: not implemented"
	// TODO(chef): 还没有测试新的接口start_relay_pull，只是保证可以编译通过
	return
}

//b.Protocol = base.ProtocolRtmp
//b.Addr = pubServer.RtmpAddr
//b.AppName = info.AppName
//b.StreamName = info.StreamName
//b.UrlParam = config.PullSecretParam
