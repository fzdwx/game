package game

import (
	"game/banwords"
	"game/config"
)

// ModPlayer 用户基础信息实体
type ModPlayer struct {
	// 玩家用户id
	Uid            int64
	Icon           int
	Card           int
	Name           string
	Sign           string
	PlayerLevel    int
	PlayerExp      int64
	WorldLevel     int64
	WorldLevelCool int64
	Birth          int64
	ShowTeam       int
	ShowCard       []int
	IsProhibit     int
	IsGM           int
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
