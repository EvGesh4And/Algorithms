type KthLargest struct {
	k    int
	heap []int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		k:    k,
		heap: make([]int, 0, k),
	}
	for _, num := range nums {
		kth.Add(num)
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	if len(this.heap) < this.k {
		this.heap = append(this.heap, val)
		shiftUp(this.heap, len(this.heap)-1)
	} else if val > this.heap[0] {
		this.heap[0] = val
		heapify(this.heap, 0)
	}
	return this.heap[0]
}

func heapify(arr []int, i int) {
	n := len(arr)
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] < arr[smallest] {
		smallest = left
	}
	if right < n && arr[right] < arr[smallest] {
		smallest = right
	}

	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		heapify(arr, smallest)
	}
}

func shiftUp(arr []int, i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if arr[parent] <= arr[i] {
			break
		}
		arr[parent], arr[i] = arr[i], arr[parent]
		i = parent
	}
}
