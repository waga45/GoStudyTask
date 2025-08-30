package task2

import "fmt"

/*
*
指针参数
*/
func addValue(v *int) {
	*v += 10
}
func addValue1(v int) {
	v += 10
}

/*
*编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10
 */
func PtrNumCount() {
	inital := 10
	fmt.Println("before value:", inital)
	addValue(&inital)
	fmt.Println("after value:", inital)
	//this is copy an available to function
	addValue1(inital)
	fmt.Println("after value:", inital)
}

// 指针切片元素*2
func arrayItemMutil(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] *= 2
	}
}

/*
*
实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
*/
func PtrArrayMut() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("before arr:", arr)
	arrayItemMutil(&arr)
	fmt.Println("after arr:", arr)
}
