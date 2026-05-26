// Copyright 2023, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package base

import "crypto/tls"

// TODO(chef): 移入naza中 2302
// TODO(chef): 将lal中使用tls.Config的地方都聚合到这里来 2302

func DefaultTlsConfigClient() *tls.Config { _ = "STUB: not implemented"; return nil }

func DefaultTlsConfigServer(certFile, keyFile string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
