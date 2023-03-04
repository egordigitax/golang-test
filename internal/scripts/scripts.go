package scripts

import "farm/internal/types"

func GetDefaultConfig() types.AppConfig {
	return types.AppConfig{
		ChickenMinEgg:      1,
		ChickenMaxEgg:      4,
		ChickenMinInterval: 4,
		ChickenMaxInterval: 10}
}
