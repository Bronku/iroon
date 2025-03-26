package main

import "github.com/Bronku/iroon/cmd/iroon"

type test struct {
	data []int
}

func (t *test) xd() []int {
	return t.data
}

func (t *test) setData(input []int) {
	t.data = make([]int, len(input))
	copy(t.data, input)
}

func main() {
	iroon.Run()
}
