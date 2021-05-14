package main

import (
	"fmt"
)

func main(){
	fmt.Printf("\n")
	testIf()
	testFor()

}
func testIf(){
	a1 := 11
	if a1 == 11 {
		fmt.Println("a1 = 11")
	}else{
		fmt.Println("a1 != 11")
	}	

	//多个判断条件
	age := 19
	if age > 35 {
		fmt.Println("中年")
	}else if age > 18 {
		fmt.Println("青年")
	}else {
		fmt.Println("小学生")
	}

	//在if作用域声明变量
	if treeNum := 9; treeNum > 8 {
		fmt.Println("有很多树")
	}else if treeNum < 8 {
		fmt.Println("有一些树")
	}else {
		fmt.Println("没有树")
	}
}
func testFor(){
	musicName := "关于我转生到异世界成为史莱姆这件事"
	fmt.Println(musicName)
	musicNameRune := []rune(musicName)
	for i := 0; i < 17; i++{
		fmt.Println(string(musicNameRune[i]))
		//切片之后转换为字符串
	}
}
