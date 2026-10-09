package conf

import "encoding/json"

// FinalMask is accepted only so that configs which still carry the key fail
// loudly: packet masks and QUIC parameters were removed from this build, and
// silently ignoring them would send traffic unmasked.
type FinalMask struct {
	Tcp        []json.RawMessage `json:"tcp"`
	Udp        []json.RawMessage `json:"udp"`
	QuicParams json.RawMessage   `json:"quicParams"`
}

func (m *FinalMask) isEmpty() bool {
	return len(m.Tcp) == 0 && len(m.Udp) == 0 && len(m.QuicParams) == 0
}
