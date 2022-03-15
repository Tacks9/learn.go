package main

import "fmt"

/**
## 03.strfmt 字符串的格式化

fmt相关文档： http://doc.golang.ltd/

字符串格式化
- %s	直接输出字符串或者[]byte
- %q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
- %x	每个字节用两字符十六进制数表示（使用a-f）
- %X	每个字节用两字符十六进制数表示（使用A-F）

*/
func main() {
	fmt.Println("中文 | ABC | 123 | 😎")

	// 字符串格式化
	var name string = "爱吃麻薯的盖先生"
	fmt.Printf("我的名字是 %s\n", name)
	fmt.Printf("我的名字是 %q\n", name)
	fmt.Printf("我的名字是 %x\n", name)
	fmt.Printf("我的名字是 %X\n", name)
}
