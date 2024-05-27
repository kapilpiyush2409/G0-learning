package datastructure

import "fmt"

type Stack struct {
	data     []interface{}
	DataType string
}

type Stacked interface {
	Push(data interface{})
	Pop()
	Top() (interface{}, error)
	IsEmpty() bool
	Print()
}

func NewStack(dataType string) Stacked {
	return &Stack{DataType: dataType}
}

func (S *Stack) Push(data interface{}) {
	if checkDataType(S.DataType, data) {
		S.data = append(S.data, data)
	} else {
		fmt.Println("wrong Datatype")
	}
}

func (S *Stack) Pop()  {
	if !S.IsEmpty(){
		S.data = S.data[:len(S.data)-1]
	}
}

func (S Stack) Top() (interface{}, error) {
	if len(S.data) == 0 {
		return 0, fmt.Errorf("Stack is empty")
	} else {
		return S.data[len(S.data)-1], nil
	}
}

func (S Stack) IsEmpty() bool {
	return len(S.data) == 0
}

func (S Stack) Print() {
	fmt.Println(S.data)
}
