package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

/**
- Golang是强类型语言
- 变量赋值过程中，类型必须保持一致
- 变量需要先定义后使用，且必须被用到
- 变量都会有一个默认值
- 变量不能重名
- 会根据变量的值推断变量类型
*/
func main() {
	// 变量的声明
	var hello string = "hello"
	var world = "world"
	fmt.Println(hello, world)

	// 变量的类型推断
	var int2 = 18
	fmt.Println(int2)

	// 中文也可以当作变量名
	中文变量名 := 1.1
	fmt.Println(中文变量名)

	// 一行定义一组变量
	var a, b uint = 1, 2
	fmt.Println(a, b)

	// 多行定义变量
	var (
		c int     = 1
		d float64 = 2
	)
	fmt.Println(c, d)

	// 变量运算的时候左右类型保持一致
	var height = 10.1
	var width float64 = 1.1
	fmt.Println(height * width)

	// 变量的默认值
	var namedefault string
	fmt.Println(namedefault)
	var agedefault int
	fmt.Println(agedefault)

	// 变量类型的推断 reflect
	a1 := "hello"
	fmt.Println(reflect.TypeOf(a1))
	a2 := 9
	fmt.Println(reflect.TypeOf(a2))
	a3 := 5.4
	fmt.Println(reflect.TypeOf(a3))
	a4 := &a3
	fmt.Println(reflect.TypeOf(a4)) // *float64

	// 变量类型转化的代价
	len1, len2 := 1, 2.1
	// 数据类型的强制转化
	fmt.Println(len1 * int(len2))
	// 2^64 - 1 无符号的最大值
	var len3 uint = math.MaxUint64 // 18446744073709551615
	fmt.Println(len3)
	// 无符号的最大值转化成，有符号的，那么精度就会缺失
	var len4 int = int(len3)
	// 原码 补码
	fmt.Println(len4) // -1

	// 采用:声明变量的时候，变量是否是重定义
	intNums := 10
	fmt.Printf("%p\n", &intNums)
	intNums, stringName := 19, "intNums"
	fmt.Printf("%p\n", &intNums)
	fmt.Println(intNums, stringName)
	// 每个花括号是一个作用域
	{
		intNums := 26
		fmt.Printf("&intNums=%p   intNums=%d \n", &intNums, intNums)
	}

	// 字符串转化整形，不一定会成功
	stringChange := "测试"
	changeValue, err := strconv.Atoi(stringChange)
	fmt.Println(stringChange, changeValue, err)
	// Yes 123 123 <nil>
	// No  0 strconv.Atoi: parsing "测试": invalid syntax
}
