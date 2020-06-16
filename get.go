package jsoncached

func Get(key string) []byte {
	return getFile(key)
}

func GetString(key string) string {
	return getFileString(key)
}
