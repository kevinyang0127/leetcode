package validParentheses

func isValid(s string) bool {
	stack := &Stack{}

	for _, rune := range s {
		str := string(rune)
		top, ok := stack.TopVal()

		if !ok {
			stack.Push(str)
			continue
		}

		if match(top, str) {
			stack.Pop()
		} else {
			stack.Push(str)
		}
	}

	return len(*stack) == 0
}

func match(leftCh string, rightCh string) bool {
	return leftCh == "(" && rightCh == ")" || leftCh == "{" && rightCh == "}" || leftCh == "[" && rightCh == "]"
}
