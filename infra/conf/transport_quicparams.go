package conf

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/xtls/xray-core/common/errors"
)

const (
	Byte     = 1
	Kilobyte = 1024 * Byte
	Megabyte = 1024 * Kilobyte
	Gigabyte = 1024 * Megabyte
	Terabyte = 1024 * Gigabyte
)

type Bandwidth string

func (b Bandwidth) Bps() (uint64, error) {
	s := strings.TrimSpace(strings.ToLower(string(b)))
	if s == "" {
		return 0, nil
	}

	idx := len(s)
	for i, c := range s {
		if (c < '0' || c > '9') && c != '.' {
			idx = i
			break
		}
	}

	numStr := s[:idx]
	unit := strings.TrimSpace(s[idx:])

	val, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, err
	}

	mul := uint64(1)
	switch unit {
	case "", "b", "bps":
		mul = Byte
	case "k", "kb", "kbps":
		mul = Kilobyte
	case "m", "mb", "mbps":
		mul = Megabyte
	case "g", "gb", "gbps":
		mul = Gigabyte
	case "t", "tb", "tbps":
		mul = Terabyte
	default:
		return 0, errors.New("unsupported unit: " + unit)
	}

	return uint64(val*float64(mul)) / 8, nil
}

type QuicParamsConfig struct {
	Congestion                    string    `json:"congestion"`
	Debug                         bool      `json:"debug"`
	BbrProfile                    string    `json:"bbrProfile"`
	BrutalUp                      Bandwidth `json:"brutalUp"`
	BrutalDown                    Bandwidth `json:"brutalDown"`
	BrutalDisableLossCompensation bool      `json:"brutalDisableLossCompensation"`
	InitStreamReceiveWindow       uint64    `json:"initStreamReceiveWindow"`
	MaxStreamReceiveWindow        uint64    `json:"maxStreamReceiveWindow"`
	InitConnectionReceiveWindow   uint64    `json:"initConnectionReceiveWindow"`
	MaxConnectionReceiveWindow    uint64    `json:"maxConnectionReceiveWindow"`
	MaxIdleTimeout                int64     `json:"maxIdleTimeout"`
	KeepAlivePeriod               int64     `json:"keepAlivePeriod"`
	DisablePathMTUDiscovery       bool      `json:"disablePathMTUDiscovery"`
	DisableChromeParrot           bool      `json:"disableChromeParrot"`
	DisableGSO                    bool      `json:"disableGSO"`
	MaxIncomingStreams            int64     `json:"maxIncomingStreams"`
	DisableStatelessReset         bool      `json:"disableStatelessReset"`
}

// FinalMask only carries QUIC parameters in this build. The packet-mask
// transforms (tcp/udp) were removed; configs that still list any fail loudly
// instead of silently sending unmasked traffic.
type FinalMask struct {
	Tcp        []json.RawMessage `json:"tcp"`
	Udp        []json.RawMessage `json:"udp"`
	QuicParams *QuicParamsConfig `json:"quicParams"`
}
