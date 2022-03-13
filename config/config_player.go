package config

type playerLevelCfg struct {
	PlayerLevel int
	PlayerExp   int64
	WorldLevel  int
	ChapterId   int
}

var playerLevels []*playerLevelCfg

func init() {
	// todo load config
}

// GetLevelConfig 获取对应玩家等级的配置
func GetLevelConfig(level int) *playerLevelCfg {

	if level < 0 || level >= len(playerLevels) {
		return nil
	}

	return playerLevels[level]
}
