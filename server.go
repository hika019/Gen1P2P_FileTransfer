package main

import (
	"Gen1P2P_FileTransfer/lib"
	"fmt"
	"net"
	"os"
)

func main() {
	protocol := "tcp"
	port := ":55555"

	tcpAddr, err := net.ResolveTCPAddr(protocol, port)
	lib.CheckErrExit(err)

	listner, err := net.ListenTCP(protocol, tcpAddr)
	lib.CheckErrExit(err)

	for {
		conn, err := listner.Accept()
		if err != nil {
			continue
		}

		handle(conn)
	}

}

func handle(conn net.Conn) {
	addr, ok := conn.RemoteAddr().(*net.TCPAddr)
	if !ok {
		return
	}

	senderIP := addr.IP.String()

	fmt.Println(senderIP)

	messBuff := make([]byte, 120)
	messLen, err := conn.Read(messBuff)
	if messLen == 0 {
		return
	}
	if lib.CheckErr(err) {
		return
	}

	fp, err := os.OpenFile("test.jpg", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if lib.CheckErr(err) {
		return
	}

	defer fp.Close()

	_, err = fp.Write(messBuff)
	lib.CheckErrExit(err)
}
