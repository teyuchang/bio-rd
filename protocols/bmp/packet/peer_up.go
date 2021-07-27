package packet

import (
	"bytes"

	bgppkt "github.com/bio-routing/bio-rd/protocols/bgp/packet"
	"github.com/bio-routing/bio-rd/util/decoder"
	"github.com/pkg/errors"
)

const (
	// OpenMsgMinLen is the minimal length of a BGP open message
	OpenMsgMinLen = 29
)

// PeerUpNotification represents a peer up notification
type PeerUpNotification struct {
	CommonHeader    *CommonHeader
	PerPeerHeader   *PerPeerHeader
	LocalAddress    [16]byte
	LocalPort       uint16
	RemotePort      uint16
	SentOpenMsg     *bgppkt.BGPMessage
	ReceivedOpenMsg *bgppkt.BGPMessage
	Information     []byte
}

// MsgType returns the type of this message
func (p *PeerUpNotification) MsgType() uint8 {
	return p.CommonHeader.MsgType
}

func decodePeerUpNotification(buf *bytes.Buffer, ch *CommonHeader) (*PeerUpNotification, error) {
	p := &PeerUpNotification{
		CommonHeader: ch,
	}

	pph, err := decodePerPeerHeader(buf)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to decode per peer header")
	}

	p.PerPeerHeader = pph

	fields := []interface{}{
		&p.LocalAddress,
		&p.LocalPort,
		&p.RemotePort,
	}

	err = decoder.Decode(buf, fields)
	if err != nil {
		return nil, err
	}

	p.SentOpenMsg, err = bgppkt.Decode(buf, &bgppkt.DecodeOptions{
		AddPathIPv4Unicast: false,
		AddPathIPv6Unicast: false,
		Use32BitASN:        true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Unable to get OPEN message")
	}

	p.ReceivedOpenMsg, err = bgppkt.Decode(buf, &bgppkt.DecodeOptions{
		AddPathIPv4Unicast: false,
		AddPathIPv6Unicast: false,
		Use32BitASN:        true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Unable to get OPEN message")
	}

	if buf.Len() == 0 {
		return p, nil
	}

	p.Information = make([]byte, buf.Len())
	fields = []interface{}{
		&p.Information,
	}

	// This can not fail as p.Information has exactly the size of what is left in buf
	decoder.Decode(buf, fields)

	return p, nil
}

func getOpenMsg(buf *bytes.Buffer) ([]byte, error) {
	msg := make([]byte, OpenMsgMinLen)

	fields := []interface{}{
		&msg,
	}
	err := decoder.Decode(buf, fields)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read")
	}

	if msg[OpenMsgMinLen-1] == 0 {
		return msg, nil
	}

	optParams := make([]byte, msg[OpenMsgMinLen-1])
	fields = []interface{}{
		&optParams,
	}

	err = decoder.Decode(buf, fields)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read")
	}

	msg = append(msg, optParams...)
	return msg, nil
}
