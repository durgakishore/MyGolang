package main

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Number of philosophers is simply the length of this list.
var ph = []string{"Mark", "Russell", "Rocky", "Haris", "Root"}

const hunger = 3                // Number of times each philosopher eats
const think = time.Second / 100 // Mean think time
const eat = time.Second / 100   // Mean eat time

var fmtLog = log.New(os.Stdout, "", 0)

var wgDining sync.WaitGroup

func diningProblem(phName string, dominantHand, otherHand *sync.Mutex) {
	fmtLog.Println(phName, "Seated")
	h := fnv.New64a()
	h.Write([]byte(phName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h > 0; h-- {
		fmtLog.Println(phName, "Hungry")
		dominantHand.Lock() // pick up forks
		otherHand.Lock()
		fmtLog.Println(phName, "Eating")
		rSleep(eat)
		dominantHand.Unlock() // put down forks
		otherHand.Unlock()
		fmtLog.Println(phName, "Thinking")
		rSleep(think)
	}
	fmtLog.Println(phName, "Satisfied")
	wgDining.Done()
	fmtLog.Println(phName, "Left the table")
}

func diningPhilosophersProblem() {
	fmtLog.Println("Table empty")
	wgDining.Add(5)
	fork0 := &sync.Mutex{}
	forkLeft := fork0
	for i := 1; i < len(ph); i++ {
		forkRight := &sync.Mutex{}
		go diningProblem(ph[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	go diningProblem(ph[0], fork0, forkLeft)
	wgDining.Wait() // wait for philosphers to finish
	fmtLog.Println("Table empty")
}
