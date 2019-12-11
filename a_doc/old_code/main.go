package old_code

import (
	"fmt"
)

func main() {
	const constVar float32 = 3.1415926
	var count int = 10

	fmt.Println("开始")
	println(count)
	fmt.Println("打印常量的值:  ", constVar)
	if constVar > 2 {
		fmt.Println("常量比2大")
	}
}
