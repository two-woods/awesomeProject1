package main

import "fmt"

var maxCapacity = 10

type sElement int

//SeqStack

func main() {
	sequence := []sElement{1, 2, 3, 4}
	seqStack := make([]sElement, maxCapacity, maxCapacity)
	var top int
	initStack(&top)
	fmt.Println("This stack is empty? ", isEmptyStack(top))
	fmt.Println("This stack is full? ", isFullStack(top))

	for _, item := range sequence {
		pushStack(seqStack, &top, item)
	}
	fmt.Println("Now,we will push on some element", seqStack[:top+1])

	popStack(seqStack, &top)
	fmt.Println("Now,we will pop an element", seqStack[:top+1])
}

// Initial Stack
func initStack(top *int) {
	*top = -1
}

func isEmptyStack(top int) bool {
	return top == -1
}

func isFullStack(top int) bool {
	return top == maxCapacity-1
}

func getTopElement(top int, seqStack []int) (item int) {
	if isEmptyStack(top) {
		return -1
	} else {
		item = seqStack[top]
		return item
	}
}

func pushStack(seqStack []sElement, top *int, item sElement) bool {
	if isFullStack(*top) {
		return false
	}

	*top = *top + 1
	seqStack[*top] = item

	return true
}

func popStack(seqStack []sElement, top *int) bool {
	if isEmptyStack(*top) {
		return false
	} else {
		item := seqStack[*top]
		*top = *top - 1
		fmt.Println("pop Stack", item)
		return true
	}
}
