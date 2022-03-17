package config

import (
	"game/util"
	"log"
)

type cardCfg struct {
	CardId       int
	Friendliness int
	Check        int
}

var cardCfgMap = map[int]*cardCfg{}

func init() {
	var cards []*cardCfg
	if err := util.CsvParse("../resource/Card.csv", &cards); err != nil {
		panic(err)
	}

	for _, card := range cards {
		cardCfgMap[card.CardId] = card
	}

	log.Printf("load card info success. has %d levels", len(playerLevels))
}

func IsCard(cardId int) bool {
	_, ok := cardCfgMap[cardId]
	return ok
}
