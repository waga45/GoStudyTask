package main

import (
	"StudyTask/task1"
	"StudyTask/utils"
	"fmt"
)

func main() {
	testArr := []int{1, 1, 2, 3, 4, 3, 7, 6, 7, 6}
	result_arr := utils.StaticsProcessInfo(task1.ArrayItemJustOne)
	fmt.Println("只出现1次的数字有：", result_arr(&testArr))
	result_arr1 := utils.StaticsProcessInfo(task1.ArrayItemJustOneOther)
	fmt.Println("只出现1次的数字有：", result_arr1(&testArr))

	var testNum = 12221
	ishw := task1.IsHwNumber(testNum)
	fmt.Println("数字", testNum, " 是否为回文:", ishw)
}
