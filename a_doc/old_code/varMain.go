package old_code

import (
	"fmt"
)

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
	var height = 32
	fmt.Println(" 身高是:  ", height)
}

func varTypeInfer() {
	var weight = 120
	fmt.Println("体重是: ", weight)
}

//多变量声明

func varMutipleVarDeclaration() {
	//语法如下：
	//var name1, name2 type = initialvalue1,initialvalue2
	var name, gender = "chilaoban", "nan"
	fmt.Println("姓名: " + name + ",性别: " + gender)
}

func varDeclarationShortHand() {
	//快捷方式
	a, b := 3, 2
	fmt.Println("快捷方式声明变量")
	fmt.Println(a)
	fmt.Println(b)
}

func calSum(price int, count int) int {
	var sum = price * count
	return sum
}

func main() {
	sum := calSum(20, 2)
	println(sum)
}
