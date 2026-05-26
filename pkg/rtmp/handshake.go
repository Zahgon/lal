// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtmp

import (
	"io"
)

// https://pengrl.com/p/20027

const version = uint8(3)

const (
	c0c1Len   = 1537
	c2Len     = 1536
	s0s1Len   = 1537
	s1Len     = 1536
	s2Len     = 1536
	s0s1s2Len = 3073
)

const (
	clientPartKeyLen = 30
	clientFullKeyLen = 62
	serverPartKeyLen = 36
	serverFullKeyLen = 68
	keyLen           = 32
)

var (
	clientVersionMockFromFfmpeg = []byte{9, 0, 124, 2} // emulated Flash client version - 9.0.124.2 on Linux
	clientVersion               = []byte{0x0C, 0x00, 0x0D, 0x0E}
	serverVersion               = []byte{0x0D, 0x0E, 0x0A, 0x0D}
)

// 30+32
var clientKey = []byte{
	'G', 'e', 'n', 'u', 'i', 'n', 'e', ' ', 'A', 'd', 'o', 'b', 'e', ' ',
	'F', 'l', 'a', 's', 'h', ' ', 'P', 'l', 'a', 'y', 'e', 'r', ' ',
	'0', '0', '1',

	0xF0, 0xEE, 0xC2, 0x4A, 0x80, 0x68, 0xBE, 0xE8, 0x2E, 0x00, 0xD0, 0xD1,
	0x02, 0x9E, 0x7E, 0x57, 0x6E, 0xEC, 0x5D, 0x2D, 0x29, 0x80, 0x6F, 0xAB,
	0x93, 0xB8, 0xE6, 0x36, 0xCF, 0xEB, 0x31, 0xAE,
}

// 36+32
var serverKey = []byte{
	'G', 'e', 'n', 'u', 'i', 'n', 'e', ' ', 'A', 'd', 'o', 'b', 'e', ' ',
	'F', 'l', 'a', 's', 'h', ' ', 'M', 'e', 'd', 'i', 'a', ' ',
	'S', 'e', 'r', 'v', 'e', 'r', ' ',
	'0', '0', '1',

	0xF0, 0xEE, 0xC2, 0x4A, 0x80, 0x68, 0xBE, 0xE8, 0x2E, 0x00, 0xD0, 0xD1,
	0x02, 0x9E, 0x7E, 0x57, 0x6E, 0xEC, 0x5D, 0x2D, 0x29, 0x80, 0x6F, 0xAB,
	0x93, 0xB8, 0xE6, 0x36, 0xCF, 0xEB, 0x31, 0xAE,
}

var random1528Buf []byte

type IHandshakeClient interface {
	WriteC0C1(writer io.Writer) error
	ReadS0S1(reader io.Reader) error
	WriteC2(writer io.Writer) error
	ReadS2(reader io.Reader) error
}

type HandshakeClientSimple struct {
	buf []byte
}

type HandshakeClientComplex struct {
	buf []byte
}

type HandshakeServer struct {
	isSimpleMode bool
	s0s1s2       []byte
}

func (c *HandshakeClientSimple) WriteC0C1(writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// 4字节模式串保持为0，标识是简单模式

func (c *HandshakeClientSimple) ReadS0S1(reader io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandshakeClientSimple) WriteC2(writer io.Writer) error {
	_ = "STUB: not implemented"
	// use s1 as c2
	return nil
}

func (c *HandshakeClientSimple) ReadS2(reader io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandshakeClientComplex) WriteC0C1(writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// mock ffmpeg

func (c *HandshakeClientComplex) ReadS0S1(reader io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

// simple mode

// use s1 as c2

// complex mode

func (c *HandshakeClientComplex) WriteC2(writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandshakeClientComplex) ReadS2(reader io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *HandshakeServer) ReadC0C1(reader io.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// s1

// s1

// s2
// make digest to s2 suffix position

func (s *HandshakeServer) WriteS0S1S2(writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *HandshakeServer) ReadC2(reader io.Reader) error { _ = "STUB: not implemented"; return nil }

// c0c1 clientPartKey serverFullKey
// s0s1 serverPartKey clientFullKey
func parseChallenge(b []byte, peerKey []byte, key []byte) []byte {
	_ = "STUB: not implemented"
	//if b[0] != version {
	//	return nil, ErrRtmp
	//}
	return nil
}

// use c0c1 digest to make a new digest

// @param b c1或s1
func findDigest(b []byte, base int, key []byte) int {
	_ = "STUB: not implemented"
	// calc offs
	return 0
}

// calc digest

// compare origin digest in buffer with calced digest

// <b> could be `c1` or `s1` or `s2`
func makeDigestWithoutCenterPart(b []byte, offs int, key []byte, out []byte) {
	_ = "STUB: not implemented"
	return
}

// left

// right

// calc

func makeDigest(b []byte, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func random1528(out []byte) { _ = "STUB: not implemented"; return }
