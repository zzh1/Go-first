package main

import "fmt"

func main(){

	/*
	//切片的创建
	var s1 []int = make([]int,5,8)
	var s2 []int = make([]int,8) //满容切片

	fmt.Println(s1)
	fmt.Println(s2)
	
	var s []int = []int{1,2,3,4,5} //满容的
	fmt.Println(s);

	var len int = len(s1)
	fmt.Println(len)
	var cap int = cap(s1)
	fmt.Println(cap) 
	*/

	/*	
	// 空切片
	var s1 []int
	var s2 []int = []int{}
	var s3 []int = make([]int,0)

	fmt.Println(s1,s2,s3)
	fmt.Println(len(s1),len(s2),len(s3))
	fmt.Println(cap(s1),cap(s2),cap(s3))
	*/

	/*
	// 切片的赋值
	var s1 = make([]int,5,8)
	// 切片的访问和数组差不多
	for i := 0; i < len(s1); i++{
		s1[i] = i+1
	}		
	
	var s2 = s1

	fmt.Println(s1,len(s1),cap(s1));
	fmt.Println(s2,len(s2),cap(s2));

	//尝试修改切片的内容
	s2[0] = 255
	fmt.Println(s1)
	fmt.Println(s2)
	//拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。
	*/

	//切片的遍历
	// var s = []int{1,2,3,4,5}
	// for index:=range s{
	// 	fmt.Println(index,s[index])
	// }
	
	// for index,value:=range s{
	// 	fmt.Println(index,value)
	// }
	
	/*
	//切片遍历
	var s1 = []int{1,2,3,4,5}
	fmt.Println(s1,len(s1),cap(s1))
	
	// 对满容的切片进行追加会分离底层数组
	var s2 = append(s1,6)

	fmt.Println(s1,len(s1),cap(s1))
	fmt.Println(s2,len(s2),cap(s2))

	// 对非满容的切片进行追加会共享底层数组
	var s3 = append(s2,7)

	fmt.Println(s2,len(s2),cap(s2))
	fmt.Println(s3,len(s3),cap(s3))
	
	// var s1 = []int{1,2,3,4,5}
	// append(s1,6)
	// fmt.Println(s1)

	var s1 = []int{1,2,3,4,5}
	_ = append(s1,6)
	fmt.Println(s1)
	*/

	// 切割切割
	// var s1 = []int{1,2,3,4,5,6,7}
	
	// // start_index 和 end_index, 不包含end_index
	// // [start_index,end_index)
	// var s2 = s1[2:5]
	// fmt.Println(s1,len(s1),cap(s1))
	// fmt.Println(s2,len(s2),cap(s2))

	// var s1 = []int{1,2,3,4,5,6,7}
	// var s2 = s1[:5]
	// var s3 = s1[3:]
	// var s4 = s1[:]

	// fmt.Println(s1,len(s1),cap(s1))
	// fmt.Println(s2,len(s2),cap(s2))
	// fmt.Println(s3,len(s3),cap(s3))
	// fmt.Println(s4,len(s4),cap(s4))

	// var s = make([]int,5,8)
	// for i:=0;i<len(s);i++{
	// 	s[i] = i+1
	// }
	
	// fmt.Println(s,len(s),cap(s))

	// var s2 = s
	// var s3 = s[:]

	// fmt.Println(s2,len(s2),cap(s2))
	// fmt.Println(s3,len(s3),cap(s3))

	// //修改母切片
	// s[0] = 255
	// fmt.Println(s,len(s),cap(s))
	// fmt.Println(s2,len(s2),cap(s2))
	// fmt.Println(s3,len(s3),cap(s3))

	/*
	//数组变切片
	var a = [10]int{1,2,3,4,5,6,7,8,9,10}
	var b = a[2:6]
	fmt.Println(b)
	a[4]=100
	fmt.Println(b)
	*/

	/*
	//copy函数
	var s = make([]int,5,8)
	for i:=0;i<len(s);i++{
		s[i] = i+1
	}

	fmt.Println(s)
	var d = make([]int,2,6)
	var n = copy(d,s)

	fmt.Println(n,d)
	*/

	s1 := make([]int,6)
	s2 := make([]int,1024)
	s1 = append(s1,1)
	s2 = append(s2,2)
	fmt.Println(len(s1),cap(s1))
	fmt.Println(len(s2),cap(s2))







}

