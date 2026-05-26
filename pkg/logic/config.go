// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"github.com/q191201771/lal/pkg/hls"
	"github.com/q191201771/lal/pkg/rtsp"
	"github.com/q191201771/naza/pkg/nazalog"
)

const (
	defaultHlsCleanupMode    = hls.CleanupModeInTheEnd
	defaultHttpflvUrlPattern = "/live/"
	defaultHttptsUrlPattern  = "/live/"
	defaultHlsUrlPattern     = "/hls/"
)

type Config struct {
	ConfVersion           string                `json:"conf_version"`
	RtmpConfig            RtmpConfig            `json:"rtmp"`
	InSessionConfig       InSessionConfig       `json:"in_session"`
	DefaultHttpConfig     DefaultHttpConfig     `json:"default_http"`
	HttpflvConfig         HttpflvConfig         `json:"httpflv"`
	HlsConfig             HlsConfig             `json:"hls"`
	HttptsConfig          HttptsConfig          `json:"httpts"`
	RtspConfig            RtspConfig            `json:"rtsp"`
	RecordConfig          RecordConfig          `json:"record"`
	RelayPushConfig       RelayPushConfig       `json:"relay_push"`
	StaticRelayPullConfig StaticRelayPullConfig `json:"static_relay_pull"`

	HttpApiConfig    HttpApiConfig    `json:"http_api"`
	ServerId         string           `json:"server_id"`
	HttpNotifyConfig HttpNotifyConfig `json:"http_notify"`
	SimpleAuthConfig SimpleAuthConfig `json:"simple_auth"`
	PprofConfig      PprofConfig      `json:"pprof"`
	LogConfig        nazalog.Option   `json:"log"`
	DebugConfig      DebugConfig      `json:"debug"`
}

type RtmpConfig struct {
	Enable               bool   `json:"enable"`
	Addr                 string `json:"addr"`
	RtmpsEnable          bool   `json:"rtmps_enable"`
	RtmpsAddr            string `json:"rtmps_addr"`
	RtmpsCertFile        string `json:"rtmps_cert_file"`
	RtmpsKeyFile         string `json:"rtmps_key_file"`
	GopNum               int    `json:"gop_num"` // TODO(chef): refactor 更名为gop_cache_num
	SingleGopMaxFrameNum int    `json:"single_gop_max_frame_num"`
	MergeWriteSize       int    `json:"merge_write_size"`
}

type InSessionConfig struct {
	AddDummyAudioEnable      bool `json:"add_dummy_audio_enable"`
	AddDummyAudioWaitAudioMs int  `json:"add_dummy_audio_wait_audio_ms"`
}

type DefaultHttpConfig struct {
	CommonHttpAddrConfig
}

type HttpflvConfig struct {
	CommonHttpServerConfig

	GopNum               int `json:"gop_num"`
	SingleGopMaxFrameNum int `json:"single_gop_max_frame_num"`
}

type HttptsConfig struct {
	CommonHttpServerConfig

	GopNum               int `json:"gop_num"`
	SingleGopMaxFrameNum int `json:"single_gop_max_frame_num"`
}

type HlsConfig struct {
	CommonHttpServerConfig

	UseMemoryAsDiskFlag bool `json:"use_memory_as_disk_flag"`
	hls.MuxerConfig
	SubSessionTimeoutMs int    `json:"sub_session_timeout_ms"`
	SubSessionHashKey   string `json:"sub_session_hash_key"`
}

type RtspConfig struct {
	Enable              bool   `json:"enable"`
	Addr                string `json:"addr"`
	RtspsEnable         bool   `json:"rtsps_enable"`
	RtspsAddr           string `json:"rtsps_addr"`
	RtspsCertFile       string `json:"rtsps_cert_file"`
	RtspsKeyFile        string `json:"rtsps_key_file"`
	OutWaitKeyFrameFlag bool   `json:"out_wait_key_frame_flag"`
	WsRtspEnable        bool   `json:"ws_rtsp_enable"`
	WsRtspAddr          string `json:"ws_rtsp_addr"`
	rtsp.ServerAuthConfig
}

type RecordConfig struct {
	EnableFlv     bool   `json:"enable_flv"`
	FlvOutPath    string `json:"flv_out_path"`
	EnableMpegts  bool   `json:"enable_mpegts"`
	MpegtsOutPath string `json:"mpegts_out_path"`
}

type RelayPushConfig struct {
	Enable   bool     `json:"enable"`
	AddrList []string `json:"addr_list"`
}

type StaticRelayPullConfig struct {
	Enable bool   `json:"enable"`
	Addr   string `json:"addr"`
}

type HttpApiConfig struct {
	Enable bool   `json:"enable"`
	Addr   string `json:"addr"`
}

type HttpNotifyConfig struct {
	Enable            bool   `json:"enable"`
	UpdateIntervalSec int    `json:"update_interval_sec"`
	OnServerStart     string `json:"on_server_start"`
	OnUpdate          string `json:"on_update"`
	OnPubStart        string `json:"on_pub_start"`
	OnPubStop         string `json:"on_pub_stop"`
	OnSubStart        string `json:"on_sub_start"`
	OnSubStop         string `json:"on_sub_stop"`
	OnRelayPullStart  string `json:"on_relay_pull_start"`
	OnRelayPullStop   string `json:"on_relay_pull_stop"`
	OnRtmpConnect     string `json:"on_rtmp_connect"`
	OnHlsMakeTs       string `json:"on_hls_make_ts"`
}

type SimpleAuthConfig struct {
	Key                string `json:"key"`
	DangerousLalSecret string `json:"dangerous_lal_secret"`
	PubRtmpEnable      bool   `json:"pub_rtmp_enable"`
	SubRtmpEnable      bool   `json:"sub_rtmp_enable"`
	SubHttpflvEnable   bool   `json:"sub_httpflv_enable"`
	SubHttptsEnable    bool   `json:"sub_httpts_enable"`
	PubRtspEnable      bool   `json:"pub_rtsp_enable"`
	SubRtspEnable      bool   `json:"sub_rtsp_enable"`
	HlsM3u8Enable      bool   `json:"hls_m3u8_enable"`
}

type PprofConfig struct {
	Enable bool   `json:"enable"`
	Addr   string `json:"addr"`
}

type DebugConfig struct {
	LogGroupIntervalSec       int `json:"log_group_interval_sec"`
	LogGroupMaxGroupNum       int `json:"log_group_max_group_num"`
	LogGroupMaxSubNumPerGroup int `json:"log_group_max_sub_num_per_group"`
}

type CommonHttpServerConfig struct {
	CommonHttpAddrConfig

	Enable      bool   `json:"enable"`
	EnableHttps bool   `json:"enable_https"`
	UrlPattern  string `json:"url_pattern"`
}

type CommonHttpAddrConfig struct {
	HttpListenAddr  string `json:"http_listen_addr"`
	HttpsListenAddr string `json:"https_listen_addr"`
	HttpsCertFile   string `json:"https_cert_file"`
	HttpsKeyFile    string `json:"https_key_file"`
}

func LoadConfAndInitLog(rawContent []byte) *Config {
	_ = "STUB: not implemented"

	// 读取配置并解析原始内容
	return nil
}

// 初始化日志模块，注意，这一步尽量提前，使得后续的日志内容按我们的日志配置输出
//
// 日志配置项不存在时，设置默认值
//
// 注意，由于此时日志模块还没有初始化，所以有日志需要打印时，我们采用先缓存后打印（日志模块初始化成功后再打印）的方式

// 打印Logo

// 检查配置版本号是否匹配

// 做个全量字段检查，缺失的字段，Go中会先设置为零值

// 日志字段检查，缺失的字段，打印前面设置的默认值

// 如果具体的HTTP应用没有设置HTTP监听相关的配置，则尝试使用全局配置

// 为缺失的字段中的一些特定字段，设置特定默认值

// 没有设置超时值，或者超时为0时

// 对一些常见的格式错误做修复
// 确保url pattern以`/`开始，并以`/`结束

// 打印配置文件中的元素内容，以及解析后的最终值
// 把配置文件原始内容中的换行去掉，使得打印日志时紧凑一些

// ---------------------------------------------------------------------------------------------------------------------

func mergeCommonHttpAddrConfig(dst, src *CommonHttpAddrConfig) { _ = "STUB: not implemented"; return }

func ensureStartWithSlash(in string) (out string, changed bool) {
	_ = "STUB: not implemented"
	return "", false
}

func ensureEndWithSlash(in string) (out string, changed bool) {
	_ = "STUB: not implemented"
	return "", false
}

func ensureStartAndEndWithSlash(in string) (out string, changed bool) {
	_ = "STUB: not implemented"
	return "", false
}
