package jsoncached

func Set(key, val string) bool {
	return SetFile(key, val)
}

func SetByte(key string, val []byte) bool {
	return SetFileByte(key, val)
}
