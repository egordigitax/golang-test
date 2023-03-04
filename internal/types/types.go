package types

type AppConfig struct {
	ChickenMinEgg      int
	ChickenMaxEgg      int
	ChickenMinInterval int
	ChickenMaxInterval int

	ConsumerMinInterval int
	ConsumerMaxInterval int
	ConsumerMinAmount   int
	ConsumerMaxAmount   int

	DebugMode bool
}
