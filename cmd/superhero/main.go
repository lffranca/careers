package main

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/lffranca/careers/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Not loading .env file")
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go api.InitAPI(wg)

	wg.Wait()
}
