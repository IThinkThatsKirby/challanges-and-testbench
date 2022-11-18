package main

import (
	"testing"
)

func TestToBinary(t *testing.T) {
	if toBinary(2149583361) != "10000000001000000000101000000001" {
		t.Error("it aint the right binary dog")

	}
}
func TestToDecimal(t *testing.T) {
	if toDecimal("1101001") != "105" {
		t.Error("WRONG DECIMAL")
	}
}

func TestToIPv4(t *testing.T) {
	if toIPv4("10000000001000000000101000000001") != "128.32.10.1" {
		t.Error("the ip address aint right.")
	}
}

func TestInt32ToIp(t *testing.T) {
	if Int32ToIp(2149583361) != "128.32.10.1" {
		t.Error("the address aint right.")
	}
}
