// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package sdp

func ParseAsc(a *AFmtPBase) ([]byte, error) {
	_ = "STUB: not implemented"
	// AAC的这个地方除了97，还遇到过104的，暂时不要这个判断了
	//
	//	if a.Format != base.RtpPacketTypeAac {
	//		return nil, nazaerrors.Wrap(base.ErrSdp)
	//	}
	return nil, nil
}

func ParseVpsSpsPps(a *AFmtPBase) (vps, sps, pps []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// ParseSpsPps
//
// 解析AVC/H264的sps，pps
// 例子见单元测试
func ParseSpsPps(a *AFmtPBase) (sps, pps []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
