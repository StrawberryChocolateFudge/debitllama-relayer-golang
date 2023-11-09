package main

import (
	"log"
	"os"

	cmd "debitllama.com/relayer/cmd"
)

func main() {
	app := cmd.GetApp()
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
