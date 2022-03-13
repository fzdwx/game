package banwords

import (
	"testing"
)

func TestIsBanWord(t *testing.T) {

	if !(IsBanWord("我操你的fzdwx")) {
		t.Errorf("违禁词匹配失败 fzdwx")
	}

	if IsBanWord("我操") {
		t.Errorf("违禁词匹配失败 我操")
	}

	if !IsBanWord("外挂") {
		t.Errorf("违禁词匹配失败 外挂")
	}
}
