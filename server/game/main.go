package main

import "moba/game/logic"

//go中定义变量

//var a,b,c int

//var x=10.21

//ab:=10

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	//for true {
	//	fmt.Printf("game loop")
	//}

	logic.Loop()
	logic.Init()
	logic.Close()
}
