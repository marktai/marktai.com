package desktopIP

import(
    "net"
)

var desktopIP net.IP

func Set(stringIP string) {
    desktopIP = net.ParseIP(stringIP)
}

func Get() net.IP {
    return desktopIP
}

func Clear() {
    desktopIP = nil
}

