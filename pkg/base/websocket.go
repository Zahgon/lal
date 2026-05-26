// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: joestarzxh

package base

import (
	"bufio"
)

// WsOpcode The WebSocket Protocol
// https://tools.ietf.org/html/rfc6455
//
// 0                   1                   2                   3
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-------+-+-------------+-------------------------------+
// |F|R|R|R| opcode|M| Payload len |    Extended payload length    |
// |I|S|S|S|  (4)  |A|     (7)     |             (16/64)           |
// |N|V|V|V|       |S|             |   (if payload len==126/127)   |
// | |1|2|3|       |K|             |                               |
// +-+-+-+-+-------+-+-------------+ - - - - - - - - - - - - - - - +
// |     Extended payload length continued, if payload len == 127  |
// + - - - - - - - - - - - - - - - +-------------------------------+
// |                               |Masking-key, if MASK set to 1  |
// +-------------------------------+-------------------------------+
// | Masking-key (continued)       |          Payload Data         |
// +-------------------------------- - - - - - - - - - - - - - - - +
// :                     Payload Data continued ...                :
// + - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - +
// |                     Payload Data continued ...                |
// +---------------------------------------------------------------+
// opcode:
// *  %x0 denotes a continuation frame
// *  %x1 denotes a text frame
// *  %x2 denotes a binary frame
// *  %x3-7 are reserved for further non-control frames
// *  %x8 denotes a connection close
// *  %x9 denotes a ping
// *  %xA denotes a pong
// *  %xB-F are reserved for further control frames
// Payload length:  7 bits, 7+16 bits, or 7+64 bits
// Masking-key:  0 or 4 bytes
// mark 加密
//
//	for i := 0; i < datalen; i {
//	    m := markingkeys[i%4]
//	    data[i] = msg[i] ^ m
//	}
type WsOpcode = uint8

const (
	Wso_Continuous WsOpcode = iota //连续消息片断
	Wso_Text                       //文本消息片断,
	Wso_Binary                     //二进制消息片断,

	// Wso_Rsv3 非控制消息片断保留的操作码,
	Wso_Rsv3
	Wso_Rsv4
	Wso_Rsv5
	Wso_Rsv6
	Wso_Rsv7
	Wso_Close //连接关闭,
	Wso_Ping  //心跳检查的ping,
	Wso_Pong  //心跳检查的pong,

	// Wso_RsvB 为将来的控制消息片断的保留操作码
	Wso_RsvB
	Wso_RsvC
	Wso_RsvD
	Wso_RsvE
	Wso_RsvF
)

type WsHeader struct {
	Fin    bool
	Rsv1   bool
	Rsv2   bool
	Rsv3   bool
	Opcode WsOpcode

	PayloadLength uint64

	Masked  bool
	MaskKey uint32
}

const WsMagicStr = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

func MakeWsFrameHeader(wsHeader WsHeader) (buf []byte) { _ = "STUB: not implemented"; return nil }

func UpdateWebSocketHeader(secWebSocketKey, protocol string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func ReadWsPayload(r *bufio.Reader) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func cipher(payload []byte, mask []byte, offset int) { _ = "STUB: not implemented"; return }

// Calculate position in mask due to previously processed bytes number.

// Count number of bytes will processed one by one from the beginning of payload.

// Count number of bytes will processed one by one from the end of payload.
// This is done to process payload by 8 bytes in each iteration of main loop.

// NOTE: we use here binary.LittleEndian regardless of what is real
// endianness on machine is. To do so, we have to use binary.LittleEndian in
// the masking loop below as well.

// Skip already processed right part.
// Get number of uint64 parts remaining to process.

// remain maps position in masking key [0,4) to number
// of bytes that need to be processed manually inside Cipher().
var remain = [4]int{0, 3, 2, 1}
