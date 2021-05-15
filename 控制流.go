package main

import (
	"fmt"
)

func main(){
	fmt.Printf("\n")
	testIf()
	testFor()
	testSwitch()

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
	s := "hello树鹏"
	for v,d := range s {
		fmt.Printf("%d %c\n",v ,d)
		//数组，切片，字符串返回索引和值，索引按字节数来
		//一个汉字一般占三字节
	}
	for y := 1; y <= 9; y++ {
        // 遍历, 决定这一行有多少列
        for x := 1; x <= y; x++ {
            fmt.Printf(" %d x %d = %d ", x, y, x*y)
        }
        // 手动生成回车
        fmt.Println()
    }
	
	for i := 9; i > 0; i-- {
		for j := 1; j <= i; j++ {
			if i == 9 || i == 7 {
				break
				//跳出单层循环
				
			}
			if i == 5 || i == 3 {
				continue
				//跳过单层循环的单次循环，继续下一次循环
			}
			if i == 1 {
				goto xend
				//go 跳到对应标签
			}
			fmt.Printf(" %d x %d = %d ",j,i,i*j)
		}
		fmt.Println()
	}
	xend:
	fmt.Println("End")
}
func testSwitch(){
	n := '水'
	switch n {
	case '水':
		fmt.Println("水之呼吸-一字型-水面斩")
	case '炎':
		fmt.Println("炎之呼吸-一字型-阳炎")
	case '雷':
		fmt.Println("雷之呼吸-一字型-霹雳一闪")
	case '虫':
		fmt.Println("虫之呼吸-一字型-蝶之舞")
	default:
		fmt.Println("血鬼术！")
		//兜底操作
	}
	a := 9
	switch {
	case a >= 9:
		fmt.Println("a >= 9")
	case a < 9:
		fmt.Println("a < 9")
    default:
		fmt.Println("Error")
	}
	switch a := 9; a {
	case 1,2,3,4,5:
		fmt.Println()
	case 6,7,8,9,10:
		fmt.Println()
	default:
		fmt.Println("Error")
	}
}

