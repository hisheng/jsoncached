package jsoncached

import "errors"

func Get(key string) ([]byte, error) {
	rs := getFile(key)
	if len(rs) == 0 {
		return []byte{}, errors.New("没有文件")
	}
	return rs, nil
}

func GetString(key string) string {
	return getFileString(key)
}
