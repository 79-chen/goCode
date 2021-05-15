package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	testArray()
	testSlice()

}
func testArray() {
	//数组长度是数据类型的一部分，长度有限制
	var a0 [5]int = [5]int{1, 2, 3, 4, 5}
	a1 := [...]int{1, 2, 3}  //自动判断长度
	a2 := [4]int{0: 1, 3: 4} //根据索引初始化
	fmt.Println(a0)
	fmt.Println(a1)
	fmt.Println(a2)
	a3 := [...]string{"大小", "方位", "指向"}
	//遍历
	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}
	for i, v := range a3 {
		fmt.Println(i, v)
	}

	//多维数组--删除多余代码:声明与初始化合并，一维数组直接用大括号
	var aa [3][2]int = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(aa)
	for _, v := range aa {
		fmt.Println(v)
	}

	a4 := [...]int{1, 2, 3, 5, 7, 8}
	var r int
	for i := 0; i < len(a4); i++ {
		r += a4[i]
	}
	fmt.Println(r)

	for i := 0; i < len(a4); i++ {
		for j := i + 1; j < len(a4); j++ {
			if a4[i]+a4[j] == 9 {
				fmt.Printf("(%d,%d)\n", a4[i], a4[j])
			}
		}
	}
}
func testSlice() {
	//切片是一个引用数据类型，都指向一个底层数组
	var s1 []int
	var s2 []string
	s1 = []int{1, 2, 3}
	s2 = []string{"chen", "shu", "peng"}
	fmt.Println(s1, s2)
	fmt.Printf("长度：%d,容量：%d\n", len(s1), cap(s1))

	//由数组切割得到切片
	a1 := [...]int{1, 2, 3, 4, 5}
	s3 := a1[0:3] //左闭右开切
	s4 := a1[3:]  //切完
	s5 := a1[:]   //全切
	s6 := s3[0:2] //一切再切
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)
	a1[0] = 999
	fmt.Println(s3)
	fmt.Println(s6)
	s3[1] = 999
	fmt.Println(s3)
	fmt.Println(s6)
	/*
			底层数组的值（堆内存）改变了，切片的值（堆内存）也改变
		    切片由索引改变了的值，实际上改变的是底层数组的值，切片仅是引用（指针）（栈内存）
	*/
	fmt.Printf("长度%d,容量%d\n", len(s3), cap(s3))
	//切片的容量是从它的第一个元素到其底层数组元素末尾的个数

	ss := make([]int, 5, 10)
	fmt.Println(ss, len(ss), cap(ss))
	//make()函数创建切片,(类型，长度，容量)

	for i := 0; i < len(s3); i++ {
		fmt.Println(i, s3[i])
	}
	//遍历
	for i, v := range s3 {
		fmt.Println(i, v)
	}

	//切片扩容append函数
	ssp := []string{"哈密瓜", "西瓜", "荔枝"}
	fmt.Printf("长度：%d,容量：%d\n", len(ssp), cap(ssp))
	ssp = append(ssp, "苹果")
	fmt.Printf("长度：%d,容量：%d\n", len(ssp), cap(ssp))
	ssr := []string{"菠萝", "火龙果"}
	ssp = append(ssp, ssr...) //...表示拆开一个一个追加
	/*
		调用append函数必须用原来的切片变量接收返回值
		append追加元素时，底层数组放不下时，Go底层会把底层数组换一个
	*/
	fmt.Printf("长度：%d,容量：%d\n", len(ssp), cap(ssp))
	/*切片扩容规则
		首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
	    否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
	    否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，
		即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，
		即（newcap >= cap）
	    如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）
	*/

	//copy()函数
	sscc := [...]int{1, 2, 3, 4}
	ssc0 := sscc[:]
	ssc1 := ssc0 //赋值
	var ssc2 []int = make([]int, 4)
	copy(ssc2, ssc0) //copy
	fmt.Println(ssc0, ssc1, ssc2)
	ssc0[0] = 999
	fmt.Println(ssc0, ssc1, ssc2)
	//copy的切片，其底层数组改变

	//删除元素
	ssc0 = append(ssc0[0:1], ssc0[2:]...)
	//删除索引为1的元素，容量不变
	//但是修改了底层数组索引为1的元素的值为索引为2的值，索引2为索引3...！！！
	fmt.Println(ssc0)
	fmt.Println(sscc)

	//练习
	var test []int = make([]int, 5, 10)
	//长度5的默认零值
	for i := 0; i < 9; i++ {
		test = append(test, i)
	}
	fmt.Println(test, cap(test))

}
