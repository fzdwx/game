package game

import "game/banwords"

// ModPlayer 用户基础信息实体
type ModPlayer struct {
	// 玩家用户id
	Uid            int64
	Icon           int
	Card           int
	Name           string
	Sign           string
	PlayerLevel    int64
	PlayerExp      int64
	WorldLevel     int64
	WorldLevelCool int64
	Birth          int64
	ShowTeam       int
	ShowCard       []int
	IsProhibit     int
	IsGM           int
}

func (this *ModPlayer) RecvSetIcon(iconId int, player *Player) {
	if !player.modIcon.IsHasIcon(iconId) {
		// todo 没有对应的icon 通知客户端，操作违法
		return
	}

	player.modPlayer.Icon = iconId
}

func (this *ModPlayer) RecvSetCard(cardId int, player *Player) {
	if !player.modCard.IsHasCard(cardId) {
		// todo 没有对应的card 通知客户端，操作违法
		return
	}

	player.modPlayer.Card = cardId
}

func (this *ModPlayer) RecvSetName(name string) {
	if banwords.IsBanWord(name) {
		return
	}
	this.Name = name
}

func (this *ModPlayer) RecvSetSign(sign string) {
	if banwords.IsBanWord(sign) {
		return
	}
	this.Sign = sign
}
