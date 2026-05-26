// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

// Package base 提供被其他多个package依赖的基础内容，自身不依赖任何package
package base

import (
	"time"
)

// TODO chef: 考虑部分内容放入关联的协议package的子package中

var startTime string

var readableTimeLayout = "2006-01-02 15:04:05.999 Z0700 MST"

// ReadableNowTime 当前时间，可读字符串形式
func ReadableNowTime() string { _ = "STUB: not implemented"; return "" }

func ParseReadableTime(t string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func GetWd() string { _ = "STUB: not implemented"; return "" }

func LogoutStartInfo() { _ = "STUB: not implemented"; return }

func WrapReadConfigFile(theConfigFile string, defaultConfigFiles []string, hookBeforeExit func()) []byte {
	_ = "STUB: not implemented"
	// TODO(chef): 统一本函数内的Log和stderr输出 202405
	return nil
}

// 如果没有指定配置文件，则尝试从默认路径找配置文件

// 如果默认路径也没有配置文件，则退出

// 读取配置文件

func init() {
	startTime = ReadableNowTime()
}
