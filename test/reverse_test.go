package test

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestRun_reverse(t *testing.T) {

	var num1 int
	var num2 int
	num1 = 38
	num2 = 39
	num1, num2 = num2, num1
	fmt.Println(num1, num2)
	fmt.Println("------------------------------------------")
	src := "猪猪是个笨！"
	fmt.Println(len(src))
	dst := reverse([]rune(src))
	fmt.Printf("%v\n", string(dst))

	//2代表精度，这种方式会有小数点后无效的0的情况
	//fmt.Println(fmt.Sprintf("%.2f", 123.126))
	//// g可以去掉小数点后无效的0
	//fmt.Println(fmt.Sprintf("%g", 123.603))

	fmt.Println(math.Max(568.6, 256.9))
	s := FormatFloat(5.0265, 2)
	fmt.Println(s)
}

func reverse(s []rune) []rune {
	fmt.Println(len(s))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// 主要逻辑就是先乘，trunc之后再除回去，就达到了保留N位小数的效果
func FormatFloat(num float64, decimal int) string {

	//// 2代表精度，这种方式会有小数点后无效的0的情况
	//strconv.FormatFloat(123.123 'f', 2, 64)
	//// 效果同上
	//fmt.Sprintf("%.2f", 123.123)
	//// g可以去掉小数点后无效的0
	//fmt.Sprintf("%g", 123.00)
	//// 效果同上，可以去掉0，但是达不到保留指定位数的效果
	//strconv.FormatFloat(a, 'g', -1, 64)

	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	return strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)

}

func TestString_Run(t *testing.T) {
	var str1 = "hello "
	//var str2 = "world"
	//fmt.Println(str1 + str2)
	fmt.Scanf("%s", &str1)
	fmt.Println(str1)
}
