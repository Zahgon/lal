// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package hls

import (
	"sync"

	"github.com/q191201771/naza/pkg/filesystemlayer"
)

var (
	fslCtx  filesystemlayer.IFileSystemLayer
	setOnce sync.Once
)

func SetUseMemoryAsDiskFlag(flag bool) { _ = "STUB: not implemented"; return }

func ReadFile(filename string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func RemoveAll(path string) error { _ = "STUB: not implemented"; return nil }

func init() {
	fslCtx = filesystemlayer.FslFactory(filesystemlayer.FslTypeDisk)
}
