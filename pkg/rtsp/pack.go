// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

// rfc2326 10.3 ANNOUNCE
//var RequestAnnounceTmpl = "not impl"

// ResponseAnnounceTmpl CSeq
var ResponseAnnounceTmpl = "RTSP/1.0 200 OK\r\n" +
	"CSeq: %s\r\n" +
	"\r\n"

// rfc2326 10.2 DESCRIBE

// ResponseDescribeTmpl CSeq, Date, Content-Length,
var ResponseDescribeTmpl = "RTSP/1.0 200 OK\r\n" +
	"CSeq: %s\r\n" +
	"Date: %s\r\n" +
	"Content-Type: application/sdp\r\n" +
	"Content-Length: %d\r\n" +
	"\r\n" +
	"%s"

// ResponseSetupTmpl rfc2326 10.4 SETUP
// CSeq, Date, Session, Transport
var ResponseSetupTmpl = "RTSP/1.0 200 OK\r\n" +
	"CSeq: %s\r\n" +
	"Date: %s\r\n" +
	"Session: %s\r\n" +
	"Transport: %s\r\n" +
	"\r\n"

// rfc2326 10.11 RECORD
//var RequestRecordTmpl = "not impl"

// ResponseRecordTmpl CSeq, Session
var ResponseRecordTmpl = "RTSP/1.0 200 OK\r\n" +
	"CSeq: %s\r\n" +
	"Session: %s\r\n" +
	"\r\n"

// rfc2326 10.5 PLAY

// ResponsePlayTmpl CSeq Date
var ResponsePlayTmpl = "RTSP/1.0 200 OK\r\n" +
	"CSeq: %s\r\n" +
	"Date: %s\r\n" +
	"\r\n"

// rfc2326 10.7 TEARDOWN
//var RequestTeardownTmpl = "not impl"

// ResponseTeardownTmpl CSeq
var ResponseTeardownTmpl = "RTSP/1.0 200 OK\r\n" +
	"CSeq: %s\r\n" +
	"\r\n"

var ResponseAuthorizedTmpl = "RTSP/1.0 401 Unauthorized\r\n" +
	"CSeq: %s\r\n" +
	"Date: %s\r\n" +
	"WWW-Authenticate: %s\r\n" +
	"\r\n"

func PackResponseOptions(cseq string) string { _ = "STUB: not implemented"; return "" }

func PackResponseAnnounce(cseq string) string { _ = "STUB: not implemented"; return "" }

func PackResponseDescribe(cseq, sdp string) string { _ = "STUB: not implemented"; return "" }

func PackResponseSetup(cseq string, htv string) string { _ = "STUB: not implemented"; return "" }

func PackResponseRecord(cseq string) string { _ = "STUB: not implemented"; return "" }

func PackResponsePlay(cseq string) string { _ = "STUB: not implemented"; return "" }

func PackResponseTeardown(cseq string) string { _ = "STUB: not implemented"; return "" }

func PackResponseAuthorized(cseq, authenticate string) string { _ = "STUB: not implemented"; return "" }

// PackRequest @param body 可以为空
func PackRequest(method, uri string, headers map[string]string, body string) (ret string) {
	_ = "STUB: not implemented"
	return ""
}
