package main

import "fmt"

/**
## 03.operator 常见运算符
- 数学运算符 + - * / %
- 数学运算符 += -= *= /= %=
- 数学运算符 i++ i--
- 比较运算符 &&且 ||或 !取反
- 比较运算符 == != > < >= <=
- 位运算符 & | ^ <<1 >>1

运算符中的一些注意事项
- 运算符计算的过程中，小心乘法加法的溢出，例如+ *
- 运算符计算的过程中，小心除法/的舍去，只保留整数部分，例如 5/2 = 2
- 运算符计算的过程中，小心除数不能为0，否则报错 panic: runtime error: integer divide by zero
- 运算符计算的过程中，小心取余%只能用于整数类型的值，如果用浮点数则会报错
- 异或^小技巧，用来寻找两两一样的，其中有一个值不一样的情况
*/
func main() {

	// 基础运算符
	var a, b int = 10, 3
	fmt.Println("a+b = ", a+b)
	fmt.Println("a-b = ", a-b)
	fmt.Println("a*b = ", a*b)
	fmt.Println("a/b = ", a/b)
	fmt.Println("a%b = ", a%b)

	fmt.Println(true && false)

	fmt.Println("a > b  ", a > b)

	// 异或^小技巧
	// 用来寻找两两一样的，其中有一个值不一样的情况
	arr := []int{1, 4, 3, 2, 1, 2, 9, 3, 9, 8, 9, 8, 4}
	result := -1
	for _, item := range arr {
		if result < 0 {
			result = item
		} else {
			result = result ^ item
		}
	}
	fmt.Println(arr, result)
}
