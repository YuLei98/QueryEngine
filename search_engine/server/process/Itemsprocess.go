package process

import (
	"Search_Engines/search_engines/server/model"
	"fmt"
)

//优先级队列实现
type PriorityQueueInterface interface {
	GetSize() int
	IsEmpty() bool
	Push(node model.Item)
	Pop() *model.Item
	Top() *model.Item
}

//小根堆
type PriorityQueue struct {
	Size int
	Item []model.Item
}

//用于堆过滤
type Itemsprocess struct {
}

func NewQueue(max int) PriorityQueueInterface {
	return &PriorityQueue{Size: 0, Item: make([]model.Item, max)}
}

func (p *PriorityQueue) GetSize() int {
	return p.Size
}

func (p *PriorityQueue) IsEmpty() bool {
	return len(p.Item) == 0
}

func (p *PriorityQueue) Push(node model.Item) {

	i := p.Size
	p.Size++
	for {
		if i <= 0 {
			break
		}

		parent := (i - 1) / 2

		if p.Item[parent].Score >= node.Score {
			break
		}
		p.Item[parent], p.Item[i] = p.Item[i], p.Item[parent]
		i = parent
	}

	p.Item[i] = node
}

func (p *PriorityQueue) Pop() *model.Item {

	if p.Size == 0 {
		return nil
	}
	root := p.Item[0]
	i := 0
	p.Size--

	last := p.Item[p.Size]
	p.Item[p.Size] = model.Item{}

	for {
		left := 2*i + 1
		right := 2*i + 2

		if left >= p.Size {
			break
		}

		if right > p.Size && p.Item[left].Score < p.Item[right].Score {
			left = right
		}

		if p.Item[left].Score < last.Score {
			break
		}

		p.Item[i], p.Item[left] = p.Item[left], p.Item[i]
		i = left
	}

	p.Item[i] = last
	return &root
}

func (p *PriorityQueue) Top() *model.Item {
	return &p.Item[0]
}

func (this *Itemsprocess) Heap_Fast(limit int, items []model.Item) (result []model.Item) {
	fmt.Println("go go heap fast ");
	q := NewQueue(limit + 1)
	for _, v := range items {
		v.Score = -v.Score
		q.Push(v)
		if q.GetSize() > limit {
			q.Pop()
		}
	}
	length := q.GetSize()
	result = make([]model.Item, length)
	for q.GetSize() > 0 {
		v := q.Top()
		v.Score = -v.Score
		length = length - 1
		result[length] =(*v);
		if length==0{break;}
		q.Pop()
	}
	return
}
