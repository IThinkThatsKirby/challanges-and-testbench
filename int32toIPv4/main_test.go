package int32toipv4

import (
	"testing"
)

func Test_ToBinary(t *testing.T) {
	if toBinary(2149583361) != "10000000001000000000101000000001" {
		t.Error("it aint the right binary dog")

	}
}
func Test_ToDecimal(t *testing.T) {
	if toDecimal("1101001") != "105" {
		t.Error("WRONG DECIMAL")
	}
}

func Test_ToIPv4(t *testing.T) {
	if toIPv4("10000000001000000000101000000001") != "128.32.10.1" {
		t.Error("the ip address aint right.")
	}
}

func Test_Int32ToIp(t *testing.T) {
	if Int32ToIp(2149583361) != "128.32.10.1" {
		t.Error("the address aint right.")
	}
}
