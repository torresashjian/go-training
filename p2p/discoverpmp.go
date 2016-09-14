package p2p

import (
	"fmt"
	"net"
)

var (
	// LAN IP ranges
	_, lan10, _  = net.ParseCIDR("10.0.0.0/8")
	_, lan176, _ = net.ParseCIDR("172.16.0.0/12")
	_, lan192, _ = net.ParseCIDR("192.168.0.0/16")
)

func DiscoverAllPMPs() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	for _, iface := range ifaces {
		fmt.Println(fmt.Sprintf("Found Interface with name: %s", iface.Name))
		ifaddrs, err := iface.Addrs()
		if err != nil {
			fmt.Errorf(err.Error())
			return
		}
		for _, addr := range ifaddrs {
			fmt.Println(fmt.Sprintf("Address of type %T: %s", addr, addr))
			switch x := addr.(type) {
			case *net.IPNet:
				if lan10.Contains(x.IP) || lan176.Contains(x.IP) || lan192.Contains(x.IP) {
					ip := x.IP.Mask(x.Mask).To4()
					if ip != nil {
						ip[3] = ip[3] | 0x01
						fmt.Println(fmt.Sprintf("Ip : %s", ip))
					}
				}
			}
		}
	}
}
