package app

import (
	"farm/internal/scripts"
	"farm/internal/services"
	"farm/internal/transport/rest"
)

func StartApp() {
	config := scripts.GetDefaultConfig()
	chickens := services.NewChickens(5, config)
	f := services.NewFarm(chickens)
	f.Start()
	h := rest.NewHandler()
	h.R = f.Refrigerator
	h.Start()
}
