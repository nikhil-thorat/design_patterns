package main

import (
	"fmt"
	"sync"
)

type single struct{}

var (
	lock           = &sync.Mutex{}
	singleInstance *single
)

func getInstance() *single {

	if singleInstance != nil {
		fmt.Println("SINGLE INSTANCE ALREADY CREATED")
		return singleInstance
	}

	lock.Lock()

	defer lock.Unlock()

	if singleInstance != nil {
		fmt.Println("SINGLETON INSTANCE ALREADY CREATED")
		return singleInstance
	}

	fmt.Println("CREATING SINGLETON INSTANCE")
	singleInstance = &single{}

	return singleInstance

}

func main() {

	for i := 0; i < 10; i++ {
		go getInstance()
	}

	fmt.Scanln()

}
