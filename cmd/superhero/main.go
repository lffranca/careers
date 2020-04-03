package main

import (
	"sync"

	"github.com/lffranca/careers/api"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go api.InitAPI(wg)

	wg.Wait()
}
