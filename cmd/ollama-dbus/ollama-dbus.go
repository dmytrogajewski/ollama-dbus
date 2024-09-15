package main

import (
	"fmt"
	"os"

	"github.com/dmytrogajewski/ollama-dbus/internal/app"
	"github.com/dmytrogajewski/ollama-dbus/internal/logger"
)

func main() {
	err := (app.New()).Run(os.Args)

	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%v", err))
	}
}
