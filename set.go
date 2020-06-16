package jsoncached

func Set(key string, val []byte) bool {
	return setFileByte(key, val)
}

func SetByte(key string, val []byte) bool {
	return setFileByte(key, val)
}

func SetString(key string, val string) bool {
	return setFile(key, val)
}
