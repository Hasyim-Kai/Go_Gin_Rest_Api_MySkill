package practice_goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

func doPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func RunGoroutines() {
	runtime.GOMAXPROCS(8) // penggunaan processor dalam cpu kita
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("data %d", i)
		wg.Add(1)
		go doPrint(&wg, data)
	}
	wg.Wait()
}
