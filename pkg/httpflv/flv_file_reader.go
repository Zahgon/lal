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

type FlvFileReader struct {
	fp               *os.File
	hasReadFlvHeader bool
}

func (ffr *FlvFileReader) Open(filename string) (err error) { _ = "STUB: not implemented"; return nil }

func (ffr *FlvFileReader) ReadFlvHeader() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ffr *FlvFileReader) ReadTag() (Tag, error) {
	_ = "STUB: not implemented"
	// lazy read flv header
	return *new(Tag), nil
}

func (ffr *FlvFileReader) Dispose() { _ = "STUB: not implemented"; return }
