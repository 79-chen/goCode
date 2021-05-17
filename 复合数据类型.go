package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	testArray()
	testSlice()
	testPointer()
	testMap()
	function()
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
	var test = make([]int, 5, 10)
	//长度5的默认零值
	for i := 0; i < 9; i++ {
		test = append(test, i)
	}
	fmt.Println(test, cap(test))

}
func testPointer() {
	//取地址符：& 获得变量的指针变量;
	//根据地址取值：* 对指针变量取值，获得指针变量指向的原变量的值
	a := 10
	n := &a
	fmt.Println(n)
	fmt.Printf("类型：%T\n", n) //int类型的指针
	p := *n                  //指针变量名为n
	fmt.Println(p)

	//new函数
	var a1 *int
	fmt.Println(a1)   //空指针，没有指向内存空间
	var a2 = new(int) //开辟一块存放对应类型变量的地址，指针指向的变量默认零值
	fmt.Println(*a2)
	fmt.Println(a2)
	*a2 = 100
	fmt.Println(a2)
	fmt.Println(*a2)

	/*
		make 和 new 都是用来申请内存的
		new 用于基本数据类型，返回值为对应类型的指针，内存对应的值为类型零值
		make 用于map slice channel三种数据类型的初始化，返回值为对应的类型（引用类型本身）
	*/
}
func testMap() {
	var m1 = make(map[string]int, 10)
	//引用数据类型，初始化（开辟内存空间）后才能使用（避免空指针异常）
	//估算好map的容量，避免在程序运行期间再动态扩容
	m1["chen"] = 10
	m1["shu"] = 11
	m1["peng"] = 12
	fmt.Println(m1["shu"])
	fmt.Println(m1) //无顺序
	fmt.Println(m1["chenshupeng"])
	//如无对应key，返回对应类型的零值
	value, ok := m1["chenshupeng"]
	if !ok {
		fmt.Println("查无此key！")
	} else {
		fmt.Println(value)
	}

	//遍历，key对应下标
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//删除
	delete(m1, "chen")
	fmt.Println(m1)

	//元素类型为map的切片
	var s1 = make([]map[int]string, 1, 10)
	s1[0] = make(map[int]string, 2)
	//map也需要初始化
	s1[0][999] = "chenshupeng"
	fmt.Println(s1[0][999]) //map的key对应的值
	fmt.Println(s1[0])      //map数据类型
	fmt.Println(s1)         //切片数据类型

	//值为切片类型的map
	var s2 = make(map[string][]int, 10)
	s2["chenshupeng"] = []int{1, 2, 3}
	fmt.Println(s2)
}
func function() func(int, int) string {
	//函数也是一种类型
	var f1 func(int, int)string
	
	// func f1(a1,a2,a3 int, b1,b2,b3 int) (ret int){
	// 	ret = 10//ret已经声明
	// 	return a+b
	// }
	// func f2(a int, b int) int{
	// 	ret := 10//ret未声明
	// 	return a+b
	// }
	// func f3(x int, y ...int) {
	// 	fmt.Println()//y的类型是切片[]int
	// }

	defer fmt.Println("end") //返回值赋值-->defer先进(状态完整)后出-->return
	fmt.Println("next")

	//匿名函数：函数内部无法声明带名字的函数
	f2 := func() string {
		return "ohhhhhhhh!!"
	}
	fmt.Println(f2())
	//只用一次的匿名函数
	func(a, b int) {
		fmt.Println(a + b)
	}(1, 2)

	//闭包内在要求：
	//1.返回值可以是函数
	//2.返回的函数存在外层函数的变量
	//先外层入参返回内层函数，再内层函数入参输出

	return f1
}
