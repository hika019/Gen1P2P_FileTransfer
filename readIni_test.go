package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigStruct struct {
	max_user   int
	use_crypto bool
}

var cfg ConfigStruct

func init() {
	loadCfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("Fail to read file: ", err.Error())
		os.Exit(1)
	}

	cfg = ConfigStruct{
		max_user:   loadCfg.Section("receive").Key("max_user").MustInt(),
		use_crypto: loadCfg.Section("send").Key("use_crypto").MustBool(),
	}
}

func main() {
	fmt.Println(cfg.max_user)
	fmt.Println(cfg.use_crypto)
}
