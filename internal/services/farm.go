package services

type Farm struct {
	Chickens     []*Chicken
	Refrigerator *Refrigerator
	Consumer     *Consumer
	eggsChan     chan int
}

func NewFarm(chickens []*Chicken, consumer *Consumer) *Farm {
	return &Farm{
		Chickens: chickens,
		Consumer: consumer,
		eggsChan: make(chan int),
	}
}

func (f *Farm) Start() bool {
	f.Refrigerator = NewRefrigerator(f.eggsChan)

	for _, item := range f.Chickens {
		go item.makeEgg(f.eggsChan)
	}

	go f.Refrigerator.startCollecting()
	go f.Consumer.StartConsume(f.Refrigerator)

	return true
}
