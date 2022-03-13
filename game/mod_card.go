package game

type ModCard struct {
}

// IsHasCard 查看该玩家是否存在这个card
func (this ModCard) IsHasCard(iconId int) bool {
	return true
}
