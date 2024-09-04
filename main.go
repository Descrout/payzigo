package main

import (
	"log"

	"github.com/Descrout/payzigo/payzigo/utils"
)

func main() {
	//cli := payzigo.WithOptions(&payzigo.PayzigoOptions{})
	log.Println(utils.HashSha1("Hello World"))
}
