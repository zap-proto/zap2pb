// Package zap2pb is the sanctioned ZAP↔protobuf-wire boundary — the single
// place in the Hanzo stack allowed to import google.golang.org/protobuf.
package zap2pb

import (
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
)

// Marshal encodes m to protobuf wire bytes.
func Marshal(m proto.Message) ([]byte, error) { return proto.Marshal(m) }

// Equal reports whether a and b are semantically equal messages.
func Equal(a, b proto.Message) bool { return proto.Equal(a, b) }

// AppendMessage marshals m and appends it to dst as length-delimited field num;
// repeated calls with one num build a `repeated` field, byte-identical to
// marshaling the enclosing message.
func AppendMessage(dst []byte, num int, m proto.Message) ([]byte, error) {
	b, err := Marshal(m)
	if err != nil {
		return nil, err
	}
	dst = protowire.AppendTag(dst, protowire.Number(num), protowire.BytesType)
	return protowire.AppendBytes(dst, b), nil
}
