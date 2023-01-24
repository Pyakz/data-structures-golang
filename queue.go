package main

import "fmt"



type Queue struct {
	data []string
}


func (q *Queue) Enqueue(data string) {
	q.data = append(q.data, data)
}

func (q *Queue) Dequeue() {
	if(len(q.data) == 0) {
		return
	}
	q.data = q.data[1:]
}

func main() {

	queue := Queue{}


	queue.Enqueue("First")
	queue.Enqueue("Second")
	queue.Enqueue("Third")
	queue.Dequeue()
	queue.Dequeue()


	fmt.Println(queue.data)

}