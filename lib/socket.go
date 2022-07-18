package lib

import (
	"net"
)

const protocol = "tcp"

/*
//soket
type	|DataLen|Payload|
		|		|		|
	1	|	4	|65536	|(byte)
		|		|		|
*/

const soketLen = 1 + 4 + 65536
const soketType = 0
const LenBegin = 1
const LenEnd = 4
const payloadBegin = 5
const payloadEnd = soketLen

//Upload FileHashTable
func UpFHT(conf ConfigStruct, hashs []byte) {
	//サーバへのFHTの転送
	protocol := "tcp"
	serverIP := conf.Server_addr
	serverPort := conf.Fht_port

	tcpAddr, err := net.ResolveTCPAddr(protocol, serverIP+":"+serverPort)
	CheckErrExit(err)

	myAddr := new(net.TCPAddr)
	myAddr.IP = net.ParseIP(MyIp())
	myAddr.Port = conf.Use_port

	conn, err := net.DialTCP(protocol, myAddr, tcpAddr)
	CheckErrExit(err)

	defer conn.Close()

}

func MyIp() string {
	interfaces, err := net.Interfaces()
	CheckErrExit(err)

	for _, inter := range interfaces {
		addrs, err := inter.Addrs()
		CheckErrExit(err)

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok {
				if !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
	}
	return ""
}
