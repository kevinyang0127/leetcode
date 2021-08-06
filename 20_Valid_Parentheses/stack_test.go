package validParentheses

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	myStack := new(Stack)
	myStack.Push("kevin")
	myStack.Push("peter")
	myStack.Push("curry")

	for _, v := range *myStack {
		fmt.Println(v)
	}

	myStack.Pop()

	for _, v := range *myStack {
		fmt.Println(v)
	}

}
