package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2))
}

type Item struct {
	num   int
	count int
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count < pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	numsCount := make(map[int]int)
	for _, num := range nums {
		numsCount[num]++
	}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for num, count := range numsCount {
		if pq.Len() < k {
			heap.Push(&pq, &Item{num: num, count: count})
			continue
		}
		if count > pq[0].count {
			heap.Pop(&pq)
			heap.Push(&pq, &Item{num: num, count: count})
		}
	}

	result := make([]int, 0, k)
	for pq.Len() != 0 {
		result = append(result, heap.Pop(&pq).(*Item).num)
	}
	return result
}
