package zap2pb

import (
	"reflect"
	"testing"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ZAP struct -> pb wire (AppendMessage as repeated field 1) -> parse == expected.
func TestAppendMessageRoundTrip(t *testing.T) {
	want := []string{"chat zen", "hanzo"}
	var body []byte
	for _, s := range want {
		var err error
		if body, err = AppendMessage(body, 1, wrapperspb.String(s)); err != nil {
			t.Fatal(err)
		}
	}

	var got []string
	for len(body) > 0 {
		num, typ, n := protowire.ConsumeTag(body)
		if n < 0 || num != 1 || typ != protowire.BytesType {
			t.Fatalf("bad tag num=%d typ=%d n=%d", num, typ, n)
		}
		body = body[n:]
		b, n := protowire.ConsumeBytes(body)
		if n < 0 {
			t.Fatal("bad bytes")
		}
		body = body[n:]
		var sv wrapperspb.StringValue
		if err := proto.Unmarshal(b, &sv); err != nil {
			t.Fatal(err)
		}
		got = append(got, sv.Value)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestEqual(t *testing.T) {
	if !Equal(wrapperspb.String("x"), wrapperspb.String("x")) {
		t.Fatal("want equal")
	}
	if Equal(wrapperspb.String("x"), wrapperspb.String("y")) {
		t.Fatal("want not equal")
	}
}
