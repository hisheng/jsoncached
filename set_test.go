package jsoncached

import (
	"encoding/json"
	"testing"
)

func TestSet(t *testing.T) {
	SetString("name2", "hisheng1112222221")

	s, _ := json.Marshal([]string{"sss", "133", "海生"})
	SetByte("name21", s)
	t.Log(Get("name21"))
}
