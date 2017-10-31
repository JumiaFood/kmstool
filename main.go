package main

import (
	"github.com/JumiaFood/kmstool/src"
	"os"
)

func main() {
	a := kmstool.NewApp()
	a.Run(os.Args)
}
