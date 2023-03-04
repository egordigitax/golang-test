package services

import (
	"farm/internal/scripts"
	"strconv"
)

type Refrigerator struct {
	Eggs     int
	eggsChan chan int
}

func NewRefrigerator(eggsChan chan int) *Refrigerator {
	return &Refrigerator{0, eggsChan}
}

func (r *Refrigerator) startCollecting() {
	debugMode := scripts.GetDefaultConfig().DebugMode

	for {
		a := <-r.eggsChan
		r.Eggs += a

		if debugMode == true {
			println("Eggs: " + strconv.Itoa(r.Eggs))
		}
	}
}
