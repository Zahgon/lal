// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtmp

import (
	"net"
	"sync"

	"github.com/q191201771/lal/pkg/base"

	"github.com/q191201771/naza/pkg/connection"
)

// TODO chef: 没有进化成Pub Sub时的超时释放

type IServerSessionObserver interface {
	OnRtmpConnect(session *ServerSession, opa ObjectPairArray)

	// OnNewRtmpPubSession
	//
	// 上层代码应该在这个事件回调中注册音视频数据的监听
	//
	// @return 上层如果想关闭这个session，则回调中返回不为nil的error值
	//
	OnNewRtmpPubSession(session *ServerSession) error

	OnNewRtmpSubSession(session *ServerSession) error
}

type IPubSessionObserver interface {
	// OnReadRtmpAvMsg 注意，回调结束后，内部会复用Payload内存块
	OnReadRtmpAvMsg(msg base.RtmpMsg)
}

func (s *ServerSession) SetPubSessionObserver(observer IPubSessionObserver) {
	_ = "STUB: not implemented"
	return
}

type ServerSessionType int

const (
	ServerSessionTypeUnknown ServerSessionType = iota // 收到客户端的publish或者play信令之前的类型状态
	ServerSessionTypePub
	ServerSessionTypeSub
)

type ServerSession struct {
	url                    string
	tcUrl                  string
	streamNameWithRawQuery string // const after set
	appName                string // const after set
	streamName             string // const after set
	rawQuery               string //const after set

	observer      IServerSessionObserver
	hs            HandshakeServer
	chunkComposer *ChunkComposer
	packer        *MessagePacker

	conn        connection.Connection
	sessionStat base.BasicSessionStat

	// only for PubSession
	avObserver IPubSessionObserver

	// IsFresh ShouldWaitVideoKeyFrame
	//
	// 只有sub类型需要
	//
	// IsFresh
	//  表示是新加入的session，需要新发送meta，vsh，ash以及gop等数据，再转发实时数据。
	//
	// ShouldWaitVideoKeyFrame
	//  表示是新加入的session，正在等待视频关键帧。
	//  注意，需要考虑没有纯音频流的场景。
	//
	IsFresh                 bool
	ShouldWaitVideoKeyFrame bool

	disposeOnce sync.Once

	DisposeByObserverFlag bool

	peerWinAckSize int
	recvLastAck    uint64
	seqNum         uint32
}

func NewServerSession(observer IServerSessionObserver, conn net.Conn) *ServerSession {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServerSession) RunLoop() (err error) { _ = "STUB: not implemented"; return nil }

func (s *ServerSession) Write(msg []byte) error { _ = "STUB: not implemented"; return nil }

func (s *ServerSession) Writev(msgs net.Buffers) error { _ = "STUB: not implemented"; return nil }

func (s *ServerSession) Flush() error { _ = "STUB: not implemented"; return nil }

// ----- IServerSessionLifecycle ---------------------------------------------------------------------------------------

func (s *ServerSession) Dispose() error { _ = "STUB: not implemented"; return nil }

// ----- ISessionUrlContext --------------------------------------------------------------------------------------------

func (s *ServerSession) Url() string { _ = "STUB: not implemented"; return "" }

func (s *ServerSession) AppName() string { _ = "STUB: not implemented"; return "" }

func (s *ServerSession) StreamName() string { _ = "STUB: not implemented"; return "" }

func (s *ServerSession) RawQuery() string {
	_ = "STUB: not implemented"

	// ----- IObject -------------------------------------------------------------------------------------------------------
	return ""
}

func (s *ServerSession) UniqueKey() string { _ = "STUB: not implemented"; return "" }

// ----- ISessionStat --------------------------------------------------------------------------------------------------

func (s *ServerSession) UpdateStat(intervalSec uint32) { _ = "STUB: not implemented"; return }

func (s *ServerSession) GetStat() base.StatSession {
	_ = "STUB: not implemented"
	return *new(base.StatSession)
}

func (s *ServerSession) IsAlive() (readAlive, writeAlive bool) {
	_ = "STUB: not implemented"
	return false, false
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *ServerSession) runReadLoop() error { _ = "STUB: not implemented"; return nil }

func (s *ServerSession) handshake() error { _ = "STUB: not implemented"; return nil }

func (s *ServerSession) doMsg(stream *Stream) error { _ = "STUB: not implemented"; return nil }

// noop
// 因为底层的 chunk composer 已经处理过了，这里就不用处理

func (s *ServerSession) doWinAckSize(stream *Stream) error { _ = "STUB: not implemented"; return nil }

func (s *ServerSession) doAck(stream *Stream) error { _ = "STUB: not implemented"; return nil }

func (s *ServerSession) doUserControl(stream *Stream) error {
	_ = "STUB: not implemented"
	// TODO(chef): 检查buff长度有效性 202301
	return nil
}

func (s *ServerSession) doDataMessageAmf0(stream *Stream) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO chef: 下面注释掉的代码包含的逻辑：
// 1. 去除metadata中@setDataFrame
// 2. 判断一些错误格式
// 如果这个逻辑不是必须的，就可以删掉了
// 另外，如果返回给上层的msg是删除了内容的buf，应该注意和header中的len保持一致
//
//switch val {
//case "|RtmpSampleAccess":
//	Log.Warnf("[%s] read data message, ignore it. val=%s", s.UniqueKey(), val)
//	return nil
//case "@setDataFrame":
//	// macos obs and ffmpeg
//	// skip @setDataFrame
//	val, err = stream.msg.readStringWithType()
//
//	val, err := stream.msg.peekStringWithType()
//	if err != nil {
//		return err
//	}
//	if val != "onMetaData" {
//		Log.Errorf("[%s] read unknown data message. val=%s, %s", s.UniqueKey(), val, stream.toDebugString())
//		return ErrRtmp
//	}
//case "onMetaData":
//	// noop
//default:
//	Log.Errorf("[%s] read unknown data message. val=%s, %s", s.UniqueKey(), val, stream.toDebugString())
//	return nil
//}
//
//s.avObserver.OnReadRtmpAvMsg(stream.toAvMsg())
//return nil

func (s *ServerSession) doCommandMessage(stream *Stream) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServerSession) doCommandAmf3Message(stream *Stream) error {
	_ = "STUB: not implemented"
	//去除前面的0就是Amf0的数据
	return nil
}

func (s *ServerSession) writeAcknowledgementIfNeeded(stream *Stream) error {
	_ = "STUB: not implemented"
	return nil
}

//此次接收小于窗口大小一半，不处理

//当序列号溢出时，将其重置

//时间戳暂时先发0

func (s *ServerSession) doConnect(tid int, stream *Stream) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServerSession) doCreateStream(tid int, stream *Stream) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServerSession) doPublish(tid int, stream *Stream) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// 兼容 https://github.com/q191201771/lal/issues/280
// 没有 pubType 时，继续走后面的流程

// 回复完信令后修改 connection 的属性

func (s *ServerSession) doPlay(tid int, stream *Stream) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// TODO chef: start duration reset

// 回复完信令后修改 connection 的属性

func (s *ServerSession) modConnProps() { _ = "STUB: not implemented"; return }

func (s *ServerSession) dispose(err error) error { _ = "STUB: not implemented"; return nil }
