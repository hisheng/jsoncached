package jsoncached

func Set(key, val string) bool {
	return setFile(key, val)
}

func SetByte(key string, val []byte) bool {
	return setFileByte(key, val)
}
