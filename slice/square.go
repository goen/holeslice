package slice

func Append1(h [2][2][]T, s ...T) [2][2][]T {
	return Push1(h, s)
}

func Push1(h [2][2][]T, s []T) [2][2][]T {
	i := itoi(Len0(h[1]))
	h[i] = Push0(h[i], s)
	return h
}

func Left1(h [2][2][]T, n int) [2][2][]T {
	return Delete1(h, 0, n)
}

func Right1(h [2][2][]T, n int) [2][2][]T {
	return Delete1(h, n, Len1(h)-n)
}

func Len1(h [2][2][]T) int {
	return Len0(h[0]) + Len0(h[1])
}

func Delete1(h [2][2][]T, n, l int) [2][2][]T {
	l0 := Len0(h[0])
	if n >= l0 {
		h[1] = Delete0(h[1], n-l0, l)
		return h
	} else if n + l == l0 {
		h[0] = Right0(h[0], n)
	}
	if n + l >= l0 {
		h[1] = Left0(h[1], n+l-l0)

	} else {
		h[0] = Delete0(h[0], n, l)
	}
	return h
}

func At1(h [2][2][]T, off int) T {
	d := [2]int{off, off - Len0(h[0])}
	return At0(h[posi(d[1])], d[posi(d[1])])
}
