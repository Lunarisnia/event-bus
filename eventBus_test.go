package ebus

import (
	"fmt"
	"testing"
)

func Test_EventBus(t *testing.T) {
	bus, err := NewEBus()
	if err != nil {
		t.FailNow()
	}
	testSignal, err := bus.CreateSignal("Test")
	if err != nil {
		t.FailNow()
	}
	onTest := func() {
		fmt.Println("Succeeded")
	}
	onTest2 := func() {
		fmt.Println("Again Succeeded")
	}

	//fmt.Println(testSignal)
	testSignal.Connect(onTest)
	testSignal.Connect(onTest2)
	testSignal.Emit()
}
