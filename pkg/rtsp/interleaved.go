// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

import (
	"bufio"
)

// rfc2326 10.12 Embedded (Interleaved) Binary Data

func readInterleaved(r *bufio.Reader) (isInterleaved bool, packet []byte, channel uint8, err error) {
	_ = "STUB: not implemented"
	return false, nil, 0, nil
}

// TODO chef: 这里为了安全性，应该检查大小

func packInterleaved(channel int, rtpPacket []byte) []byte { _ = "STUB: not implemented"; return nil }
