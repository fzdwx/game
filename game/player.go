package game

// Player 游戏玩家实体
type Player struct {
	ModPlayer *ModPlayer
	ModIcon   *ModIcon
}

// CreateTestPlayer 创建测试玩家
func CreateTestPlayer() *Player {
	player := Player{&ModPlayer{Icon: 0}, &ModIcon{}}
	return &player
}