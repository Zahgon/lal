// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package httpflv

import (
	"os"
)

// TODO chef: 结构体重命名为FileWriter，文件名重命名为file_writer.go。所有写流文件的（flv,hls,ts）统一重构

type FlvFileWriter struct {
	fp *os.File
}

func (ffw *FlvFileWriter) Open(filename string) (err error) { _ = "STUB: not implemented"; return nil }

func (ffw *FlvFileWriter) WriteRaw(b []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (ffw *FlvFileWriter) WriteFlvHeader() (err error) { _ = "STUB: not implemented"; return nil }

func (ffw *FlvFileWriter) WriteTag(tag Tag) (err error) { _ = "STUB: not implemented"; return nil }

func (ffw *FlvFileWriter) Dispose() error { _ = "STUB: not implemented"; return nil }

func (ffw *FlvFileWriter) Name() string { _ = "STUB: not implemented"; return "" }
