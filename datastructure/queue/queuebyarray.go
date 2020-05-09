package queue

//数组实现循环队列

type Queue struct {
	tab []int
	head int
	tail int
}

func GetQueueInstance(cap int) *Queue {
	if cap < 1 {
		cap = 10
	}
	
	return &Queue{
		tab:  make([]int, cap+1),
		head: 0,
		tail: 0,
	}
}

func (q *Queue) Push(n int) bool {
	if q.full() {
		return false
	}

	q.tab[q.tail] = n
	q.tail++
	return true
}

func (q *Queue) Pop() int {
	if q.empty() {
		return -1
	}

	res := q.tab[q.head]
	q.head++
	return res
}

func (q *Queue) empty() bool {
	return q.head == q.tail
}

func (q *Queue) full() bool {
	return (q.tail + 1) % len(q.tab) == q.head
}