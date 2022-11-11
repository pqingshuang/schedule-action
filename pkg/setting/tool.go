package setting

func Map[V any](slice []V, f func(V, int) V) []V {
	ls := make([]V, 0)
	for ind, ele := range slice {
		ls = append(ls, f(ele, ind))

	}
	return ls
}
