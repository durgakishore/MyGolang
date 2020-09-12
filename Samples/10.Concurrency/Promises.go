package main

import (
	"errors"
	"fmt"
)

type promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

type purchaseOrder1 struct {
	Number int
	Value  float64
}

func savePO1(po *purchaseOrder1, shouldFail bool) *promise {

	result := new(promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		if shouldFail {
			result.failureChannel <- errors.New("failed to save purchase order")
		} else {
			po.Number = 1234
			result.successChannel <- po
		}

	}()
	return result
}

func (p *promise) then(success func(interface{}) error, failure func(error)) *promise {

	result := new(promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {

		select {

		case obj := <-p.successChannel:
			newErr := success(obj)
			if newErr == nil {
				result.successChannel <- obj
			} else {
				result.failureChannel <- newErr
			}
		case err := <-p.failureChannel:
			failure(err)
			result.failureChannel <- err
		}
	}()
	return result
}

func samplePromise() {

	po := new(purchaseOrder1)
	po.Value = 42.57
	savePO1(po, false).then(func(obj interface{}) error {
		po := obj.(*purchaseOrder1)
		fmt.Println("Purchase order saved with ID: ", po.Number)
		return nil
	}, func(err error) {
		fmt.Println("failed to save purchase order " + err.Error())
	})

	fmt.Scanln()
}
