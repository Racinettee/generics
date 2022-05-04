package generics

type Stack[X any] List[X]

func (stack Stack[X]) Top() X {
	return stack[len(stack)-1]
}
