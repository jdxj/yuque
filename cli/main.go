package main

import (
	"fmt"
	"os"

	"github.com/jdxj/yuque/cli/cmd/app"
)

func main() {
	app := app.NewApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("%s\n", err)
	}
}
