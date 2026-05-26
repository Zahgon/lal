// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package hls

import (
	"github.com/q191201771/naza/pkg/filesystemlayer"
)

type Fragment struct {
	fp filesystemlayer.IFile
}

func (f *Fragment) OpenFile(filename string) (err error) { _ = "STUB: not implemented"; return nil }

func (f *Fragment) WriteFile(b []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (f *Fragment) CloseFile() error { _ = "STUB: not implemented"; return nil }
