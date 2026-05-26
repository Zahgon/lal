// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package base

import "github.com/q191201771/naza/pkg/unique"

const (
	UkPreCustomizePubSessionContext = SessionProtocolCustomizeStr + SessionBaseTypePubStr // "CUSTOMIZEPUB"
	UkPreRtmpServerSession          = SessionProtocolRtmpStr + SessionBaseTypePubSubStr   // "RTMPPUBSUB" // 两种可能，pub或者sub
	UkPreRtmpPushSession            = SessionProtocolRtmpStr + SessionBaseTypePushStr     // "RTMPPUSH"
	UkPreRtmpPullSession            = SessionProtocolRtmpStr + SessionBaseTypePullStr     // "RTMPPULL"
	UkPreRtspPubSession             = SessionProtocolRtspStr + SessionBaseTypePubStr      // "RTSPPUB"
	UkPreRtspSubSession             = SessionProtocolRtspStr + SessionBaseTypeSubStr      // "RTSPSUB"
	UkPreRtspPushSession            = SessionProtocolRtspStr + SessionBaseTypePushStr     // "RTSPPUSH"
	UkPreRtspPullSession            = SessionProtocolRtspStr + SessionBaseTypePullStr     // "RTSPPULL"
	UkPreFlvSubSession              = SessionProtocolFlvStr + SessionBaseTypePubSubStr    // "FLVSUB"
	UkPreFlvPullSession             = SessionProtocolFlvStr + SessionBaseTypePullStr      // "FLVPULL"
	UkPreTsSubSession               = SessionProtocolTsStr + SessionBaseTypePubSubStr     // "TSSUB"
	UkPrePsPubSession               = SessionProtocolPsStr + SessionBaseTypePubStr        // "PSPUB"
	UkPreHlsSubSession              = SessionProtocolHlsStr + SessionBaseTypeSubStr       // "HLSSUB"

	UkPreRtspServerCommandSession = "RTSPSRVCMD" // 这个不暴露给上层

	UkPreGroup              = "GROUP"
	UkPreHlsMuxer           = "HLSMUXER"
	UkPreRtmp2MpegtsRemuxer = "RTMP2MPEGTS"
)

//func GenUk(prefix string) string {
//	return unique.GenUniqueKey(prefix)
//}

func GenUkCustomizePubSession() string { _ = "STUB: not implemented"; return "" }

func GenUkRtmpServerSession() string { _ = "STUB: not implemented"; return "" }

func GenUkRtmpPushSession() string { _ = "STUB: not implemented"; return "" }

func GenUkRtmpPullSession() string { _ = "STUB: not implemented"; return "" }

func GenUkRtspServerCommandSession() string { _ = "STUB: not implemented"; return "" }

func GenUkRtspPubSession() string { _ = "STUB: not implemented"; return "" }

func GenUkRtspSubSession() string { _ = "STUB: not implemented"; return "" }

func GenUkRtspPushSession() string { _ = "STUB: not implemented"; return "" }

func GenUkRtspPullSession() string { _ = "STUB: not implemented"; return "" }

func GenUkFlvSubSession() string { _ = "STUB: not implemented"; return "" }

func GenUkTsSubSession() string { _ = "STUB: not implemented"; return "" }

func GenUkFlvPullSession() string { _ = "STUB: not implemented"; return "" }

func GenUkHlsSubSession() string { _ = "STUB: not implemented"; return "" }

func GenUkPsPubSession() string { _ = "STUB: not implemented"; return "" }

func GenUkGroup() string { _ = "STUB: not implemented"; return "" }

func GenUkHlsMuxer() string { _ = "STUB: not implemented"; return "" }

func GenUkRtmp2MpegtsRemuxer() string { _ = "STUB: not implemented"; return "" }

var (
	siUkCustomizePubSession      *unique.SingleGenerator
	siUkRtmpServerSession        *unique.SingleGenerator
	siUkRtmpPushSession          *unique.SingleGenerator
	siUkRtmpPullSession          *unique.SingleGenerator
	siUkRtspServerCommandSession *unique.SingleGenerator
	siUkRtspPubSession           *unique.SingleGenerator
	siUkRtspSubSession           *unique.SingleGenerator
	siUkRtspPushSession          *unique.SingleGenerator
	siUkRtspPullSession          *unique.SingleGenerator
	siUkFlvSubSession            *unique.SingleGenerator
	siUkTsSubSession             *unique.SingleGenerator
	siUkFlvPullSession           *unique.SingleGenerator
	siUkPsPubSession             *unique.SingleGenerator
	siUkHlsSubSession            *unique.SingleGenerator

	siUkGroup              *unique.SingleGenerator
	siUkHlsMuxer           *unique.SingleGenerator
	siUkRtmp2MpegtsRemuxer *unique.SingleGenerator
)

func init() {
	siUkCustomizePubSession = unique.NewSingleGenerator(UkPreCustomizePubSessionContext)
	siUkRtmpServerSession = unique.NewSingleGenerator(UkPreRtmpServerSession)
	siUkRtmpPushSession = unique.NewSingleGenerator(UkPreRtmpPushSession)
	siUkRtmpPullSession = unique.NewSingleGenerator(UkPreRtmpPullSession)
	siUkRtspServerCommandSession = unique.NewSingleGenerator(UkPreRtspServerCommandSession)
	siUkRtspPubSession = unique.NewSingleGenerator(UkPreRtspPubSession)
	siUkRtspSubSession = unique.NewSingleGenerator(UkPreRtspSubSession)
	siUkRtspPushSession = unique.NewSingleGenerator(UkPreRtspPushSession)
	siUkRtspPullSession = unique.NewSingleGenerator(UkPreRtspPullSession)
	siUkFlvSubSession = unique.NewSingleGenerator(UkPreFlvSubSession)
	siUkTsSubSession = unique.NewSingleGenerator(UkPreTsSubSession)
	siUkFlvPullSession = unique.NewSingleGenerator(UkPreFlvPullSession)
	siUkPsPubSession = unique.NewSingleGenerator(UkPrePsPubSession)
	siUkHlsSubSession = unique.NewSingleGenerator(UkPreHlsSubSession)

	siUkGroup = unique.NewSingleGenerator(UkPreGroup)
	siUkHlsMuxer = unique.NewSingleGenerator(UkPreHlsMuxer)
	siUkRtmp2MpegtsRemuxer = unique.NewSingleGenerator(UkPreRtmp2MpegtsRemuxer)
}
