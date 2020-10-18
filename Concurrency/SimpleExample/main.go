package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	rd      = rand.New(rand.NewSource(time.Now().UnixNano()))
	cache   = map[int]Book{}
	wg      = &sync.WaitGroup{}
	m       = &sync.RWMutex{}
	chCache = make(chan Book)
	chDb    = make(chan Book)
)

func main() {

	for i := 0; i < 10; i++ {

		id := rd.Intn(10) + 1

		wg.Add(2)

		go func(id int, m *sync.RWMutex, wg *sync.WaitGroup, ch chan Book) {

			if b, ok := queryFromCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, m, wg, chCache)

		go func(id int, m *sync.RWMutex, wg *sync.WaitGroup, ch chan Book) {
			if b, ok := queryFromDatabase(id); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, m, wg, chDb)

		go func(chCache, chDb chan Book) {

			select {
			case b := <-chCache:
				fmt.Println("from cache")
				fmt.Println(b)
				<-chDb
			case b := <-chDb:
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(chCache, chDb)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryFromCache(id int, m *sync.RWMutex) (Book, bool) {

	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryFromDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if id == b.ID {
			return b, true
		}
	}
	return Book{}, false
}
