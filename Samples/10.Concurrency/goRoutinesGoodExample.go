package main

import (
	"fmt"
	"sync"
)

var wgChat sync.WaitGroup

type person4 struct {
	Name       string
	messages   []string
	currentMsg int
}

func (p *person4) getMessage() string {
	var msg string
	if p.currentMsg < len(p.messages) {
		msg = p.Name + ": " + p.messages[p.currentMsg]
		p.currentMsg++
	}
	return msg
}
func goRoutunesGoodExample() {
	msgFromKishore := make(chan string)
	msgFromRishika := make(chan string)
	Rishika := &person4{"Rishika ", []string{"Hi", "How are you?", "What's going on?", "Wow, I too want to learn", "Great,Thanks", "bye"}, 0}
	Kishore := &person4{"Kishore ", []string{"Hi", "I am good", "learning Go", "Contact Romin Irani, he will add you to Go Learning group", "Welcome", "bye"}, 0}
	wgChat.Add(1)
	go writeMessage(Kishore, msgFromRishika, msgFromKishore)
	wgChat.Add(1)
	go writeMessage(Rishika, msgFromKishore, msgFromRishika)
	msgFromRishika <- Rishika.getMessage()
	wgChat.Wait()
}
func writeMessage(p *person4, readChannel chan string, writeChannel chan string) {
	for chanMsg := range readChannel {
		if chanMsg == "" {
			break
		}
		fmt.Println(chanMsg)
		writeChannel <- p.getMessage()
	}
	close(writeChannel)
	wgChat.Done()
}
