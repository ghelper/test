package test

import (
	"fmt"
	"testing"
)

//var nickname string

//func TestRun(t *testing.T) {
//	nickname = "大虾"
//	print(nickname) //大虾
//	change()
//	println(nickname) //大虾
//}
//
//func change() {
//	nickname := "小虾"
//	print(nickname) //小虾
//	show()          //大虾 上面只是nickname := 只是初始化声明变量。仅在函数内有效，全局变量nickname未变
//}
//
//func show() {
//	print(nickname)
//}

var nickname = "大虾"

func TestRun(t *testing.T) {
	show()              //大虾
	change()            //小虾
	show()              //大虾
	fmt.Print(nickname) //大虾
}
func show() { print(nickname) }

func change() {
	//nickname := "小虾"
	nickname := "小虾"
	print(nickname)
}

var ap = [5]int{3: 2} //下标3的值赋值为2
var at = [6]int{}

func TestRun1(t *testing.T) {
	//数组遍历,i是数组当前下标,p是当前下标对应的值,i可以用_下划线代替,表示忽略
	as := [3]int{1, 2, 3}
	var pa [3]*int
	for i, p := range as {
		fmt.Println(i, "---", p)
		pa[i] = &as[i]
	}

	//还可以常用的遍历循环
	le := len(as)
	for i := 0; i < le; i++ {
		fmt.Println(as[i])
	}

	a := []int{1, 2, 3}
	b := []int{2, 3, 4, 5, 6}
	a = append(a, b...)
	fmt.Println(a)
}

func TestRun2(t *testing.T) {
	// 每一个变量在运行时都有一个地址（这个地址存储着这个变量对应的值）
	var temp int = 1
	// go语言获取地址符， temp 变量，通过&符号来获取temp这个变量在内存中的地址，我们将获取到的地址赋值给一个指针变量 pointTemp ,类型为 *int
	var pointTemp *int = &temp
	fmt.Println(pointTemp)
	fmt.Println(*pointTemp)
	//输出结果0xc00005a2d01
	// 思考一下这行代码输出什么
	fmt.Println(*&temp)
}

func TestRun3(t *testing.T) {
	a := 5
	b := 10
	// swap(a, b)
	swap(&a, &b) // &a，&b表示引用a,b的内存地址
	// a, b = swap(a, b)
	// b, a = a, b

	fmt.Println("a=", a)
	fmt.Println("b=", b)
}

// a *int, b *int表示接受a,b的内存地址
func swap(a *int, b *int) {
	// *a，*b表示a,b内存地址所指向的值
	tmp := *a
	*a = *b
	*b = tmp
}


func TestRun4(t *testing.T) {
	a := 3
	b := 4
	c := "abc"
	d := [3]int{1,2,3}
	fmt.Printf("main方法：a的值为 %v,b的值为 %v,c的值为 %v,d的值为 %v \n",a,b,c,d)
	demo(a,b,c,d)
	fmt.Printf("main方法：a的值为 %v,b的值为 %v,c的值为 %v,d的值为 %v \n",a,b,c,d)
}

func demo(a,b int,c string,d [3]int) {
	a = 5
	b = 6
	c = "efg"
	d[0] = 0
	fmt.Printf("demo函数：a的值为 %v,b的值为 %v,c的值为 %v,d的值为 %v\n",a,b,c,d)
}
