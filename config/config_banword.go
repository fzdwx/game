package config

import "fmt"

// banWordCfg ConfigBanWord 是违禁词的配置
type banWordCfg struct {
	Id  int
	Txt string
}

var words []*banWordCfg

func init() {

	words = append(words,
		&banWordCfg{Id: 0, Txt: "外挂"},
		&banWordCfg{Id: 1, Txt: "工具"},
		&banWordCfg{Id: 2, Txt: "fzdwx"},
	)

	fmt.Println("csv banWord init")
}

func GetBaseWords() []string {
	baseWords := make([]string, 0)
	for _, v := range words {
		baseWords = append(baseWords, v.Txt)
	}
	return baseWords
}
