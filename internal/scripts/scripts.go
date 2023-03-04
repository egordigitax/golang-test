package scripts

import "farm/internal/types"

func GetDefaultConfig() types.AppConfig {
	return types.AppConfig{
		ChickenMinEgg:      2,
		ChickenMaxEgg:      5,
		ChickenMinInterval: 2,
		ChickenMaxInterval: 10,

		ConsumerMaxAmount:   20,
		ConsumerMinAmount:   10,
		ConsumerMaxInterval: 15,
		ConsumerMinInterval: 10,
	}
}
