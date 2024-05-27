package datastructure

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head     *Node
	DataType string
}

type List interface {
	Insert(data interface{})
	Display()
	Length() int
	IsEmpty() bool
	Contains(data interface{}) bool
	Remove(data interface{})
	RemoveAll(data interface{})
	Get(data interface{}) interface{}
	Set(index interface{}, data interface{})
	Clear()
	ToArray() []interface{}
}

func NewLinkedList(dataType string) List {
	list := &LinkedList{DataType: dataType}
	return list
}

func (list LinkedList) IsEmpty() bool {
	if list.head == nil {
		return true
	}
	return false
}

func (list *LinkedList) Insert(data interface{}) {

	if checkDataType(list.DataType, data) {
		nextNode := Node{data: data}
		if list.IsEmpty() {
			list.head = &nextNode
		} else {
			current := list.head
			for current.next != nil {
				current = current.next
			}
			current.next = &nextNode
		}
	} else {
		fmt.Println("wrong Datatype")
	}

}

func (list *LinkedList) Remove(data interface{}) {
	if list.head.data == data {
		list.head = list.head.next
		return
	} else {
		prev := list.head
		for prev.next != nil {
			if prev.next.data == data {
				prev.next = prev.next.next
				return
			}
			prev = prev.next
		}
	}
}

func (list *LinkedList) RemoveAll(data interface{}) {
	if list.head == nil {
		return
	}

	// Remove nodes from the beginning if they contain the data
	for list.head != nil && list.head.data == data {
		list.head = list.head.next
	}

	// Traverse the list to remove nodes containing the data
	prev := list.head
	for prev != nil && prev.next != nil {
		if prev.next.data == data {
			prev.next = prev.next.next
		} else {
			prev = prev.next
		}
	}
}

func (list *LinkedList) Set(index interface{}, data interface{}) {
	if checkDataType(list.DataType, data) {
		count := -1
		if list.IsEmpty() {
			return
		} else {
			current := list.head
			for current != nil {
				count++
				if index == count {
					list.head.data = data
					return
				}
			}
		}
	} else {
		fmt.Println("wrong Datatype")
	}

}

func (list LinkedList) Get(data interface{}) interface{} {
	count := -1
	if list.IsEmpty() {
		return count
	} else {
		current := list.head
		for current != nil {
			count++
			if current.data == data {
				return count
			}
			current = current.next
		}
	}
	return count
}

func (list LinkedList) Length() int {
	count := 0
	if list.IsEmpty() {
		return count
	} else {
		current := list.head
		for current != nil {
			count++
			current = current.next
		}
	}
	return count
}

func (list *LinkedList) Clear() {
	list.head = nil
}

func (list LinkedList) Contains(data interface{}) bool {
	if list.IsEmpty() {
		return false
	} else {
		current := list.head
		for current != nil {
			if current.data == data {
				return true
			}
			current = current.next
		}
	}
	return false
}

func (list LinkedList) ToArray() []interface{} {
	slice := make([]interface{}, 0, list.Length())
	if list.IsEmpty() {
		return slice
	} else {
		current := list.head
		for current != nil {
			slice = append(slice, current.data)
			current = current.next
		}
	}
	return slice
}

func (list LinkedList) Display() {
	if list.IsEmpty() {
		return
	} else {
		current := list.head
		for current != nil {
			fmt.Printf("%d ", current.data)
			current = current.next
		}
		fmt.Println()
	}
}
