// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"github.com/q191201771/lal/pkg/base"
)

func SimpleAuthCalcSecret(key string, streamName string) string {
	_ = "STUB: not implemented"
	return ""
}

// ---------------------------------------------------------------------------------------------------------------------

// TODO(chef): [refactor] 结合 NotifyHandler 整理

const secretName = "lal_secret"

type SimpleAuthCtx struct {
	config SimpleAuthConfig
}

func NewSimpleAuthCtx(config SimpleAuthConfig) *SimpleAuthCtx {
	_ = "STUB: not implemented"
	return nil
}

func (s *SimpleAuthCtx) OnPubStart(info base.PubStartInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SimpleAuthCtx) OnSubStart(info base.SubStartInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SimpleAuthCtx) OnHls(streamName string, urlParam string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SimpleAuthCtx) check(streamName string, urlParam string) error {
	_ = "STUB: not implemented"
	return nil
}

// 注意，只有DangerousLalSecret配置了值，才验证参数是否和DangerousLalSecret相等
