package game

// Player 游戏玩家实体
type Player struct {
	modPlayer *ModPlayer
	modIcon   *ModIcon
	modCard   *ModCard
}

// CreateTestPlayer 创建测试玩家
func CreateTestPlayer() *Player {
	player := Player{&ModPlayer{Icon: 0}, &ModIcon{}, &ModCard{}}
	return &player
}

// RecvSetIcon 用来更新玩家的icon
func (this *Player) RecvSetIcon(iconId int) {
	this.modPlayer.recvSetIcon(iconId, this)
}

// RecvSetCard 用来更新玩家的card
func (this *Player) RecvSetCard(cardId int) {
	this.modPlayer.recvSetCard(cardId, this)
}

// RecvSetName 用来更新玩家的name
func (this *Player) RecvSetName(name string) {
	this.modPlayer.recvSetName(name)
}

// RecvSetSign 用来更新玩家的sign
func (this *Player) RecvSetSign(name string) {
	this.modPlayer.recvSetSign(name)
}
