package main

import (
	"log"

	"github.com/daichi-sato-design/gae-test/handlers"
)

func main(){
	log.Println("アプリの起動")
	handlers.Handlers()
}