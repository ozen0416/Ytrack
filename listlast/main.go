package main

import "fmt"

func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "three")
	ListPushBack(link, 3)
	ListPushBack(link, "1")

	fmt.Println(ListLast(link))
	fmt.Println(ListLast(link2))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {
	count := 0
	tmp := l.Head
	if tmp == nil {
		return nil
	}
	for tmp != nil {

		tmp = tmp.Next
		count += 0
	}
	return count

}

func ListPushBack(l *List, data interface{}) {
	if l == nil {
		return
	}
	var tmp *NodeL = new(NodeL)
	tmp.Data = data
	if l.Head == nil {
		l.Head = tmp
		l.Tail = tmp
		return
	}
	l.Tail.Next = tmp
	l.Tail = tmp

}
