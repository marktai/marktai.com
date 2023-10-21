package ipCircBuffer_test

import (
	"ipCircBuffer"
	"net"
	"testing"
)

type IPCircBuffer struct {
	ips    []net.IP
	cur    int
	length int
}

var (
	DesktopIP   *IPCircBuffer
	RaspberryIP *IPCircBuffer
)

func TestSet(t *testing.T) {
	testIP := ipCircBuffer.NewIPCircBuffer()
	testIP.Set("1.1.1.1")
	if !testIP.Get().Equal(net.ParseIP("1.1.1.1")) {
		t.Error("Expected equal to 1.1.1.1")
	}

	testIP.Set("2.2.2.2")
	if !testIP.Get().Equal(net.ParseIP("2.2.2.2")) {
		t.Error("Expected equal to 2.2.2.2")
	}

}

func TestClear(t *testing.T) {
	testIP := ipCircBuffer.NewIPCircBuffer()
	testIP.Set("1.1.1.1")
	testIP.Set("2.2.2.2")
	testIP.Clear()
	if testIP.Get() != nil {
		t.Error("Expected nil")
	}
}

func TestRevert(t *testing.T) {
	testIP := ipCircBuffer.NewIPCircBuffer()
	testIP.Set("0.0.0.0")
	testIP.Set("0.0.0.0")
	testIP.Set("0.0.0.0")
	testIP.Set("0.0.0.0")
	testIP.Set("0.0.0.0")
	testIP.Set("1.1.1.1")
	testIP.Set("2.2.2.2")
	testIP.Set("3.3.3.3")

	if !testIP.Get().Equal(net.ParseIP("3.3.3.3")) {
		t.Error("Expected equal to 3.3.3.3")
	}

	testIP.Revert()
	if !testIP.Get().Equal(net.ParseIP("2.2.2.2")) {
		t.Error("Expected equal to 2.2.2.2")
	}

	testIP.Revert()
	if !testIP.Get().Equal(net.ParseIP("1.1.1.1")) {
		t.Error("Expected equal to 1.1.1.1")
	}
}
