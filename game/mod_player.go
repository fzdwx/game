package game

import (
	"game/banwords"
	"game/config"
	"log"
	"time"
)

// ModPlayer 用户基础信息实体
type ModPlayer struct {
	// 玩家用户id
	Uid            int64  // uid
	Icon           int    // 头像
	Card           int    // 名片
	Name           string // 名字
	Sign           string // 前面
	PlayerLevel    int    // 等级
	PlayerExp      int64  // 经验
	WorldLevel     int    // 大世界等级
	WorldLevelNow  int    // 大世界等级(实际)
	WorldLevelCool int64  // 操作大世界等级的冷却时间
	Birth          int    // 生日
	ShowTeam       []int  // 展示阵容
	ShowCard       []int  // 展示名片
	IsProhibit     int    //封禁状态
	IsGM           int    // GM 标志
}

func (this *ModPlayer) recvSetIcon(iconId int, player *Player) {
	if !player.modIcon.isHasIcon(iconId) {
		// todo 没有对应的icon 通知客户端，操作违法
		return
	}

	player.modPlayer.Icon = iconId
}

func (this *ModPlayer) recvSetCard(cardId int, player *Player) {
	if !player.modCard.isHasCard(cardId) {
		// todo 没有对应的card 通知客户端，操作违法
		return
	}

	player.modPlayer.Card = cardId
}

func (this *ModPlayer) recvSetName(name string) {
	if banwords.IsBanWord(name) {
		return
	}
	this.Name = name
}

func (this *ModPlayer) recvSetSign(sign string) {
	if banwords.IsBanWord(sign) {
		return
	}
	this.Sign = sign
}

func (this *ModPlayer) addExp(exp int64, player *Player) {
	this.PlayerExp += exp

	for {
		levelConfig := config.GetLevelConfig(this.PlayerLevel)
		if levelConfig == nil {
			break
		}

		if levelConfig.PlayerExp == 0 {
			break
		}

		// 是否完成任务
		if levelConfig.ChapterId > 0 && !player.modUniqueTask.IsFinish(levelConfig.ChapterId) {
			break
		}

		if this.PlayerExp >= levelConfig.PlayerExp {
			this.PlayerLevel++
			this.PlayerExp -= levelConfig.PlayerExp
		} else {
			break
		}

	}
}

func (this *ModPlayer) reduceWorldLevel() {
	if this.WorldLevel < config.REDUCE_WORLD_LEVLE_START {
		log.Println("降低世界等级失败. --- 当前世界等级:", this.WorldLevel)
		return
	}

	if this.WorldLevel-this.WorldLevelNow >= config.REDUCE_WORLD_LEVLE_MAX {
		log.Println("降低世界等级失败,已经降低过. --- 当前世界等级:", this.WorldLevel, "真实等级:", this.WorldLevelNow)
	}

	if time.Now().Unix() < this.WorldLevelCool {
		log.Println("降低世界等级失败,还在冷却中")
		return
	}

	this.WorldLevelNow -= 1
	this.WorldLevelCool = time.Now().Unix() + int64(config.REDUCE_WORLD_LEVLE_COOL_TIME.Seconds())
}

func (this *ModPlayer) returnWorldLevel() {
	if this.WorldLevel-this.WorldLevelNow >= config.REDUCE_WORLD_LEVLE_MAX {
		log.Println("返回世界等级失败,已经降低过. --- 当前世界等级:", this.WorldLevel, "真实等级:", this.WorldLevelNow)
	}

	if time.Now().Unix() < this.WorldLevelCool {
		log.Println("返回世界等级失败,还在冷却中")
		return
	}

	this.WorldLevelNow += 1
	this.WorldLevelCool = time.Now().Unix() + int64(config.REDUCE_WORLD_LEVLE_COOL_TIME.Seconds())
}

func (this *ModPlayer) setBirth(birth int) {
	if this.Birth > 0 {
		log.Println("已经设置过生日了")
		return
	}

	month := birth / 100
	day := birth % 100
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day <= 0 || day > 31 {
			log.Println(month, "月 没有", day, "日!")
			return
		}
	case 4, 6, 9, 11:
		if day <= 0 || day > 30 {
			log.Println(month, "月 没有", day, "日!")
			return
		}
	case 2:
		if day <= 0 || day > 29 {
			log.Println(month, "月 没有", day, "日!")
			return
		}
	default:
		log.Println("没有", month, "月")
		return
	}
	this.Birth = birth
}

func (this *ModPlayer) setShowCard(card []int, player *Player) {
	i := len(card)
	if i > 9 {
		return
	}

	cardExist := make(map[int]int)
	var newCard []int
	for i, c := range card {
		if _, ok := cardExist[c]; ok {
			continue
		}

		if !player.modCard.isHasCard(c) {
			continue
		}
		cardExist[c] = i
		newCard = append(newCard, c)
	}

	this.ShowCard = newCard
}
