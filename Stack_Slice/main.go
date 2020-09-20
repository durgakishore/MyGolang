package main

import (
	"fmt"
	"sync"
)

//ItemType interface represnts the type of the item
type ItemType interface{}

type stack struct {
	items  []ItemType
	rwLock sync.RWMutex
}

func (s *stack) Push(t ItemType) {

	if s.items == nil {
		s.items = []ItemType{}
	}

	s.rwLock.Lock()
	s.items = append(s.items, t)
	s.rwLock.Unlock()
}

func (s *stack) Pop() *ItemType {

	if s.items == nil {
		return nil
	}

	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	item := s.items[len(s.items)-1]

	s.items = s.items[0 : len(s.items)-1]

	return &item
}

func (s *stack) GetAllItems() []ItemType {

	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	return s.items
}

func main() {
	s := &stack{}
	s.Push(1)
	s.Push(2)
	s.Push("gRPC")
	s.Push("GoLang")
	fmt.Println(s.GetAllItems())
	s.Pop()
	fmt.Println(s.GetAllItems())
}
