package main

import (
	"command-line/app"
	"log"
	"os"
)

func main() {
	application := app.Gerar()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}