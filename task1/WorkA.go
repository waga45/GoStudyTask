package task1

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
  - Task1 136.找到[]int中只出现了一次的元素
    大于100元素该方法处理快
*/
func ArrayItemJustOne(arr *[]int) ([]int, error) {
	if arr == nil {
		return nil, errors.New("err: the parameter of arr is empty or nil!")
	}
	var tempMap = make(map[int]int)
	for _, v := range *arr {
		//出现次数，是否存在
		count, flag := tempMap[v]
		if !flag {
			tempMap[v] = 1
		} else {
			tempMap[v] = count + 1
		}
	}
	result := []int{}
	for k, v := range tempMap {
		if v <= 1 {
			result = append(result, k)
		}
	}
	return result, nil
}

/*
*
Task1 136.找到[]int中只出现了一次的元素
小于100元素下该方法快
*/
func ArrayItemJustOneOther(arr *[]int) ([]int, error) {
	if arr == nil {
		return nil, errors.New("err: the parameter of arr is empty or nil!")
	}
	var resultList = []int{}
	for _, v := range *arr {
		var showTime = 0
		var currentItem = v
		for j := 0; j < len(*arr); j++ {
			if currentItem == (*arr)[j] {
				showTime++
				if showTime > 1 {
					break
				}
			}
		}
		if showTime == 1 {
			resultList = append(resultList, v)
		}
	}
	return resultList, nil
}

/*
*
是否回文整数

	（回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数 如 121  1331）
*/
func IsHwNumber(num int) bool {
	strNum := strconv.Itoa(num)
	//转Unicode
	strNumArr := []rune(strNum)
	if len(strNumArr) <= 1 {
		return false
	}
	if len(strNumArr) == 2 {
		return string(strNumArr[0]) == string(strNumArr[1])
	}

	var leftArr = []rune{}
	var rightArr = []rune{}
	if len(strNumArr)%2 == 0 {
		//偶数
		size := len(strNumArr) / 2
		leftArr = append(leftArr, strNumArr[:size]...)
		rightArrReverse := reverseArray(strNumArr[size:])
		rightArr = append(rightArr, rightArrReverse...)
	} else {
		//基数
		size := (len(strNumArr) - 1) / 2
		leftArr = append(leftArr, strNumArr[:size]...)
		rightArrReverse := reverseArray(strNumArr[size+1:])
		rightArr = append(rightArr, rightArrReverse...)
	}
	return strings.Compare(string(leftArr), string(rightArr)) == 0
}

// 反序
func reverseArray(ss []rune) []rune {
	swapArr := make([]rune, len(ss))
	for i := 0; i < len(ss); i++ {
		swapArr[i] = (ss)[len(ss)-1-i]
	}
	return swapArr
}

/*
*

	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/
func EnableParticularlyStr(input string) bool {
	if (len(input) <= 1) || (len(input)%2 != 0) {
		return false
	}
	if len(input) == 2 {
		return charCompare(string(input[0]), string(input[1]))
	}
	inputArr := []rune(input)
	var itemCount = findNoRepeatLen(&inputArr)
	var trueCount = 0
	loopFind(&inputArr, &trueCount)
	return trueCount == itemCount
}
func loopFind(arr *[]rune, trueCount *int) {
	for i := 0; i < len(*arr); i++ {
		fmt.Println("loop start :", *arr)
		var cr = string((*arr)[i])
		var cn string
		if (i + 1) < len(*arr) {
			cn = string((*arr)[i+1])
		} else {
			cn = string((*arr)[i])
		}
		b := charCompare(cr, cn)
		if b == true {
			*arr = append((*arr)[:i], (*arr)[i+2:]...)
			*trueCount += 1
			loopFind(arr, trueCount)
		}
	}
}
func findNoRepeatLen(arr *[]rune) int {
	var tempMap = map[string]int{}
	for _, v := range *arr {
		_, flag := tempMap[string(v)]
		if !flag {
			tempMap[string(v)] = 1
		}
	}
	delete(tempMap, ")")
	delete(tempMap, "}")
	delete(tempMap, "]")
	return len(tempMap)
}
func charCompare(c1 string, c2 string) bool {
	switch c1 {
	case "(":
		if c2 == ")" {
			return true
		}
	case "{":
		if c2 == "}" {
			return true
		}
	case "[":
		if c2 == "]" {
			return true
		}
	}
	return false
}

/*
计算最长公共长缀
*/
func FrequentlyMaxChar(content *[]string) string {
	if len(*content) <= 0 {
		return ""
	}
	var maxItem = ""
	for _, v := range *content {
		if len(v) > len(maxItem) {
			maxItem = v
		}
	}
	maxArray := []string{}
	arrSize := len(*content)
	for i := 0; i < len(maxItem); i++ {
		//current char
		var chr = maxItem[i]
		var count = 0
		for _, c := range *content {
			if (len(c) >= i+1) && (chr == c[i]) {
				count++
			}
		}
		if count == arrSize {
			maxArray = append(maxArray, string(chr))
		}
	}
	return strings.Join(maxArray, "")
}

/*
*
数+1
*/
func DigitsAddition(nums []int) []int {
	if len(nums) <= 0 {
		return nums
	}
	ss := make([]string, len(nums))
	for i, v := range nums {
		ss[i] = strconv.Itoa(v)
	}
	num, _ := strconv.Atoi(strings.Join(ss, ""))
	num += 1

	var resultArr = []int{}
	for _, v := range strconv.Itoa(num) {
		fmt.Println(v)
		n, _ := strconv.Atoi(string(v))
		resultArr = append(resultArr, n)
	}
	return resultArr
}

/*
*
移除重复
*/
func RemoveRepeatItem(arr []int) (int, []int) {
	if len(arr) <= 1 {
		return len(arr), arr
	}
	slow := 0
	// 4,1,2,2,3,4,2
	for i := 1; i < len(arr); i++ {
		if (arr)[i] != (arr)[slow] {
			slow++
			arr[slow] = arr[i]
		}
	}
	fmt.Println(arr[:slow+1])
	return slow + 1, arr[:slow+1]
}

/*
* 56.合并区间
 */
func RemoveRepeatDemension(list [][]int) [][]int {
	if len(list) <= 1 {
		return list
	}
	fmt.Println("输入：", list)
	var resultArr = [][]int{}
	for rowIndex := 0; rowIndex < len(list); rowIndex++ {
		one := list[rowIndex]
		onceMax := one[1]
		for j := rowIndex + 1; j < len(list); j++ {
			next := list[j]
			if one[1]-next[0] >= 0 {
				onceMax = next[1]
				rowIndex = j
			}
		}
		resultArr = append(resultArr, []int{one[0], onceMax})
	}
	fmt.Println("合并区间结果：", resultArr)
	return resultArr
}

/*
*两数之和
 */
func CaculateToTarget(arr []int, target int) []int {
	if len(arr) <= 0 {
		return nil
	}
	for i := 0; i < len(arr); i++ {
		var c = arr[i]
		find := false
		findRightIndex := 0
		for j := i + 1; j < len(arr); j++ {
			if (c + arr[j]) == target {
				find = true
				findRightIndex = j
				break
			}
		}
		if find {
			return []int{i, findRightIndex}
		}
	}
	return nil
}
