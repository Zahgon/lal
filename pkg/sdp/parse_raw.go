// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package sdp

type RawContext struct {
	MediaDescList []MediaDesc
}

type MediaDesc struct {
	M         M
	ARtpMap   ARtpMap
	AFmtPBase *AFmtPBase
	AControl  AControl
}

type M struct {
	Media string
	PT    int // 暂时只支持m只有一个pt值的情况
}

type ARtpMap struct {
	PayloadType        int
	EncodingName       string
	ClockRate          int
	EncodingParameters string
}

type AFmtPBase struct {
	Format     int               // same as PayloadType
	Parameters map[string]string // name -> value
}

type AControl struct {
	Value string
}

// ParseSdp2RawContext 例子见单元测试
func ParseSdp2RawContext(b []byte) (RawContext, error) {
	_ = "STUB: not implemented"
	return *new(RawContext), nil
}

// TestCase13，再尝试抢救一下

// TODO(chef): 如果换行的数据刚好是`m=`或`a=`开头呢？

func ParseM(s string) (ret M, err error) { _ = "STUB: not implemented"; return *new(M), nil }

// ParseARtpMap 例子见单元测试
func ParseARtpMap(s string) (ret ARtpMap, err error) {
	_ = "STUB: not implemented"
	// rfc 3640 3.3.1.  General
	// rfc 3640 3.3.6.  High Bit-rate AAC
	//
	// a=rtpmap:<payload type> <encoding name>/<clock rate>[/<encoding parameters>]
	//
	return *new(ARtpMap), nil
}

// ParseAFmtPBase 例子见单元测试
func ParseAFmtPBase(s string) (ret AFmtPBase, err error) {
	_ = "STUB: not implemented"
	// rfc 3640 4.4.1.  The a=fmtp Keyword
	//
	// a=fmtp:<format> <parameter name>=<value>[; <parameter name>=<value>]
	//
	return *new(AFmtPBase), nil
}

// 见TestCase11

// 见TestCase12

func ParseAControl(s string) (ret AControl, err error) {
	_ = "STUB: not implemented"
	return *new(AControl), nil
}

// ---------------------------------------------------------------------------------------------------------------------

func parseSdp2RawContext(lines []string) (RawContext, error) {
	_ = "STUB: not implemented"
	return *new(RawContext), nil
}
