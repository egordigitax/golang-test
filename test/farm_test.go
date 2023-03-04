package test

import (
	"farm/internal/scripts"
	"farm/internal/services"
	"testing"
)

func TestChickensAmount(t *testing.T) {
	got := len(services.NewChickens(5, scripts.GetDefaultConfig()))
	want := 5

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCreateFarm(t *testing.T) {
	chickens := services.NewChickens(5, scripts.GetDefaultConfig())
	consumer := services.NewConsumer(scripts.GetDefaultConfig())
	farm := services.NewFarm(chickens, consumer)
	got := farm.Start()

	if got != true {
		t.Errorf("Service can't start")
	}
}
