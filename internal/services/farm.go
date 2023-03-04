package services

type Farm struct {
	Chickens     []Chicken
	Refrigerator *Refrigerator
	eggsChan     chan int
}

func NewFarm(chickens []Chicken) *Farm {
	return &Farm{
		Chickens: chickens,
		eggsChan: make(chan int),
	}
}

func (f *Farm) Start() bool {
	f.Refrigerator = NewRefrigerator(f.eggsChan)

	for _, item := range f.Chickens {
		go item.makeEgg(f.eggsChan)
	}

	go f.Refrigerator.startCollecting()

	return true
}
