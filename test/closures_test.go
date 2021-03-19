package test

import (
	"fmt"
	"testing"
)

func fib() func() int {

	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func TestRun_Closures(t *testing.T) {

	//var a int = 1
	//var ip *int = &a
	//fmt.Printf("ip的值为：%x", ip)
	//fmt.Println()
	f00 := fib()
	fmt.Println(f00(), f00(), f00(), f00(), f00())
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		a := fibonacci(n - 1)
		b := fibonacci(n - 2)
		res = a + b
	}
	return res
}

func TestRun_Closures1(t *testing.T) {
	result := 0
	//for i := 0; i <= 10; i++ {
		result = fibonacci(4)
		fmt.Printf("fibonacci(%d) is: %d\n", 2, result)
	//}
}
