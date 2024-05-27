package datastructure


type Queue[T any] struct{
	data []T
}

type queued [T any] interface{
	Enqueue(T)
	Dequeue() T
}


func (q *Queue[T])Enqueue(item T){
	q.data = append(q.data, item)
}


func (q *Queue[T]) Dequeue() T{
	if len(q.data) == 0{
		panic("queue is empty")
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item
}