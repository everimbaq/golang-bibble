package syntax

import (
	"fmt"
	"reflect"
	"unsafe"
)

// go和c语言指针和数组的区别
// 在c语言中，数组名不等于指针，数组是数组，指针是指针，指针需要占用内存空间来存储地址
// 数组不能被直接复制，当数组名作为函数参数出现时，要么是数组的引用，要么是指向第一个元素的指针
// 数组和指针对于sizeof来说是不同的，指针变量占用的空间通常等于当前CPU的最大位数，数组名取sizeof的话，得到的则是数组的大小。
// 对数组名取地址&是合法的，但有些编译器不推荐这样做，对数组名取地址的结果与直接使用数组名的结果是一致的，这是C语言的一种特殊规定。
// 当你对一个数组做&的时候，他提取的是指向数组的指针，然后仍然可以隐式转换成指向第一个元素的指针，而且它们的值是相等的。
func test() {
	var arr = [3]int{1, 2, 3}
	fmt.Println("array:", arr, &arr, reflect.TypeOf(arr), unsafe.Sizeof(arr))
	paar := arr
	// 怎么讲一个指针变量赋给数组？
	fmt.Println(" =array:", paar, &paar, reflect.TypeOf(paar), unsafe.Sizeof(paar))
	var t = []int{6, 7, 8}
	fmt.Println("slice:", t, &t, reflect.TypeOf(t), unsafe.Sizeof(t))
	var arr2 *[]int
	arr2 = &t
	// 指向slice的指针， type是一个*[]int指针，指向t的真实地址， size是8，因为指针地址的长度为8
	fmt.Println("pointer to slice", arr2, &arr2, reflect.TypeOf(arr2), unsafe.Sizeof(arr2))

}
