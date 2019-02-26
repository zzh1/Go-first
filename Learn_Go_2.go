package main

import "fmt"

//全局变量，如果全局变量的首字母大写，那么它就是公开的全局变量。如果全局变量小写，那么它就是内部的全局变量。
//var globali int = 24

//const globali int = 24

func main(){

	//有符号整数，可以表示正负
	var a int8 = 1 //1字节
	var b int16 = 2 //2字节
	var c int32 = 3 //3字节
	var d int64 = 4 //4字节
	fmt.Println(a,b,c,d)

	//无符号整数，只能表示非负数  
	var ua uint8 = 1 
	var ub uint16 = 2
	var uc uint32 = 3
	var ud uint64 = 4 
	fmt.Println(ua,ub,uc,ud)

	//int 类型，在32位机器上占4个字节，在63位及其上占8个字节
	var e int = 5
	var ue uint = 5
	fmt.Println(e,ue)

	//bool 类型
	var f bool = true
	fmt.Println(f)

	//字节类型
	var j byte = 'a'
	fmt.Println(j)

	//字符串类型
	var g string = "abcdefg"
	fmt.Println(g)

	//浮点数
	var h float32 = 3.14
	var i float64 = 3.141592653
	fmt.Println(h,i)



	// 注释：        先CTRL+K，然后CTRL+C
	// 取消注释： 先CTRL+K，然后CTRL+U
	// var value int = 42
	// var p1 *int = &value
	//var p2 **int = &p1
	//var p3 ***int = &p2
	//fmt.Println(p1,p2,p3)
	//fmt.Println(*p1,**p2,***p3)


	//指针
	//var value int = 42
	//var pointer *int = &value
	//fmt.Println(pointer,*pointer)

	//变量与常量
	//const locali int = 42
	//fmt.Println(globali,locali)

	// 全局变量与局部变量
	//var locali int = 42
	//fmt.Println(globali,locali);


	//第2课 变量
	//变量的三种初始化方式
	//var s int = 42
	//var s = 42
	//s := 42
	//fmt.Println(s);
}


