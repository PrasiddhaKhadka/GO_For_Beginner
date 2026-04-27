package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Doing task", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// why 1 :-> because a single func is only called
		// would be more like say 3 :-> if a 3 different function were run concurently at once !!!!
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()

	// time.Sleep(time.Second * 2).  // BAD APPROACH !!

}
