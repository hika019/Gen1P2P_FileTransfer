package lib

import "gopkg.in/ini.v1"

type ConfigStruct struct {
	Share_dir_name   string
	Max_connect_user int
	Use_crypto       bool
	Use_port         int
	Server_addr      string
	Fht_port         string
}

const confPath string = "config.ini"

func loadConf() *ini.File {
	loadConf, err := ini.Load(confPath)
	CheckErrExit(err)
	return loadConf
}

func ReadConf() *ConfigStruct {
	loadConf := loadConf()
	conf := ConfigStruct{
		Share_dir_name:   loadConf.Section("setting").Key("share_dir_name").String(),
		Max_connect_user: loadConf.Section("setting").Key("max_connect_user").MustInt(),
		Use_crypto:       loadConf.Section("setting").Key("use_crypto").MustBool(),
		Use_port:         loadConf.Section("setting").Key("use_port").MustInt(),
		Server_addr:      loadConf.Section("server").Key("server_addr").String(),
		Fht_port:         loadConf.Section("server").Key("fhb_port").String(),
	}

	return &conf
}
