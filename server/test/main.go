package main

import (
	"fmt"
	"strconv"
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

	testTypeValue()
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
