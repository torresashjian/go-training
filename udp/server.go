package udp

import (
	"fmt"
	"github.com/torresashjian/go-training/log"
	"net"
	"os"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func Listen() {
	addr, err := net.ResolveUDPAddr("udp", ":10001")
	CheckError(err)
	log.Info.Printf("Address resolution IP '%s', Port '%d', Zone '%s'", addr.IP, addr.Port, addr.Zone)
	/* Now listen at selected port */
	conn, err := net.ListenUDP("udp", addr)
	CheckError(err)
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buf)
		log.Info.Printf("Received '%s' from '%s'", string(buf[0:n]), addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
