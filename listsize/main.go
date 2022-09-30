package main

import "fmt"

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	count := 0
	tmp := l.Head
	for tmp != nil {
		tmp = tmp.Next
		count += 1
	}
	return count
}

func ListPushFront(l *List, data interface{}) {
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
	tmp.Next = l.Head
	l.Head = tmp

}
