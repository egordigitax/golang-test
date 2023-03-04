package services

import "math/rand"

type Chicken struct {
}

func (c *Chicken) makeEgg(minAmount int, maxAmount int) int {
	return rand.Intn(maxAmount-minAmount) + minAmount
}
