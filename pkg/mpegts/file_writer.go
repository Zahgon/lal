// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package mpegts

import (
	"os"
)

type FileWriter struct {
	fp *os.File
}

func (fw *FileWriter) Create(filename string) (err error) { _ = "STUB: not implemented"; return nil }

func (fw *FileWriter) Write(b []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (fw *FileWriter) Dispose() error { _ = "STUB: not implemented"; return nil }

func (fw *FileWriter) Name() string { _ = "STUB: not implemented"; return "" }
