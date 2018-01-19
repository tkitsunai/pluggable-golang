package protocol_test

import (
	"testing"

	"github.com/tkitsunai/pluggable-golang/protocol"
)

func TestProtocol_UNIXドメインソケットは文字列で表現することができる(t *testing.T) {
	if protocol.UNIX.String() != "unix" {
		t.Fatal("failed UNIX enum idiom")
	}
}

func TestProtocol_TCPは文字列で表現することができる(t *testing.T) {
	if protocol.TCP.String() != "tcp" {
		t.Fatal("failed TCP enum idiom")
	}
}
