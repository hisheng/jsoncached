package jsoncached

import (
	"io/ioutil"
)

const dir = "./cached/"

func getFile(key string) []byte {
	file := getFileName(key)
	return readFile(file)
}

func getFileString(key string) string {
	file := getFileName(key)
	return readFileString(file)
}

func readFile(file string) []byte {
	f, _ := ioutil.ReadFile(file)
	return f
}

func readFileString(file string) string {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	return string(f)
}

func setFile(key, val string) bool {
	fileName := getFileName(key)
	return writeFile(fileName, val)
}

func setFileByte(key string, val []byte) bool {
	fileName := getFileName(key)
	return writeFileByte(fileName, val)
}

func writeFile(fileName, val string) bool {
	var d = []byte(val)
	err := ioutil.WriteFile(fileName, d, 0666)
	if err != nil {
		return false
	}
	return true
}

func writeFileByte(fileName string, val []byte) bool {
	err := ioutil.WriteFile(fileName, val, 0666)
	if err != nil {
		return false
	}
	return true
}

func getFileName(key string) string {
	return dir + key + ".json"
}
