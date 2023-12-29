package data_structures

import "fmt"

type node struct {
	nextpointer *node
	value       int
}

type listnode struct {
	headpointer, tailpointer *node
	length                   int
}

func newNode(value int) *node {
	return &node{nextpointer: nil, value: value}
}

func NewListNode() *listnode {
	return &listnode{headpointer: nil, tailpointer: nil, length: 0}
}

func (list *listnode) Push(value int) {
	head := newNode(value)
	if list.headpointer == nil {
		list.headpointer = head
		list.tailpointer = head
		list.length++
	} else {
		list.tailpointer.nextpointer = head
		list.tailpointer = head
		list.length++
	}
}

func (listnode *listnode) swap(n1, n2 *node) {
	temp := n1.value
	n1.value = n2.value
	n2.value = temp
}

func (list *listnode) LenA() int {
	return list.length
}

func (list *listnode) PrintList() {
	var obj_pointer *node = list.headpointer
	for obj_pointer != nil {
		fmt.Printf("%d, ", obj_pointer.value)
		obj_pointer = obj_pointer.nextpointer

	}
}

func (list *listnode) BubbleSortListNodeAsc() {
	for i := 0; i < list.length; i++ {
		is_swaped := false
		head_pointer := list.headpointer
		for j := 0; j < list.length-1 && head_pointer.nextpointer != nil; j++ {
			if head_pointer.value > head_pointer.nextpointer.value {
				list.swap(head_pointer.nextpointer, head_pointer)
				is_swaped = true
			}
			head_pointer = head_pointer.nextpointer
		}
		if is_swaped == false {
			break
		}
	}
}

func (list *listnode) LinkedListLinearSerch(target int) int {
	pointer_node := list.headpointer
	for i := 0; i < list.length; i++ {
		if pointer_node.value == target {
			return i
		}
		pointer_node = pointer_node.nextpointer
	}
	return -1
}

func (l *listnode) InsertAtNthPositionInListNode(position, value int) {
	n_node := node{nextpointer: nil, value: value}
	toInsertAfter := l.headpointer
	for i := 1; i <= l.length; i++ {
		if i == position {
			n_node.nextpointer = toInsertAfter.nextpointer
			toInsertAfter.nextpointer = &n_node
			l.length++
		}
		toInsertAfter = toInsertAfter.nextpointer
	}
}

type DoubleLinkNode struct {
	prev, next         *DoubleLinkNode
	doublelinklistnode *DoubleLinkListNode
	Value              any
}

type DoubleLinkListNode struct {
	pointer *DoubleLinkNode
	length  int
}

func (doublelinklistnode *DoubleLinkListNode) Init() *DoubleLinkListNode {
	doublelinklistnode.pointer.prev = doublelinklistnode.pointer
	doublelinklistnode.pointer.next = doublelinklistnode.pointer
	doublelinklistnode.length = 0
	return doublelinklistnode
}

func (double *DoubleLinkListNode) newInit() {
	if double.pointer.next == nil {
		double.Init()
	}
}

func NewLinkListNode() *DoubleLinkListNode {
	return new(DoubleLinkListNode).Init()
}

func (doublelinklistnode *DoubleLinkListNode) Insert(datanode, at *DoubleLinkNode) *DoubleLinkListNode {
	datanode.prev = at
	datanode.next = at.next
	datanode.prev.next = datanode
	datanode.next.prev = datanode
	datanode.doublelinklistnode = doublelinklistnode
	return doublelinklistnode
}

func (doublelinklistnode *DoubleLinkListNode) InserValue(value any, datanode *DoubleLinkNode) *DoubleLinkListNode {
	return doublelinklistnode.Insert(&DoubleLinkNode{Value: value}, datanode)
}
