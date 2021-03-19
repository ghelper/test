package test

import (
	"fmt"
	"testing"
	"time"
)

func TestRun_Slice(t *testing.T) {

	a := []int{1, 2, 3}
	slice := make([]int, 0)
	for _, v := range a {
		slice = append(slice, v)
	}
	fmt.Println(slice)

	//我们创建了一个切片，有四个元素，下标命名为0，1，2，3
	//slice11 := []int{1, 2, 3, 4}
	//假设我们要删除下标为0的元素，这段代码的含义是
	//创建一个空的切片 fmt.Println(slice11[:0])
	//创建一个从下标1到末尾元素的切片 slice11[1:]
	//给空切片添加元素并返回一个新的切片
	//fmt.Println(slice11[2:])
	//slice12 := append(slice11[:0], slice11[1:]...)
	//fmt.Println(slice12)
	//同上我们得出删除下标为i的元素的切片的公式时
	//sliceTemp := append(slice11[:i],slice11[i+1:]...)

	array11 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array11[0:2])
	fmt.Println(array11[3:])
	slice11 := array11[0:3]
	//slice11[0] = 10
	sliceTemp := append(array11[0:2], array11[3:]...)
	//slice11[0] = 11
	fmt.Println(array11, slice11, sliceTemp)
	//输出结果：[11 2 4 5 5] [11 2 4] [11 2 4 5]
}

func TestSlice_Run(t *testing.T) {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		//vtype := v
		go func(in string) {
			fmt.Println(in)
		}(v)
	}
	time.Sleep(3 * time.Second)


	x := []int{0,1,2,3,4}
	fmt.Println(x[1:2])

}
