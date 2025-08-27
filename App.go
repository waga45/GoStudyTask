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

	var ln = "([{}])"
	ln = "[()]"
	b := task1.EnableParticularlyStr(ln)
	fmt.Println("是否有效字符串:", b)

	var ss = []string{"flat", "flower", "fliter"}
	res := task1.FrequentlyMaxChar(&ss)
	fmt.Println("最长公共前缀：", res)

	var dd = []int{1, 2, 3, 6}
	dd_res := task1.DigitsAddition(dd)
	fmt.Println("值+1 结果：", dd_res)

	task1.RemoveRepeatItem([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})

	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	task1.RemoveRepeatDemension(intervals)

	nums := []int{2, 7, 11, 15}
	indexResult := task1.CaculateToTarget(nums, 17)
	fmt.Println(indexResult)
}
