package validParentheses

type Stack []string

func (s *Stack) Push(e string) {
	*s = append(*s, e)
}

func (s *Stack) Pop() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}

	popVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return popVal, true
}

func (s *Stack) TopVal() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}

	return (*s)[len(*s)-1], true
}
