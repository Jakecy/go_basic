package main

import "fmt"

func main() {
	//declare1()
	//以short hand创建数组并赋值
	//declareArrayUsingShortHand()
	//使用三个点...创建消息
	//declareArrayUsing3Dots()
	//获取数组长度
	//getArrayLegth()
	//使用for循环遍历数组
	//literateArrayUsingForLoop()
	//使用range遍历数组
	literateArrayUsingRange()

}

func literateArrayUsingRange() {
	sum := 0
	var arr = [...]int{23, 25, 78, 36, 88, 520}
	for i, v := range arr {
		fmt.Println("每次遍历的：i 是,v 是： ", i, v)
		sum += v
	}
}

func literateArrayUsingForLoop() {
	var arr = [...]int{23, 25, 78, 36, 88, 520}
	//遍历
	for i := 0; i < len(arr); i++ {
		fmt.Println(" 当前的数组元素是: ", arr[i])
	}
}

func getArrayLegth() {
	//创建一个float64数组
	var arr = [...]float64{23.25, 78.36, 88.520}
	fmt.Println("数组arr的长度是: ", len(arr))
}

func arraySizeIsType() {
	/*	a :=[...]int{12,18,19}

		var b [5]int
		b=a*/
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
