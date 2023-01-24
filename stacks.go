package main

import "fmt"

type Stack []string

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

func (s *Stack) Pop() (string, bool) {

	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.

		return element, true
	}
}

func main() {
	var stack Stack // create a stack variable of type Stack

	stack.Push("this")
	stack.Push("is")
	stack.Push("sparta!!")
	stack.Pop()
	fmt.Println(stack)

}
