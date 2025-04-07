type RecentCounter struct {
	stack []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	tPast := t - 3000
	this.stack = append(this.stack, t)
	for len(this.stack) > 0 && this.stack[0] < tPast {
		this.stack = this.stack[1:]
	}
	return len(this.stack)
}