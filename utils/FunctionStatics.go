package utils

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// 统计函数执行情况
func StaticsProcessInfo(fn interface{}) func(...interface{}) []interface{} {
	return func(args ...interface{}) []interface{} {
		// 打印函数名和输入参数
		fnType := reflect.TypeOf(fn)
		fmt.Printf("===== 执行函数: %s =====\n", fnType.Name())
		fmt.Printf("输入参数: ")
		for i, arg := range args {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%v(%T)", arg, arg)
		}
		fmt.Println()
		// 记录开始时间
		start := time.Now()
		//记录内存状态
		var m1, m2 runtime.MemStats
		runtime.ReadMemStats(&m1)
		// 调用原函数
		fnValue := reflect.ValueOf(fn)
		in := make([]reflect.Value, len(args))
		for i, arg := range args {
			in[i] = reflect.ValueOf(arg)
		}
		out := fnValue.Call(in)
		runtime.ReadMemStats(&m2)
		// 计算执行时间
		duration := time.Since(start)
		fmt.Printf("执行耗时: %v\n", duration)
		fmt.Printf("内存分配: %d bytes, 分配次数: %d次\n",
			m2.TotalAlloc-m1.TotalAlloc,
			m2.Mallocs-m1.Mallocs)
		// 打印返回结果
		fmt.Printf("返回结果: ")
		result := make([]interface{}, len(out))
		for i, val := range out {
			result[i] = val.Interface()
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%v(%T)", result[i], result[i])
		}
		fmt.Println("\n========================\n")

		return result
	}
}
