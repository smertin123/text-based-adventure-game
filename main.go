package main

import "sync"

var wg sync.WaitGroup

func main() {
	//choose a character
	Introduce()
	wg.Wait()
}
