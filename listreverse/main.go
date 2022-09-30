package main

import "fmt"

func main() {
	link := &List{}

	ListPushBack(link, 1)
	ListPushBack(link, 2)
	ListPushBack(link, 3)
	ListPushBack(link, 4)

	ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	var tmp *NodeL
	var tmp2 *NodeL
	if l == nil {
		return
	}
	for l.Head != nil {
		tmp, l.Head.Next = l.Head.Next, tmp2
		tmp2, l.Head = l.Head, tmp
	}
	l.Head = tmp2
}_

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
