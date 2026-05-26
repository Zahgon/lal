// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package hls

import (
	"github.com/q191201771/lal/pkg/base"
)

// 聚合以下功能：
// - 落盘策略： 生成HLS（m3u8文件+ts文件）时，文件命名规则，以及文件存放规则
// - 路由策略： HTTP请求HLS时，request URI和文件路径的映射规则

type RequestInfo struct {
	StreamName       string // uri结合策略
	FileNameWithPath string // uri结合策略, 从磁盘打开文件时使用
}

type IPathStrategy interface {
	IPathRequestStrategy
	IPathWriteStrategy
}

// IPathRequestStrategy
//
// 路由策略
// 接到HTTP请求时，对应文件路径的映射逻辑
type IPathRequestStrategy interface {
	// GetRequestInfo
	//
	// 解析HTTP请求，得到流名称、文件所在路径
	//
	GetRequestInfo(urlCtx base.UrlContext, rootOutPath string) RequestInfo
}

// IPathWriteStrategy 落盘策略
type IPathWriteStrategy interface {
	// GetMuxerOutPath 获取单个流对应的文件根路径
	GetMuxerOutPath(rootOutPath string, streamName string) string

	// GetLiveM3u8FileName 获取单个流对应的m3u8文件路径
	//
	// @param outPath: func GetMuxerOutPath的结果
	GetLiveM3u8FileName(outPath string, streamName string) string

	// GetRecordM3u8FileName 获取单个流对应的record类型的m3u8文件路径
	//
	// live m3u8和record m3u8的区别：
	// live记录的是当前最近的可播放内容，record记录的是从流开始时的可播放内容
	//
	// @param outPath: func GetMuxerOutPath的结果
	GetRecordM3u8FileName(outPath string, streamName string) string

	// GetTsFileNameWithPath 获取单个流对应的ts文件路径
	//
	// @param outPath: func GetMuxerOutPath的结果
	GetTsFileNameWithPath(outPath string, fileName string) string

	// GetTsFileName ts文件名的生成策略
	GetTsFileName(streamName string, index int, timestamp int) string
}

// ---------------------------------------------------------------------------------------------------------------------

const (
	playlistM3u8FileName = "playlist.m3u8"
	recordM3u8FileName   = "record.m3u8"
)

// DefaultPathStrategy 默认的路由，落盘策略
//
// 每个流在<rootPath>下以流名称生成一个子目录，目录下包含:
//
// - playlist.m3u8              实时的HLS文件，定期刷新，写入当前最新的TS文件列表，淘汰过期的TS文件列表
// - record.m3u8                录制回放的HLS文件，包含了从流开始至今的所有TS文件
// - test110-1620540712084-0.ts TS分片文件，命名格式为{liveid}-{timestamp}-{index}.ts
// - test110-1620540716095-1.ts
// - ...                        一系列的TS文件
//
// 假设
// 流名称="test110"
// rootPath="/tmp/lal/hls/"
//
// 则
// http://127.0.0.1:8080/hls/test110/playlist.m3u8              -> /tmp/lal/hls/test110/playlist.m3u8
// http://127.0.0.1:8080/hls/test110/record.m3u8                -> /tmp/lal/hls/test110/record.m3u8
// http://127.0.0.1:8080/hls/test110/test110-1620540712084-0.ts -> /tmp/lal/hls/test110/test110-1620540712084-0.ts
//
// http://127.0.0.1:8080/hls/test110.m3u8                       -> /tmp/lal/hls/test110/playlist.m3u8
// http://127.0.0.1:8080/hls/test110-1620540712084-0.ts         -> /tmp/lal/hls/test110/test110-1620540712084-0.ts
// 最下面这两个做了特殊映射
type DefaultPathStrategy struct {
}

// GetRequestInfo
//
// RequestURI example:
// uri                                    -> FileName                  StreamName FileType FileNameWithPath
// /hls/test110.m3u8                      -> test110.m3u8              test110    m3u8     {rootOutPath}/test110/playlist.m3u8
// /hls/test110/playlist.m3u8             -> playlist.m3u8             test110    m3u8     {rootOutPath}/test110/playlist.m3u8
// /hls/test110/record.m3u8               -> record.m3u8               test110    m3u8     {rootOutPath}/test110/record.m3u8
// /hls/test110/test110-1620540712084-.ts -> test110-1620540712084-.ts test110    ts       {rootOutPath/test110/test110-1620540712084-.ts
// /hls/test110-1620540712084-.ts         -> test110-1620540712084-.ts test110    ts       {rootOutPath/test110/test110-1620540712084-.ts
func (dps *DefaultPathStrategy) GetRequestInfo(urlCtx base.UrlContext, rootOutPath string) (ri RequestInfo) {
	_ = "STUB: not implemented"
	return *new(RequestInfo)
}

// GetMuxerOutPath <rootOutPath>/<streamName>
func (*DefaultPathStrategy) GetMuxerOutPath(rootOutPath string, streamName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (*DefaultPathStrategy) GetLiveM3u8FileName(outPath string, streamName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (*DefaultPathStrategy) GetRecordM3u8FileName(outPath string, streamName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (*DefaultPathStrategy) GetTsFileNameWithPath(outPath string, fileName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (*DefaultPathStrategy) GetTsFileName(streamName string, index int, timestamp int) string {
	_ = "STUB: not implemented"
	return ""
}

func (*DefaultPathStrategy) getStreamNameFromTsFileName(fileName string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetTsFileName 格式固定为%s-%d-%d.ts streamName取%s部分
