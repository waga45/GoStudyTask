package task1

import (
	"errors"
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
