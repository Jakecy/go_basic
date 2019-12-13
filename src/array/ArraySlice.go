package main

import "fmt"

func main() {
	//declare1()
	//以short hand创建数组并赋值
	//declareArrayUsingShortHand()
	//使用三个点...创建消息
	declareArrayUsing3Dots()
}

func declareArrayUsing3Dots() {
	a := [...]int{12, 18, 19}
	fmt.Println(a)
}

func declareArrayUsingShortHand() {
	//
	a := [3]int{12, 18}
	fmt.Println(a)
}

func declare1() {
	//声明一个包含3个int元素的数组
	var a [3]int
	//为数组元素赋值
	a[0] = 12
	a[1] = 78
	a[2] = 50
	fmt.Println(a)
}
