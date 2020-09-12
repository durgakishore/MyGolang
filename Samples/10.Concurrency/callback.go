package main

import (
	"fmt"
)

type purchaseOrder struct {
	Number int
	Value  float64
}

func savePO(po *purchaseOrder, callback chan *purchaseOrder) {

	po.Number = 1234
	callback <- po
}

func sampleCallback() {

	po := new(purchaseOrder)
	po.Value = 45.50

	ch := make(chan *purchaseOrder)

	go savePO(po, ch)

	newPO := <-ch

	fmt.Println("Number", newPO.Number, "Value", newPO.Value)
}
