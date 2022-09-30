package main

import "fmt"

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)

	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 1).Data)
	fmt.Println(ListAt(link.Head, 7))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	count := 0
	for l != nil {
		if count == pos {
			return l
		}
		l = l.Next
		count++
	}
	return l
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
