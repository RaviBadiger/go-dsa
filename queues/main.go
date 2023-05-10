package main

import "fmt"

type Queue struct{
	items []int
}

func (q *Queue) Enqueue(data int) {
	q.items = append(q.items, data)
}

func (q *Queue) Dequeue()int{
	toDequeue := q.items[0]
	q.items =  q.items[1:]
	return toDequeue
}

func main(){
	myQueue := Queue{}
	myQueue.Enqueue(12)
	myQueue.Enqueue(13)
	myQueue.Enqueue(14)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Print(myQueue)
}