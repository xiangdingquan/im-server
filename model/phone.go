package model

import (
	"encoding/json"

	"open.chat/mtproto"
)

type VoIPServerConfig struct {
	AudioFrameSize                 int     `json:"audio_frame_size"`                              // 60
	WebrtcNsLevelVpio              int     `json:"webrtc_ns_level_vpio,omitempty"`                // 0
	WebrtcNsLevel                  int     `json:"webrtc_ns_level,omitempty"`                     // 2
	WebrtcAgcTargetLevel           int     `json:"webrtc_agc_target_level,omitempty"`             // 9
	WebrtcAgcEnableLimiter         bool    `json:"webrtc_agc_enable_limiter,omitempty"`           // true
	WebrtcAgcCompressionGain       int     `json:"webrtc_agc_compression_gain,omitempty"`         // 20
	JitterMinDelay20               int     `json:"jitter_min_delay_20,omitempty"`                 // 6
	JitterMaxDelay20               int     `json:"jitter_max_delay_20,omitempty"`                 // 25
	JitterMaxSlots20               int     `json:"jitter_max_slots_20,omitempty"`                 // 50
	JitterMinDelay40               int     `json:"jitter_min_delay_40,omitempty"`                 // 4
	JitterMaxDelay40               int     `json:"jitter_max_delay_40,omitempty"`                 // 15
	JitterMaxSlots40               int     `json:"jitter_max_slots_40,omitempty"`                 // 30
	JitterMinDelay60               int     `json:"jitter_min_delay_60,omitempty"`                 // 2
	JitterMaxDelay60               int     `json:"jitter_max_delay_60,omitempty"`                 // 10
	JitterMaxSlots60               int     `json:"jitter_max_slots_60,omitempty"`                 // 20
	JitterInitialDelay60           int     `json:"jitter_initial_delay_60,omitempty"`             // 2
	JitterInitialDelay40           int     `json:"jitter_initial_delay_40,omitempty"`             // 4
	JitterInitialDelay20           int     `json:"jitter_initial_delay_20,omitempty"`             // 6
	JitterLossesToReset            int     `json:"jitter_losses_to_reset,omitempty"`              // 20
	JitterResyncThreshold          float64 `json:"jitter_resync_threshold,omitempty"`             // 1.0
	UseIosVpioAgc                  bool    `json:"use_ios_vpio_agc,omitempty"`                    // true
	UseOsxVpioAgc                  bool    `json:"use_osx_vpio_agc,omitempty"`                    // true
	AudioCongestionWindow          int     `json:"audio_congestion_window,omitempty"`             // 1024
	Nat64FallbackTimeout           int     `json:"nat_64_fallback_timeout,omitempty"`             // 3
	AudioVadNoVoiceBitrate         int     `json:"audio_vad_no_voice_bitrate,omitempty"`          // 6000
	AudioVadBandwidth              int     `json:"audio_vad_bandwidth,omitempty"`                 // 3
	AudioVadNoVoiceBandwidth       int     `json:"audio_vad_no_voice_bandwidth,omitempty"`        // 0
	AudioExtraEcBandwidth          int     `json:"audio_extra_ec_bandwidth,omitempty"`            // 2
	AudioMaxBitrate                int     `json:"audio_max_bitrate,omitempty"`                   // 20000
	AudioMaxBitrateGprs            int     `json:"audio_max_bitrate_gprs,omitempty"`              // 8000
	AudioMaxBitrateEdge            int     `json:"audio_max_bitrate_edge,omitempty"`              // 16000
	AudioMaxBitrateSaving          int     `json:"audio_max_bitrate_saving,omitempty"`            // 8000
	AudioInitBitrate               int     `json:"audio_init_bitrate,omitempty"`                  // 16000
	AudioInitBitrateGprs           int     `json:"audio_init_bitrate_gprs,omitempty"`             // 8000
	AudioInitBitrateEdge           int     `json:"audio_init_bitrate_edge,omitempty"`             // 8000
	AudioInitBitrateSaving         int     `json:"audio_init_bitrate_saving,omitempty"`           // 8000
	AudioBitrateStepIncr           int     `json:"audio_bitrate_step_incr,omitempty"`             // 1000
	AudioBitrateStepDecr           int     `json:"audio_bitrate_step_decr,omitempty"`             // 1000
	AudioMinBitrate                int     `json:"audio_min_bitrate,omitempty"`                   // 8000
	RelaySwitchThreshold           float64 `json:"relay_switch_threshold,omitempty"`              // 0.8
	P2pToRelaySwitchThreshold      float64 `json:"p2p_to_relay_switch_threshold,omitempty"`       // 0.6
	RelayToP2pSwitchThreshold      float64 `json:"relay_to_p2p_switch_threshold,omitempty"`       // 0.8
	ReconnectingStateTimeout       float64 `json:"reconnecting_state_timeout,omitempty"`          // 2.0
	RateFlags                      uint32  `json:"rate_flags,omitempty"`                          // 0xFFFFFFFF
	RateMinRtt                     float64 `json:"rate_min_rtt,omitempty"`                        // 0.6
	RateMinSendLoss                float64 `json:"rate_min_send_loss,omitempty"`                  // 0.2
	PacketLossForExtraEc           float64 `json:"packet_loss_for_extra_ec,omitempty"`            // 0.02
	MaxUnsentStreamPackets         int     `json:"max_unsent_stream_packets,omitempty"`           // 2
	BadCallRating                  bool    `json:"bad_call_rating,omitempty"`                     // false
	EstablishedDelayIfNoStreamData float64 `json:"established_delay_if_no_stream_data,omitempty"` // 1.5
	UseTcp                         bool    `json:"use_tcp,omitempty"`                             // true
	ForceTcp                       bool    `json:"force_tcp,omitempty"`                           // false
	UseSystemNs                    bool    `json:"use_system_ns,omitempty"`                       // true
	UseSystemAec                   bool    `json:"use_system_aec,omitempty"`                      // true
	AdspGoodImpls                  string  `json:"adsp_good_impls,omitempty"`
	AaecGoodImpls                  string  `json:"aaec_good_impls,omitempty"`
	AnsGoodImpls                   string  `json:"ans_good_impls,omitempty"`
	AudioMediumFecBitrate          int     `json:"audio_medium_fec_bitrate,omitempty"`    // 20000
	AudioMediumFecMultiplier       float64 `json:"audio_medium_fec_multiplier,omitempty"` // 0.1
	AudioStrongFecBitrate          int     `json:"audio_strong_fec_bitrate,omitempty"`    // 7000

}

const (
	MinPhoneCallLayer = 65
	MaxPhoneCallLayer = 92
)

var callConfigDataJson string

func init() {
	var callConfig = &VoIPServerConfig{
		AudioFrameSize:           60,
		JitterMinDelay60:         2,
		JitterMaxDelay60:         10,
		JitterMaxSlots60:         20,
		JitterLossesToReset:      20,
		JitterResyncThreshold:    0.5,
		AudioCongestionWindow:    1024,
		AudioMaxBitrate:          20000,
		AudioMaxBitrateEdge:      16000,
		AudioMaxBitrateSaving:    8000,
		AudioInitBitrate:         16000,
		AudioInitBitrateEdge:     8000,
		AudioInitBitrateGprs:     8000,
		AudioInitBitrateSaving:   8000,
		AudioBitrateStepIncr:     1000,
		AudioBitrateStepDecr:     1000,
		UseSystemNs:              true,
		UseSystemAec:             true,
		ForceTcp:                 false,
		JitterInitialDelay60:     2,
		AdspGoodImpls:            "(Qualcomm Fluence)",
		BadCallRating:            true,
		UseIosVpioAgc:            false,
		UseTcp:                   false,
		AudioMediumFecBitrate:    20000,
		AudioMediumFecMultiplier: 0.1,
		AudioStrongFecBitrate:    7000,
	}

	jsonData, _ := json.Marshal(callConfig)
	callConfigDataJson = string(jsonData)
}

func GetCallConfigDataJson() string {
	return callConfigDataJson
}

const (
	CallStateNone      int32 = 0
	CallStateRequested int32 = 1
	CallStateReceived  int32 = 2
	CallStateAccepted  int32 = 3
	CallStateConfirmed int32 = 4
	CallStateCall      int32 = 5
	CallStateDiscarded int32 = 6
	CallStateError     int32 = 7
)

type PhoneCallSession struct {
	Video                bool                       `json:"video"`
	Id                   int64                      `json:"id"`
	AccessHash           int64                      `json:"access_hash"`
	AdminId              int32                      `json:"admin_id"`
	AdminAuthKeyId       int64                      `json:"admin_auth_key_id"`
	RandomId             int64                      `json:"random_id"`
	AdminProtocol        *mtproto.PhoneCallProtocol `json:"admin_protocol"`
	ParticipantId        int32                      `json:"participant_id"`
	ParticipantAuthKeyId int64                      `json:"participant_auth_key_id"`
	ParticipantProtocol  *mtproto.PhoneCallProtocol `json:"participant_protocol"`
	GAHash               []byte                     `json:"ga_hash"`
	GA                   []byte                     `json:"ga"`
	GB                   []byte                     `json:"gb"`
	KeyFingerprint       int64                      `json:"key_fingerprint"`
	State                int32                      `json:"state"`
	Date                 int64                      `json:"date"`
}

const (
	MinPhoneCallLibraryVersion = "2.4.4"
	MaxPhoneCallLibraryVersion = "2.7.7"
)

func CalcPhoneCallLibraryVersion(versions ...[]string) string {
	var (
		b bool
	)

	for _, v1 := range versions {
		b = false
		for _, v2 := range v1 {
			if v2 == MaxPhoneCallLibraryVersion {
				b = true
				break
			}
		}
		if !b {
			break
		}
	}

	if b {
		return MaxPhoneCallLibraryVersion
	} else {
		return MinPhoneCallLibraryVersion
	}
}
