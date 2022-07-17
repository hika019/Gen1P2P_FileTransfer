package main

import (
	"Gen1P2P_FileTransfer/lib"
	"fmt"
	"os"
)

var shareDir string

var conf lib.ConfigStruct

func init() {
	loadConf()

	//共有dirの作成
	if !lib.IsExists(shareDir) {
		err := os.Mkdir(shareDir, 0777)
		lib.CheckErrExit(err)
	}

}

func main() {
	fmt.Println(conf)
}

func loadConf() {
	conf = *lib.ReadConf()

	shareDir = conf.Share_dir_name
}

func send() {

}

func receive() {

}
