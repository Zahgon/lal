// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"github.com/q191201771/naza/pkg/nazalog"

	"github.com/q191201771/lal/pkg/logic"
)

func main() {
	defer nazalog.Sync()

	confFilename := parseFlag()
	lals := logic.NewLalServer(func(option *logic.Option) {
		option.ConfFilename = confFilename
	})
	err := lals.RunLoop()
	nazalog.Infof("lal server loop done. err=%+v", err)
}

func parseFlag() string { _ = "STUB: not implemented"; return "" }
