package app

import (
	"farm/internal/scripts"
	"farm/internal/services"
	"farm/internal/transport/rest"
)

func StartApp() {
	config := scripts.GetDefaultConfig()

	chickens := services.NewChickens(5, config)
	consumer := services.NewConsumer(config)

	f := services.NewFarm(chickens, consumer)
	f.Start()

	h := rest.NewHandler("8000")
	h.R = f.Refrigerator

	h.Start()
}
