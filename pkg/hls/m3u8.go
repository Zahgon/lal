// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package hls

// writeM3u8File
//
// @param content     需写入文件的内容
// @param filename    m3u8文件名
// @param filenameBak m3u8临时文件名
func writeM3u8File(content []byte, filename string, filenameBak string) error {
	_ = "STUB: not implemented"
	return nil
}

// updateTargetDurationInM3u8 如果当前duration比原m3u8文件的`EXT-X-TARGETDURATION`大，则更新`EXT-X-TARGETDURATION`的值
//
// @param content      原m3u8文件的内容
// @param currDuration 当前duration
//
// @return 处理后的m3u8文件内容
func updateTargetDurationInM3u8(content []byte, currDuration int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CalcM3u8Duration
//
// @param content 传入m3u8文件内容
//
// @return durationSec m3u8中所有ts的时间总和。注意，使用的是m3u8文件中描述的ts时间，而不是读取ts文件中实际音视频数据的时间。
func CalcM3u8Duration(content []byte) (durationSec float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
