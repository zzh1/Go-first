package main

import "fmt"

func main(){

	// var a [9]int
	// fmt.Println(a);

	//三种定义形式
	// var a = [9]int{1,2,3,4,5,6,7,8,9}
	// var b [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
	// c:=[8]int{1,2,3,4,5,6,7,8}

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	/*
	var squares [9]int
	for i:=0;i<len(squares);i++{
		squares[i] = (i+1) * (i+1)
	}
	fmt.Println(squares)
	*/

	/*
	//数组下表越界情况
	// var a = [5]int{1,2,3,4,5}
	// a[101] = 255
	// fmt.Println(a)

	var a = [5]int{1,2,3,4,5}
	var b = 101
	a[b] = 255
	fmt.Println(a)
	*/


	/*
	//数组赋值
	var a = [9]int{1,2,3,4,5,6,7,8,9}
	var b [9]int
	b=a
	a[0] = 12345
	fmt.Println(a)
	fmt.Println(b)
	
	//不同长度的数组之间赋值是禁止的
	var a = [9]int{1,2,3,4,5,6,7,8,9}
	var b [10]int
	b=a
	fmt.Println(b)
	*/

	var a = [5]int{1,2,3,4,5}
	for index:=range a{
		fmt.Println(index,a[index])
	}
	/*
	for index,value:=range a {
		fmt.Println(index,value)
	}*/

}