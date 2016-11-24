package udp

import (
	"fmt"
	"github.com/torresashjian/go-training/log"
	"net"
	"strconv"
	"time"
)

func Send() {
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	CheckError(err)

	localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	conn, err := net.DialUDP("udp", localAddr, serverAddr)
	CheckError(err)

	defer conn.Close()
	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		_, err := conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		} else {
			log.Info.Printf("Message '%s' sent from client '%s' to server '%s'", msg, localAddr, serverAddr)
		}
		time.Sleep(time.Second * 1)
	}
}
