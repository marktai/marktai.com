package desktopIP

import (
	"net"
)

var desktopIPs [5]net.IP
var cur = 0

func Set(stringIP string) {
	desktopIPs[cur+1] = net.ParseIP(stringIP)
	cur = (cur + 1) % 5
}

func Get() net.IP {
	return desktopIPs[cur]
}

func Clear() {
	for i, _ := range desktopIPs {
		desktopIPs[i] = nil
	}
}

func Revert() {
	cur = cur - 1
	if cur < 0 {
		cur += 5
	}
}
