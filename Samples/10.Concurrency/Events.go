package main

import "fmt"

//Button structure
type Button struct {
	eventListeners map[string][]chan string
}

//MakeButton function returns *Button
func MakeButton() *Button {

	result := new(Button)
	result.eventListeners = make(map[string][]chan string)
	return result
}

//AddEventListener adds Button events
func (button *Button) AddEventListener(event string, responseChannel chan string) {

	if _, present := button.eventListeners[event]; present {
		button.eventListeners[event] = append(button.eventListeners[event], responseChannel)
	} else {
		button.eventListeners[event] = []chan string{responseChannel}
	}
}

// RemoveEventListener removes the event
func (button *Button) RemoveEventListener(event string, listenerChannel chan string) {

	if _, present := button.eventListeners[event]; present {
		for index, _ := range button.eventListeners[event] {
			if button.eventListeners[event][index] == listenerChannel {
				button.eventListeners[event] = append(button.eventListeners[event][:index],
					button.eventListeners[event][index+1:]...)
				break
			}
		}
	}
}

//TriggerEvent will trigger a new event
func (button *Button) TriggerEvent(event string, response string) {
	if _, present := button.eventListeners[event]; present {
		for _, handler := range button.eventListeners[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}

//SampleEvent will explain event mechanism
func SampleEvent() {

	btn := MakeButton()
	handlerOne := make(chan string)
	handlerTwo := make(chan string)
	btn.AddEventListener("click", handlerOne)
	btn.AddEventListener("click", handlerTwo)

	go func() {

		for {
			msg := <-handlerOne
			fmt.Println("Handler one " + msg)
		}
	}()

	go func() {

		for {
			msg := <-handlerTwo
			fmt.Println("Handler two " + msg)
		}
	}()

	btn.TriggerEvent("click","Button clicked")
	btn.RemoveEventListener("click",handlerOne)
 	btn.TriggerEvent("click","Button clicked again")

	fmt.Scanln()
}
