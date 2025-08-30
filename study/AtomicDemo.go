package study

import (
	"unsafe"
)

// 数据节点
type Node struct {
	value interface{}
	next  unsafe.Pointer
}

// 栈
type LockFreeStack struct {
	top unsafe.Pointer
}

//func (l *LockFreeStack) Push(_value interface{}) {
//	newTop := &Node{value: _value}
//
//	for {
//
//		currentTop := atomic.LoadUintptr(l.top)
//		newTop.next = currentTop
//	}
//}
