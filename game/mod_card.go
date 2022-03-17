package game

import "game/config"

type ModCard struct {
}

// isHasCard 查看该玩家是否存在这个card
func (this ModCard) isHasCard(cardId int) bool {
	return config.IsCard(cardId)
}
