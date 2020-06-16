package jsoncached

import (
	"io/ioutil"
)

const dir = "./cached/"

type cached struct {
}

func GetFile(key string) string {
	file := GetFileName(key)
	return ReadFile(file)
}

func ReadFile(file string) string {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	return string(f)
}

func SetFile(key, val string) bool {
	fileName := GetFileName(key)
	return WriteFile(fileName, val)
}

func WriteFile(fileName, val string) bool {
	var d = []byte(val)
	err := ioutil.WriteFile(fileName, d, 0666)
	if err != nil {
		return false
	}
	return true
}

func GetFileName(key string) string {
	return dir + key + ".json"
}
