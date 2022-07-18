package main

import (
	"Gen1P2P_FileTransfer/lib"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var shareDir string
var hashTable map[[32]byte]string

var conf lib.ConfigStruct

func init() {
	loadConf()

	//共有dirの作成
	if !lib.IsExists(shareDir) {
		err := os.Mkdir(shareDir, 0777)
		lib.CheckErrExit(err)
	}

	hashTable = map[[32]byte]string{}

	lib.UpFHT(conf, createFileHashTable(shareDir))

}

func main() {
	fmt.Println(conf)
}

func loadConf() {
	conf = *lib.ReadConf()

	shareDir = conf.Share_dir_name
}

func loadLocalFiles(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	lib.CheckErrExit(err)

	var paths []string

	for _, file := range files {
		if file.IsDir() {
			//paths = append(paths, loadLocalFiles(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join("", file.Name()))
	}
	return paths
}

func createFileHashTable(dir string) []byte {

	paths := loadLocalFiles(dir)
	for _, path := range paths {
		h := sha256.Sum256([]byte(path))
		hashTable[h] = path
	}

	var hashs []byte
	for hash := range hashTable {
		for i := 0; i < 32; i++ {
			hashs = append(hashs, hash[i])
		}
	}

	return hashs
}

func uploadFile() {

}

func downloadFile() {

}
