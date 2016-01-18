package ipCircBuffer

import (
	"net"
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

func (i *IPCircBuffer) Set(stringIP string) {
	i.cur = (i.cur + 1) % i.length
	i.ips[i.cur] = net.ParseIP(stringIP)
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
		i.cur += i.length
	}
}

func NewIPCircBuffer() *IPCircBuffer {
	temp := &IPCircBuffer{}
	temp.ips = make([]net.IP, 5)
	temp.cur = 0
	temp.length = 5
	return &temp
}

func Init() {
	DesktopIP = NewIPCircBuffer()
	RaspberryIP = NewIPCircBuffer()
}
