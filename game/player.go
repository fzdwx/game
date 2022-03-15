package game

import (
	"log"
	"time"
)

// Player 游戏玩家实体
type Player struct {
	modPlayer     *ModPlayer
	modIcon       *ModIcon
	modCard       *ModCard
	modUniqueTask *ModUniqueTask
}

// CreateTestPlayer 创建测试玩家
func CreateTestPlayer() *Player {
	player := Player{
		&ModPlayer{Icon: 0},
		&ModIcon{},
		&ModCard{},
		&ModUniqueTask{
			myTaskInfo: map[int]*TaskInfo{},
			//mux:        sync.RWMutex{},
		},
	}
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

// ReduceWorldLevel 降低世界等级
func (this *Player) ReduceWorldLevel() {
	this.modPlayer.reduceWorldLevel()
}

// ReturnWorldLevel 返回世界等级
func (this *Player) ReturnWorldLevel() {
	this.modPlayer.returnWorldLevel()
}

// SetBirth 设置生日
func (this *Player) SetBirth(birth int) {
	this.modPlayer.setBirth(birth)

	if this.modPlayer.isBirthDay() {
		log.Println("今天是你的生日~~~")
	}

}

func (this ModPlayer) isBirthDay() bool {
	now := time.Now()
	month := now.Month()
	day := now.Day()

	return int(month) == this.Birth/100 && day == this.Birth%100
}
