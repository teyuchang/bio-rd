package packet

import (
	"bytes"
	"fmt"

	"github.com/bio-routing/bio-rd/util/decoder"
)

type Stat struct {
	StatType   uint16
	StatLength uint16
	Counter    uint32
	Gauge      uint64
}

func decodeStat(buf *bytes.Buffer) (*Stat, error) {
	stat := &Stat{}

	fields := []interface{}{
		&stat.StatType,
		&stat.StatLength,
	}

	err := decoder.Decode(buf, fields)
	if err != nil {
		return stat, err
	}

	switch stat.StatLength {
	case 4:
		fields = []interface{}{
			&stat.Counter,
		}
	case 8:
		fields = []interface{}{
			&stat.Gauge,
		}
	default:
		return nil, fmt.Errorf("unsupported stat length: %d", stat.StatLength)
	}

	err = decoder.Decode(buf, fields)
	if err != nil {
		return stat, err
	}

	return stat, nil
}
