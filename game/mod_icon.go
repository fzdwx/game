package game

type ModIcon struct {
}

// IsHasIcon 查看该玩家是否存在这个icon
func (this ModIcon) IsHasIcon(iconId int) bool {
	return true
}