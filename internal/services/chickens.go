package services

import (
	"farm/internal/types"
	"math/rand"
	"time"
)

type Chicken struct {
	EggsMinAmount   int
	EggsMaxAmount   int
	EggsMinInterval int
	EggsMaxInterval int
}

func NewChickens(count int, config types.AppConfig) []*Chicken {
	var arr []*Chicken
	for i := 0; i < count; i++ {
		arr = append(arr, &Chicken{
			EggsMinAmount:   config.ChickenMinEgg,
			EggsMaxAmount:   config.ChickenMaxEgg,
			EggsMinInterval: config.ChickenMinInterval,
			EggsMaxInterval: config.ChickenMaxInterval,
		})
	}
	return arr
}

func (c *Chicken) makeEgg(eggsChan chan int) {
	for {
		eggTime := time.Duration(rand.Intn(c.EggsMaxInterval-c.EggsMinInterval) + c.EggsMinInterval)
		time.Sleep(time.Second * eggTime)
		eggsChan <- rand.Intn(c.EggsMaxAmount-c.EggsMinAmount) + c.EggsMinAmount
	}
}
