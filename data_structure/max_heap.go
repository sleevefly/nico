package data_structure

import "fmt"

// MaxHeap 结构表示大顶堆
type MaxHeap struct {
	arr []int
}

// NewMaxHeap 创建一个新的大顶堆
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// Len 返回堆的大小
func (h *MaxHeap) Len() int {
	return len(h.arr)
}

// Parent 返回指定索引的父节点索引
func (h *MaxHeap) Parent(i int) int {
	return (i - 1) / 2
}

// LeftChild 返回指定索引的左子节点索引
func (h *MaxHeap) LeftChild(i int) int {
	return 2*i + 1
}

// RightChild 返回指定索引的右子节点索引
func (h *MaxHeap) RightChild(i int) int {
	return 2*i + 2
}

// Swap 交换堆中两个元素的位置
func (h *MaxHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

// Push 向堆中添加一个元素
func (h *MaxHeap) Push(val int) {
	h.arr = append(h.arr, val)
	h.siftUp(len(h.arr) - 1)
}

// Pop 从堆中取出最大的元素
func (h *MaxHeap) Pop() (int, error) {
	if h.Len() == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	root := h.arr[0]
	lastIndex := len(h.arr) - 1

	h.Swap(0, lastIndex)
	h.arr = h.arr[:lastIndex]
	h.siftDown(0)

	return root, nil
}

// siftUp 从指定索引开始上浮元素，维护堆的性质
func (h *MaxHeap) siftUp(index int) {
	for index > 0 && h.arr[h.Parent(index)] < h.arr[index] {
		h.Swap(index, h.Parent(index))
		index = h.Parent(index)
	}
}

// siftDown 从指定索引开始下沉元素，维护堆的性质
func (h *MaxHeap) siftDown(index int) {
	maxIndex := index
	leftChild := h.LeftChild(index)
	rightChild := h.RightChild(index)

	if leftChild < h.Len() && h.arr[leftChild] > h.arr[maxIndex] {
		maxIndex = leftChild
	}

	if rightChild < h.Len() && h.arr[rightChild] > h.arr[maxIndex] {
		maxIndex = rightChild
	}

	if index != maxIndex {
		h.Swap(index, maxIndex)
		h.siftDown(maxIndex)
	}
}

type MyHeap struct {
	arr []int
}

func NewMyHeap() *MyHeap {
	return &MyHeap{}
}

func (h *MyHeap) Push(val int) {
	h.arr = append(h.arr, val)
}

func (h *MyHeap) upSkip(index int) {
	for index > 0 && h.arr[h.Parent(index)] < h.arr[index] {
		h.Swap(index, h.Parent(index))
		index = h.Parent(index)
	}
}

func (h *MyHeap) Parent(index int) int {
	return (index - 1) / 2
}
func (h *MyHeap) Swap(index, parent int) {
	h.arr[index], h.arr[parent] = h.arr[parent], h.arr[index]
}
