// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtprtcp

type Rr struct {
	senderSsrc  uint32
	mediaSsrc   uint32
	fraction    uint8
	lost        uint32
	cycles      uint16
	extendedSeq uint32
	jitter      uint32
	lsr         uint32
	dlsr        uint32 // default 0
}

func (r *Rr) Pack() []byte { _ = "STUB: not implemented"; return nil }

// TODO chef: lost的表示是否正确
