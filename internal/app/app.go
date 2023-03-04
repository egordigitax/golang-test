package app

import "farm/internal/transport/rest"

func StartApp() {
	h := rest.NewHandler()
	h.Start()
}
