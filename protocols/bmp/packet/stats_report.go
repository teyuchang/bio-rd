package packet

import (
	"bytes"

	"github.com/bio-routing/bio-rd/util/decoder"
	"github.com/pkg/errors"
)

// StatsReport represents a stats report message
type StatsReport struct {
	CommonHeader  *CommonHeader
	PerPeerHeader *PerPeerHeader
	StatsCount    uint32
	Stats         []*Stat
}

// MsgType returns the type of this message
func (s *StatsReport) MsgType() uint8 {
	return s.CommonHeader.MsgType
}

func decodeStatsReport(buf *bytes.Buffer, ch *CommonHeader) (Msg, error) {
	sr := &StatsReport{
		CommonHeader: ch,
	}

	pph, err := decodePerPeerHeader(buf)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to decode per peer header")
	}

	sr.PerPeerHeader = pph

	fields := []interface{}{
		&sr.StatsCount,
	}

	err = decoder.Decode(buf, fields)
	if err != nil {
		return sr, err
	}

	sr.Stats = make([]*Stat, sr.StatsCount)
	for i := uint32(0); i < sr.StatsCount; i++ {
		stat, err := decodeStat(buf)
		if err != nil {
			return sr, errors.Wrap(err, "Unable to decode information TLV")
		}

		sr.Stats[i] = stat
	}

	return sr, nil
}
