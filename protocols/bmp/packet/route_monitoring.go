package packet

import (
	"bytes"

	bgppkt "github.com/bio-routing/bio-rd/protocols/bgp/packet"
	"github.com/pkg/errors"
)

// RouteMonitoringMsg represents a Route Monitoring Message
type RouteMonitoringMsg struct {
	CommonHeader  *CommonHeader
	PerPeerHeader *PerPeerHeader
	BGPUpdate     *bgppkt.BGPMessage
}

// MsgType returns the type of this message
func (rm *RouteMonitoringMsg) MsgType() uint8 {
	return rm.CommonHeader.MsgType
}

func decodeRouteMonitoringMsg(buf *bytes.Buffer, ch *CommonHeader) (*RouteMonitoringMsg, error) {
	rm := &RouteMonitoringMsg{
		CommonHeader: ch,
	}

	pph, err := decodePerPeerHeader(buf)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to decode per peer header")
	}

	rm.PerPeerHeader = pph

	rm.BGPUpdate, err = bgppkt.Decode(buf, &bgppkt.DecodeOptions{
		AddPathIPv4Unicast: false,
		AddPathIPv6Unicast: false,
		Use32BitASN:        true,
	})
	if err != nil {
		return nil, err
	}

	return rm, nil
}
