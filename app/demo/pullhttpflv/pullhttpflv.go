// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"os"

	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/lal/pkg/rtmp"
	"github.com/q191201771/naza/pkg/nazalog"
)

// 拉取http-flv的流并存储为flv文件

func main() {
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.AssertBehavior = nazalog.AssertFatal
	})
	defer nazalog.Sync()
	base.LogoutStartInfo()

	url, flvname := parseFlag()
	flvfile, err := os.Create(flvname)
	if err != nil {
		nazalog.Errorf("create flv file failed, err=%+v", err)
		return
	}

	defer flvfile.Close()

	session := httpflv.NewPullSession().WithOnReadFlvTag(func(tag httpflv.Tag) {
		if tag.Header.Type == httpflv.TagTypeMetadata {
			// TODO(chef): httpflv.PullSession支持返回flv header，可供业务方选择使用 202210
			// 根据metadata填写flv头
			opa, err := rtmp.ParseMetadata(tag.Payload())
			if err != nil {
				nazalog.Errorf("ParseMetadata failed, err=%+v", err)
				return
			}

			b := make([]byte, 13)
			var flags uint8

			audiocodecid := opa.Find("audiocodecid")
			videocodecid := opa.Find("videocodecid")
			if audiocodecid != 0 {
				flags |= 0x04
			}

			if videocodecid != 0 {
				flags |= 0x01
			}

			writeFlvHeader(b, flags)
			flvfile.Write(b)
		}

		nazalog.Infof("tag Type:%d, tag Size:%d", tag.Header.Type, tag.Header.DataSize)

		flvfile.Write(tag.Raw)
	})

	err = session.Start(url)
	nazalog.Assert(nil, err)
	err = <-session.WaitChan()
	nazalog.Assert(nil, err)
}

func parseFlag() (url, flvfile string) { _ = "STUB: not implemented"; return "", "" }

func writeFlvHeader(b []byte, flags uint8) {
	_ = "STUB: not implemented"

	// 'FLV', version 1
	return
}

// DataOffset: UI32 Offset in bytes from start of file to start of body (that is, size of header)
// The DataOffset field usually has a value of 9 for FLV version 1.

// PreviousTagSize0: UI32 Always 0
