package main

import (
	"github.com/whereiskurt/splunkcloudknobb/internal/app"
	"github.com/whereiskurt/splunkcloudknobb/pkg/config"
)

func main() {
	config := config.NewConfig()

	a := app.NewApp(config)

	a.InvokeCLI()

	return
}
