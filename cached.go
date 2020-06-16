package jsoncached

import (
	"io/ioutil"
)

const dir = "./cached/"

type cached struct {
}

func GetFile(fileName string) string {
	file := dir + fileName + ".json"
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	return string(f)
}
