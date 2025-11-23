package main

import (
	"fmt"
	"sync"
)

// type post struct {
// 	views int
// }

// func (p *post) inc(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	p.views++
// }

// func main() {
// 	var wg sync.WaitGroup
// 	myPost := post{views: 0}
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go myPost.inc(&wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("Views:", myPost.views)
// }

// using mutex

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()

	}()
	p.mu.Lock()
	p.views++
}

func main() {
	var wg sync.WaitGroup
	myPost := post{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println("Views:", myPost.views)
}
