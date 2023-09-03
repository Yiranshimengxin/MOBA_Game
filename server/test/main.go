package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// 模仿C#枚举类型
const (
	p1 = iota
	p2
	p3
	p4
	p5
	p6
	p7
)

func main() {
	//v := 10
	//ptr := &v
	//fmt.Printf("%p\n", &v)
	//fmt.Printf("%p\n", ptr)
	//fmt.Printf("v的数据类型为：%T\n", v)
	//fmt.Printf("ptr的数据类型为：%T\n", ptr)
	//
	//value := *ptr
	//fmt.Printf("value的值为：%d\n", value)
	//fmt.Printf("value的数据类型为：%T\n", value)
	//
	//value = 20
	//*ptr = 30
	//fmt.Printf("value=%d\n", value)
	//fmt.Printf("v=%d\n", v)
	//
	//str := "hello world"
	//fmt.Printf(str)
	//a := &str
	//*a = "hello game"
	//fmt.Printf("\n" + *a + "\n")
	//
	//age := new(int)
	//*age = 22
	//fmt.Println(*age)
	//fmt.Printf("age的类型为：%T\n", age)
	//
	//distance := new(float64)
	//*distance = 100.1
	//fmt.Println(*distance)
	//fmt.Printf("distanc的类型为：%T\n", distance)
	//
	//const pi = 3.14
	//
	//fmt.Println(p1)
	//fmt.Println(p2)
	//fmt.Println(p3)
	//fmt.Println(p4)
	//fmt.Println(p5)
	//fmt.Println(p6)
	//fmt.Println(p7)

	//testTypeValue()

	//testIfElse()

	//testFor()

	//testSwitchCase()
	/////////////////////////////////////////////////////////////////////
	//
	//ch = make(chan int)
	//wg.Add(2)
	//
	//go running(1)
	//go running1()
	//
	////var input string
	////fmt.Scanln(&input)
	//
	//wg.Wait()
	//fmt.Printf("主程序退出,count=%d\n", atomicInt.Load())
	/////////////////////////////////////////////////////////////////////

	/////////////////////////////////////////////////////////////////////
	ch1 = make(chan string)
	wg.Add(2)
	go runingWrite()
	go runingRead()
	wg.Wait()
	fmt.Printf("主程序退出,此时读取到的字符串为%n", ch1)
	/////////////////////////////////////////////////////////////////////
}

func testTypeValue() {
	//a := 5.0
	a := float32(5.0)

	fmt.Printf("a的类型为%T\n", a)

	b1 := int(a)
	b2 := int32(a)
	b3 := int64(a)

	fmt.Printf("b1=%T,b2=%T,b3=%T", b1, b2, b3)

	c1 := uint(a)
	c2 := uint32(a)
	c3 := uint64(a)

	fmt.Printf("c1=%T,c2=%T,c3=%T", c1, c2, c3)

	//str := "hello"
	//d1 := int(str)
	//
	//fmt.Printf("d1=%T", d1)

	//var ch byte='a'
	//d:=uint8(ch)

	str := "11111"
	d1, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println(d1)
	}

	d2 := 123456
	str2 := strconv.Itoa(d2)
	fmt.Println(str2)

	d3 := int64(123456)
	str3 := strconv.FormatInt(d3, 16)
	fmt.Println(str3)

	d4 := float32(123.456)
	str4 := strconv.FormatFloat(float64(d4), 'f', 6, 64)
	fmt.Println(str4)

	var array [10]int
	array[0] = 1
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
	for k, v := range array {
		fmt.Printf("k=%d,v=%d\n", k, v)
	}

	//q:=[5]int{1,2,3,4,5}
	q := [...]int{1, 2, 3, 4, 5, 6, 7}
	for k, v := range q {
		fmt.Printf("k=%d,v=%d\n", k, v)
	}

	var strSlice []string
	//给切片申请内存
	strSlice = make([]string, 0)
	fmt.Printf("strSlice的类型为%T\n", strSlice)

	// map
	var map1 map[string]int
	map1 = make(map[string]int)
	map1["key"] = 10
	map1["key1"] = 20
	map1["key2"] = 30

	value := map1["key1"]
	fmt.Printf("%d\n", value)

	value1, ok := map1["key6"]
	if ok {
		fmt.Printf("%d\n", value1)
	} else {
		fmt.Println("no value")
	}

	for k, v := range map1 {
		fmt.Printf("k=%v,v=%d\n", k, v)
	}

	// 结构体
	type Person struct {
		name string
		age  int
	}

	type Color struct {
		R, G, B byte
	}
}

func 多返回值() (int, int, int, string) {
	return 1, 2, 3, "world"
}

func testIfElse() {
	ten := 11
	if ten > 10 {
		fmt.Println("ten>10")
	} else if ten < 10 {
		fmt.Println("ten<10")
	} else {
		fmt.Println("ten=10")
	}

	if err := Connect(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("connect success")
	}
}

func Connect() error {
	return nil
}

func testFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for {
		i++
		if i > 10 {
			break
		}
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}

func testSwitchCase() {
	cond := 10
	switch cond {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}
}

func testGoto() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			goto exit
		}
	}

exit:
	fmt.Println("exit")
}

func testFunc(a int, b int) (int, int) {
	return a, b
}

//func testFunc2() {
//	fn := testReturnFunc(1, 2, "name")
//
//}

func testReturnFunc(a int, b int, c string) func(a int, b int, c string) (int, int, string) {
	return func(a, b int, c string) (int, int, string) {
		fmt.Println("testReturnFunc")
		return a, b, c
	}
}

//var lock sync.Mutex
//var count int = 0

var atomicInt atomic.Int64

var wg sync.WaitGroup

// 定义一个int类型的通道
var ch chan int

func running(index int) {

	for i := 0; i < 100; i++ {
		//atomicInt.Add(1)
		ch <- i
	}

	//fmt.Printf("index=%d,count=%d \n", index, atomicInt.Load())
	wg.Done()
}

func running1() {
	for data := range ch {
		fmt.Printf("running1 %d\n", data)
		if data == 99 {
			break
		}
	}
	wg.Done()
}

// //////////////////////////////////////////////////////////////////////////
var ch1 chan string

func runingWrite() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			ch1 <- "over"
			break
		} else {
			ch1 <- "hello"
		}
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func runingRead() {
	for s := range ch1 {
		fmt.Printf("runningWrite %s\n", s)
		if s == "over" {
			break
		}
	}
	wg.Done()
}
