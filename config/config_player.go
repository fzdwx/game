package config

import (
	"game/util"
	"log"
)

type playerLevelCfg struct {
	PlayerLevel int
	PlayerExp   int64
	WorldLevel  int
	ChapterId   int
}

var playerLevels []*playerLevelCfg

func init() {
	if err := util.CsvParse("../resource/PlayerLevel.csv", &playerLevels); err != nil {
		panic(err)
	}

	log.Printf("load player level info success. has %d levels", len(playerLevels))
}

// GetLevelConfig 获取对应玩家等级的配置
func GetLevelConfig(level int) *playerLevelCfg {

	if level < 0 || level >= len(playerLevels) {
		return nil
	}

	return playerLevels[level]
}
