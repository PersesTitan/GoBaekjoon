package page5

func sum(a []int) int {
	s := 0
	for value := range a {
		s += a[value]
	}
	return s
}
