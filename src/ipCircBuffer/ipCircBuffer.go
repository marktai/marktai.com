package ipCircBuffer

import (
	"net"
)

type IPCircBuffer struct {
	ips [5]net.IP
	cur int
}

var (
	DesktopIP   *IPCircBuffer
	RaspberryIP *IPCircBuffer
)

func (i *IPCircBuffer) Set(stringIP string) {
	i.ips[i.cur+1] = net.ParseIP(stringIP)
	i.cur = (i.cur + 1) % 5
}

func (i *IPCircBuffer) Get() net.IP {
	return i.ips[i.cur]
}

func (i *IPCircBuffer) Clear() {
	for j, _ := range i.ips {
		i.ips[j] = nil
	}
}

func (i *IPCircBuffer) Revert() {
	i.cur = i.cur - 1
	if i.cur < 0 {
		i.cur += 5
	}
}

func NewIPCircBuffer() *IPCircBuffer {
	return &IPCircBuffer{}
}

func Init() {
	DesktopIP = NewIPCircBuffer()
	RaspberryIP = NewIPCircBuffer()
}
