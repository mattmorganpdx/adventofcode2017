package intIter


type IntIter struct {
	value   []int
	length  int
	current int
}

func NewIntIter(input []int) IntIter {
	return IntIter{input, len(input), 0}

}

func (this *IntIter) Next() int {
	if this.current+1 >= this.length {
		this.current = 1
	} else {
		this.current++
	}
	return this.value[this.current-1]
}

func (this *IntIter) Peek() int {
	var next int
	if this.current+1 >= this.length {
		next = 0
	} else {
		next = this.current + 1
	}
	return this.value[next]
}

func (this *IntIter) Current() int {
	return this.value[this.current]
}