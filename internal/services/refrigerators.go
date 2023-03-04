package services

type Refrigerator struct {
	Eggs     int
	eggsChan chan int
}

func NewRefrigerator(eggsChan chan int) *Refrigerator {
	return &Refrigerator{0, eggsChan}
}

func (r *Refrigerator) startCollecting() {
	for {
		a := <-r.eggsChan
		r.Eggs += a
		println(r.Eggs)
	}
}
