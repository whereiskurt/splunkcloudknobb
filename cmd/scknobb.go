package main

import (
	"github.com/whereiskurt/cloudcrawler/internal/app"
	"github.com/whereiskurt/cloudcrawler/pkg/config"
)

func main() {
	config := config.NewConfig()

	a := app.NewApp(config)

	a.InvokeCLI()

	return
}
