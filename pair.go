package generics

type Pair[X, Y any] struct {
	Left  X
	Right Y
}

func NewPair[X, Y any](left X, right Y) Pair[X, Y] {
	return Pair[X, Y]{left, right}
}
