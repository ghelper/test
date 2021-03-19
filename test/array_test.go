package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"testing"
)

var c = make(chan int)
var d = make(chan<- int)

func TestArray_Run1(t *testing.T) {
	fmt.Println("hello,world!")
	//fmt.Println(len(a))
	d = c         // 不能写成 c = d
	go p(d, 1000) // 虽然传给了d,但是channel按引用传递,值也传递给c了
	fmt.Println(<-c)
}
func p(ch chan<- int, i int) {
	ch <- i
}

func TestMap_Run(t *testing.T) {
	map1 := make(map[string]interface{})
	map1["slogan1"] = "hello 世界"
	if a, ok := map1["slogan"]; ok {
		fmt.Println(a)
		fmt.Println("存在")
	} else {
		fmt.Println(a)
		fmt.Println("不存在")
	}
}

func Test_Run2(t *testing.T) {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(wg *sync.WaitGroup, url1 string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			//http.Get(url)
			fmt.Println(url1)
		}(&wg, url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

func Test_Run3(t *testing.T) {
	an := "这是我的第%s个程序"
	content := fmt.Sprintf(an, "5855")
	fmt.Println(content)
}

type Person struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func Test_Run4(t *testing.T) {

	p := Person{Name: "wsq", Gender: "male"}
	// 序列化
	b, _ := json.Marshal(p)
	fmt.Println(string(b)) //{"Name":"a","Gender":"male","Age":23}

	// 反序列化
	pp := Person{}
	err := json.Unmarshal(b, &pp)
	if err != nil {
		//errors.New("unmarshal error")
		fmt.Println("unmarshal error")
	}
	fmt.Printf("%T, %v\n", pp, pp)
}

func Test_Run5(t *testing.T) {

	//只有英文的情况下
	//s1 := "hello world"
	//var s1_byte_list = []byte(s1) //打散成字符列表
	//s1_byte_list[5] = 'F'         //修改下表为6的字符为F
	//s1 = string(s1_byte_list)     //打散的字符列表在组装成字符串
	//fmt.Println(s1)               //输出 hello 6orld
	//
	////带有中文的情况
	//s2 := "天空好像下雨，我好像住你隔壁vay"
	//var s2_rune_list = []rune(s2) //打散成utf-8字符列表
	//s2_rune_list[5] = '雪'         //修改下表为5的字符为雪
	//s2 = string(s2_rune_list)     //打散的utf8字符转字符串
	//fmt.Println(s2)               //输出 天空好像下雪，我好像住你隔壁
	//
	////组合切片成一个字符串
	//fmt.Println(strings.Join([]string{"a", "1"}, ","))
	//
	////去掉首位空格放入切片
	//fmt.Println(strings.Fields("w,1,2,3 "))

	a := tokenAmountToEth("32322", 6)
	fmt.Println(a)
}

func tokenAmountToEth(amount string, decimal Uint64) string {
	dLen := int(decimal.Uint64())
	aLen := len(amount)
	if aLen < dLen {
		return fmt.Sprintf("0.%s%s", strings.Repeat("0", dLen-aLen), amount)
	} else if aLen == dLen {
		return fmt.Sprintf("0.%s", amount)
	} else {
		return fmt.Sprintf("%s.%s", amount[:aLen-dLen], amount[aLen-dLen:])
	}
}

type Uint64 uint64

func (c Uint64) Uint64() uint64 {
	return uint64(c)
}


type MS struct {
	Name string
	Cgs  int
}

func Test_Run6(t *testing.T)  {
	hashRun := make(map[string]*MS)
	mss := []MS{
		MS{Name: "m", Cgs: 100},
		MS{Name: "m", Cgs: 1},
	}
	for _, ms := range mss {
		if h, ok := hashRun[ms.Name]; ok {
			h.Cgs += ms.Cgs
		} else {
			hashRun[ms.Name] = &ms
		}
	}
	for _, ms := range hashRun {
		fmt.Println(ms.Cgs)
	}
}

func Test_Run7(t *testing.T)  {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, n := range nums {
		//i = 6
		sum += n
	}
	fmt.Println(sum)
}
