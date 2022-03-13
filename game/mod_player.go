package game

// ModPlayer 用户基础信息实体
type ModPlayer struct {
	// 玩家用户id
	Uid            int64
	Icon           int
	Card           int64
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

// RecvSetIcon 用来更新玩家的icon
func (this *Player) RecvSetIcon(iconId int) {

	if !this.ModIcon.IsHasIcon(iconId) {
		// todo 没有对应的icon 通知客户端，操作违法
		return
	}

	this.ModPlayer.Icon = iconId
}