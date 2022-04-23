package generics

type List[X any] []X

func (l *List[X]) Push(v X) {
	*l = append(*l, v)
}

func (l *List[X]) Pop() X {
	res := (*l)[len(*l)-1]
	*l = (*l)[0 : len(*l)-1]
	return res
}

func (l List[X]) Len() int {
	return len(l)
}
