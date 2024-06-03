package main

import(
	"fmt"
)


type Node[T any] struct {
	elem any 
	next *Node[T]
	prev *Node[T]
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func New [T any]() List[T] {
	return List[T]{nil, nil}
}

func IsEmpty[T any](self List[T]) bool {
	return self.head == nil
}

func Len[T any](self List[T]) int {
	length := 0
	for current := self.head; current != nil; current = current.next {
		length++
	}
	return length
}

func FrontElement[T any](self List[T]) any {
	if IsEmpty(self) {
		return nil
	}
	return self.head.elem
}

func Next[T any](self List[T]) *Node[T] {
	if IsEmpty(self) {
		return nil
	}
	return self.head.next
}

func ToString[T any](l List[T]) string {
	if IsEmpty(l) {
		return "[]"
	}
	result := ""
	for current := l.head; current != nil; current = current.next {
		result += fmt.Sprintf("(%v)\n", current.elem)
	}
	return result
}

func PushFront[T any](self *List[T], elem any) {
	aux := &Node[T]{elem: elem, next: self.head, prev: nil}
	if IsEmpty(*self) {
		self.tail = aux
	} else {
		self.head.prev = aux
	}
	self.head = aux
}

func PushBack[T any](self *List[T], elem any) {
	aux := &Node[T]{elem: elem, next: nil, prev: self.tail}
	if IsEmpty(*self) {
		self.head = aux
	} else {
		self.tail.next = aux
	}
	self.tail = aux
}

func Remove[T any](self *List[T], actual *Node[T]) {
	if actual == nil || IsEmpty(*self) {
		return
	}
	if actual == self.head {
		self.head = actual.next
		if self.head != nil {
			self.head.prev = nil
		} else {
			self.tail = nil
		}
	} else if actual == self.tail {
		self.tail = actual.prev
		if self.tail != nil {
			self.tail.next = nil
		} else {
			self.head = nil
		}
	} else {
		actual.prev.next = actual.next
		actual.next.prev = actual.prev
	}
}

func main()  {
	num:= New[int]()

	PushBack(&num,23)
	PushBack(&num,45)
	PushBack(&num,213)
	PushBack(&num,826)
	fmt.Println(ToString(num))

}