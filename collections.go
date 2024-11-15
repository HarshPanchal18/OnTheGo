package main

import (
	"container/list"
	"fmt"
)

func doublyLinkedList() {

	var x list.List
	x.PushBack(1)
	x.PushBack("Anything")

	for ele := x.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele.Value)
	}

}