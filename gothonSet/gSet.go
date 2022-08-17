package gothonSet

type T any

type Set map[T]bool

func NewSet[T comparable](a []T) Set {
	s := Set{}
	for _, v := range a {
		s[v] = true
	}
	return s
}
