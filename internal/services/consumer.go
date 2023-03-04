package services

import (
	"farm/internal/types"
	"math/rand"
	"time"
)

type Consumer struct {
	ConsumeMinInterval int
	ConsumeMaxInterval int
	ConsumeMinAmount   int
	ConsumeMaxAmount   int
}

func NewConsumer(config types.AppConfig) *Consumer {
	return &Consumer{
		ConsumeMaxInterval: config.ConsumerMaxInterval,
		ConsumeMinInterval: config.ConsumerMinInterval,
		ConsumeMaxAmount:   config.ConsumerMaxAmount,
		ConsumeMinAmount:   config.ConsumerMinAmount,
	}
}

func (c *Consumer) StartConsume(target *Refrigerator) {
	for {
		consumeTime := time.Duration(rand.Intn(c.ConsumeMaxInterval-c.ConsumeMinInterval) + c.ConsumeMinInterval)
		consumeAmount := rand.Intn(c.ConsumeMaxAmount-c.ConsumeMinAmount) + c.ConsumeMinAmount

		time.Sleep(time.Second * consumeTime)

		if target.Eggs >= consumeAmount {
			target.Eggs -= consumeAmount
		} else {
			target.Eggs -= target.Eggs
		}

		println(target.Eggs)
	}
}
