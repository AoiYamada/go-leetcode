// thread unsaved stack copied from:
// https://stackoverflow.com/questions/28541609/looking-for-reasonable-stack-implementation-in-golang

package common

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return res
}
