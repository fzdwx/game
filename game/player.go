package game

// Player 游戏玩家实体
type Player struct {
	ModPlayer *ModPlayer
	ModIcon   *ModIcon
	ModCard   *ModCard
}

// CreateTestPlayer 创建测试玩家
func CreateTestPlayer() *Player {
	player := Player{&ModPlayer{Icon: 0}, &ModIcon{}, &ModCard{}}
	return &player
}

// RecvSetIcon 用来更新玩家的icon
func (this *Player) RecvSetIcon(iconId int) {
	this.ModPlayer.RecvSetIcon(iconId, this)
}

// RecvSetCard 用来更新玩家的card
func (this *Player) RecvSetCard(cardId int) {
	this.ModPlayer.RecvSetCard(cardId, this)
}
