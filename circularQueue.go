package main

import "fmt"

type CircularQueue struct {
	data     []int
	front    int
	rear     int
	size     int
	capacity int
}

// NewCircularQueue Create a new circular queue
func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		data:     make([]int, capacity),
		front:    -1,
		rear:     -1,
		capacity: capacity,
	}
}

// IsFull Check if the queue is full
func (q *CircularQueue) IsFull() bool {
	return (q.rear+1)%q.capacity == q.front
}

// IsEmpty Check if the queue is empty
func (q *CircularQueue) IsEmpty() bool {
	return q.front == -1
}

// Enqueue Add element to the queue
func (q *CircularQueue) Enqueue(value int) error {
	if q.IsFull() {
		return fmt.Errorf("queue is full")
	}

	if q.IsEmpty() {
		q.front = 0
	}

	q.rear = (q.rear + 1) % q.capacity
	q.data[q.rear] = value
	q.size++
	return nil
}

// Dequeue Remove element from the queue
func (q *CircularQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}

	value := q.data[q.front]

	if q.front == q.rear {
		// Queue becomes empty after this operation
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front + 1) % q.capacity
	}

	q.size--
	return value, nil
}

// Peek the front element
func (q *CircularQueue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.data[q.front], nil
}

// Display all elements
func (q *CircularQueue) Display() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	fmt.Print("Queue elements: ")
	i := q.front
	for {
		fmt.Printf("%d ", q.data[i])
		if i == q.rear {
			break
		}
		i = (i + 1) % q.capacity
	}
	fmt.Println()
}

// Example usage
func main() {
	q := NewCircularQueue(5)

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)

	q.Display()

	q.Dequeue()
	q.Dequeue()

	q.Display()

	q.Enqueue(60)
	q.Enqueue(70)

	q.Display()

	val, _ := q.Peek()
	fmt.Println("Front element:", val)
}
