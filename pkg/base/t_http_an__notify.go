// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package base

// 文档见： https://pengrl.com/lal/#/HTTPNotify

// EventCommonInfo 所有事件共有的字段
type EventCommonInfo struct {
	ServerId string `json:"server_id"`
}

// SessionEventCommonInfo session相关的事件的共有的字段
type SessionEventCommonInfo struct {
	EventCommonInfo

	SessionId  string `json:"session_id"`
	Protocol   string `json:"protocol"`
	BaseType   string `json:"base_type"`
	RemoteAddr string `json:"remote_addr"`

	Url        string `json:"url"`
	AppName    string `json:"app_name"`
	StreamName string `json:"stream_name"`
	UrlParam   string `json:"url_param"`

	HasInSession  bool `json:"has_in_session"`
	HasOutSession bool `json:"has_out_session"`

	// TODO(chef): [opt] 这两个字段，实际需求出发点是有业务方需要在stop事件做流量统计，但是现在的实现为所有session事件都添加了，是否合适 202208
	ReadBytesSum  uint64 `json:"read_bytes_sum"`
	WroteBytesSum uint64 `json:"wrote_bytes_sum"`
}

type UpdateInfo struct {
	EventCommonInfo

	Groups []StatGroup `json:"groups"`
}

type PubStartInfo struct {
	SessionEventCommonInfo
}

type PubStopInfo struct {
	SessionEventCommonInfo
}

type SubStartInfo struct {
	SessionEventCommonInfo
}

type SubStopInfo struct {
	SessionEventCommonInfo
}

type PullStartInfo struct {
	SessionEventCommonInfo
}

type PullStopInfo struct {
	SessionEventCommonInfo
}

type RtmpConnectInfo struct {
	EventCommonInfo

	SessionId  string `json:"session_id"`
	RemoteAddr string `json:"remote_addr"`
	App        string `json:"app"`
	FlashVer   string `json:"flashVer"`
	TcUrl      string `json:"tcUrl"`
}

type HlsMakeTsInfo struct {
	EventCommonInfo

	Event          string  `json:"event"`
	StreamName     string  `json:"stream_name"`
	Cwd            string  `json:"cwd"`
	TsFile         string  `json:"ts_file"`
	LiveM3u8File   string  `json:"live_m3u8_file"`
	RecordM3u8File string  `json:"record_m3u8_file"`
	Id             int     `json:"id"`
	Duration       float64 `json:"duration"`
}

// ---------------------------------------------------------------------------------------------------------------------

func Session2PubStartInfo(session ISession) PubStartInfo {
	_ = "STUB: not implemented"
	return *new(PubStartInfo)
}

func Session2PubStopInfo(session ISession) PubStopInfo {
	_ = "STUB: not implemented"
	return *new(PubStopInfo)
}

func Session2SubStartInfo(session ISession) SubStartInfo {
	_ = "STUB: not implemented"
	return *new(SubStartInfo)
}

func Session2SubStopInfo(session ISession) SubStopInfo {
	_ = "STUB: not implemented"
	return *new(SubStopInfo)
}

func Session2PullStartInfo(session ISession) PullStartInfo {
	_ = "STUB: not implemented"
	return *new(PullStartInfo)
}

func Session2PullStopInfo(session ISession) PullStopInfo {
	_ = "STUB: not implemented"
	return *new(PullStopInfo)
}

// ---------------------------------------------------------------------------------------------------------------------

func session2EventCommonInfo(session ISession) SessionEventCommonInfo {
	_ = "STUB: not implemented"
	return *new(SessionEventCommonInfo)
}
