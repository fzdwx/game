package config

import "time"

const (
	REDUCE_WORLD_LEVLE_START     = 5                               // 降低世界等级的要求
	REDUCE_WORLD_LEVLE_MAX       = 1                               // 最低能降低多少级别
	REDUCE_WORLD_LEVLE_COOL_TIME = time.Duration(time.Second * 10) // 降低世界等级冷却时间
)
