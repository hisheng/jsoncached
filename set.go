package jsoncached

func Set(key, val string) bool {
	return SetFile(key, val)
}

func SetByte(key string, val []byte) error {
	return SetFileByte(key, val)
}
