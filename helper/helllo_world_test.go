package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("world")

	if result != "helloworld" {
		t.Error("result not hello world")
		// t.Fail()
		// panic("Test failed: expected helloworld")
	}
	fmt.Println("Test passed: expected hello world")
}
func TestHelloTayo(t *testing.T) {
	result := HelloTayo("tayo")

	if result != "hellotayo" {
		t.Fatal("result not hello tayo")
		// t,FailNow()
		// panic("Test failed: expected hellotayo")
	}

	fmt.Println("Test passed: expected hellotayo")
}
