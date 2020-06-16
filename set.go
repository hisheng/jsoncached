package jsoncached

func Set(key, val string) bool {
	return SetFile(key, val)
}
