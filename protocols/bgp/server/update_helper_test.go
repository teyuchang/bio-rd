package server

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/bio-routing/bio-rd/protocols/bgp/packet"

	bnet "github.com/bio-routing/bio-rd/net"
	"github.com/stretchr/testify/assert"
)

type failingUpdate struct{}

func (f *failingUpdate) SerializeUpdate(opt *packet.EncodeOptions) ([]byte, error) {
	return nil, errors.New("general error")
}

type WriterByter interface {
	Bytes() []byte
	io.Writer
}

type failingReadWriter struct {
}

func (f *failingReadWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("general error")
}

func (f *failingReadWriter) Bytes() []byte {
	return []byte{}
}

func TestSerializeAndSendUpdate(t *testing.T) {
	tests := []struct {
		name       string
		buf        WriterByter
		err        error
		testUpdate serializeAbleUpdate
		expected   []byte
	}{
		{
			name: "normal bgp update",
			buf:  bytes.NewBuffer(nil),
			err:  nil,
			testUpdate: &packet.BGPUpdate{
				WithdrawnRoutesLen: 5,
				WithdrawnRoutes: &packet.NLRI{
					IP:     bnet.IPv4FromOctets(10, 0, 0, 0),
					Pfxlen: 8,
					Next: &packet.NLRI{
						IP:     bnet.IPv4FromOctets(192, 168, 0, 0),
						Pfxlen: 16,
					},
				},
			},

			expected: []byte{
				255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, // Marker
				0, 28, // Length
				2,                               // Type = Update
				0, 5, 8, 10, 16, 192, 168, 0, 0, // 2 withdraws
			},
		},
		{
			name:       "failed serialization",
			buf:        bytes.NewBuffer(nil),
			err:        nil,
			testUpdate: &failingUpdate{},
			expected:   nil,
		},
		{
			name: "failed connection",
			buf:  &failingReadWriter{},
			err:  errors.New("Failed sending Update: general error"),
			testUpdate: &packet.BGPUpdate{
				WithdrawnRoutesLen: 5,
				WithdrawnRoutes: &packet.NLRI{
					IP:     bnet.IPv4FromOctets(10, 0, 0, 0),
					Pfxlen: 8,
					Next: &packet.NLRI{
						IP:     bnet.IPv4FromOctets(192, 168, 0, 0),
						Pfxlen: 16,
					},
				},
			},
			expected: []byte{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opt := &packet.EncodeOptions{}
			err := serializeAndSendUpdate(test.buf, test.testUpdate, opt)
			assert.Equal(t, test.err, err)

			assert.Equal(t, test.expected, test.buf.Bytes())
		})
	}
}
