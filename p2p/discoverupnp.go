package p2p

import (
	"fmt"
	"github.com/huin/goupnp"
	"github.com/huin/goupnp/dcps/internetgateway1"
	"github.com/huin/goupnp/dcps/internetgateway2"
)

// discoverUPnP searches for Internet Gateway Devices
// and returns the first one it can find on the local network.
func discoverUPnP(target string) {
	devs, err := goupnp.DiscoverDevices(target)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("Devices found for target %s: %d ", target, len(devs)))
	for i := 0; i < len(devs); i++ {
		fmt.Println(fmt.Sprintf("Device Location: %s ", devs[i].Location.Path))
	}
}

func DiscoverAllUPnP() {
	discoverUPnP(internetgateway1.URN_WANConnectionDevice_1)
	discoverUPnP(internetgateway2.URN_WANConnectionDevice_2)
}
