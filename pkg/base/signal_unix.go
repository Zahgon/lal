// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

//go:build linux || darwin || netbsd || freebsd || openbsd || dragonfly
// +build linux darwin netbsd freebsd openbsd dragonfly

package base

// RunSignalHandler 监听SIGUSR1和SIGUSR2信号并回调
//
// TODO(chef): refactor 函数名应与SIGUSR1挂钩
func RunSignalHandler(cb func()) { _ = "STUB: not implemented"; return }
