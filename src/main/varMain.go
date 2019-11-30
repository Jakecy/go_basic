package main

import "fmt"

func variableFuncV1() {
	var age int
	fmt.Println(" my age is :", age)
}

func variableFuncV2() {
	//为变量赋值
	var gender string
	gender = "男"
	fmt.Println("性别是: ", gender)
}

func variableFuncV3() {
	//声明变量时候，指明初始值
	var height int = 32
	fmt.Println(" 身高是:  ", height)
}

func main() {
	variableFuncV1()
	variableFuncV2()
	variableFuncV3()
}
