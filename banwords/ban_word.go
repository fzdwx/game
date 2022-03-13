package banwords

import (
	"fmt"
	"game/config"
	"regexp"
	"time"
)

var banWord = Of(
	config.GetBaseWords(),
	[]string{},
)

// BanWord 是一个违禁词匹配工具
type BanWord struct {
	banWordBase  []string // 配置生成
	banWordExtra []string // 动态更新
}

// IsBanWord 是用来检查 txt 是否为违禁词
func IsBanWord(txt string) bool {
	return banWord.IsBanWord(txt)
}

// Instance get BandWord 单例
func Instance() *BanWord {
	return banWord
}

// Of 是一个创建 BanWord 的工厂方法
func Of(banWordBase []string, banWordExtra []string) *BanWord {
	return &BanWord{banWordBase, banWordExtra}
}

// IsBanWord  查看txt是否为一个违禁词
func (this *BanWord) IsBanWord(txt string) bool {

	for _, v := range this.banWordBase {
		match, _ := regexp.MatchString(v, txt)
		if match {
			return match
		}
	}

	for _, v := range this.banWordExtra {
		match, _ := regexp.MatchString(v, txt)
		if match {
			return match
		}
	}

	return false
}

// RunRefreshTask 启动更新词库
func (this *BanWord) RunRefreshTask() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now().Unix())
		}
	}
}
