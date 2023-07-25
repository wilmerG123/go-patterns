package main

import (
	"fmt"
	"go-patterns/1-singleton/client_1"
	"go-patterns/1-singleton/client_2"
	"go-patterns/1-singleton/singleton"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client_1.IncrementId()
		}()
		go func() {
			defer wg.Done()
			client_2.IncrementId()
		}()
	}
	wg.Wait()
	e := singleton.GetInstance()
	id := e.GetId()
	fmt.Println(id)

}
