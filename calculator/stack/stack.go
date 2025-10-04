package stack

type Stack []int

func (s Stack) Push(elem int) Stack {
	return append(s, elem)
}

func (s Stack) Pop() (Stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s Stack) Pick() int {
	l := len(s)
	return s[l-1]
}

func (s Stack) Empty() bool {
	l := len(s)
	return l == 0
}
