package test

import (
	"fmt"
	"testing"
)

var done chan bool

func HelloWord() {
	fmt.Println("hello")
	//time.Sleep(1 * time.Second)
	done <- true
}
func TestChan_Run1(t *testing.T) {
	done = make(chan bool)
	go HelloWord()
	<-done
}

func TestChan_Run2(t *testing.T) {
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
		//<-c
	}

	for i := 0; i < 100; i++ {
		<-c
	}
}

func TestName(t *testing.T) {

	for i := 0; i < 10; i++ {
		go func(a int) {
			defer fmt.Println(6)
		}(i)
	}
}
