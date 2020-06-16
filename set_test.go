package jsoncached

import "testing"

func TestSet(t *testing.T) {
	Set("name2", "hisheng1112222221")
	t.Log(Get("name2"))
}
