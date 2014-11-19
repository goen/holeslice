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
	l1 := Len0(h[1])
	l0 := Len0(h[0])
	if l1 == 0 && n > 0 {
		h[1] = Left0(h[0], n+l)
		h[0] = Right0(h[0], n)
		return h
	}

	if n == 0 {
		if l >= l0 {
			h[0] = Left0(h[1], l-l0)
			h[1] = [2][]T{}
		} else {
			h[0] = Left0(h[0], l)
		}
	} else if n < l0 {
		if n+l == l0 {
			h[0] = Right0(h[0], n)
			return h
		} else if n+l > l0 {
			h[0] = Right0(h[0], n)
			h[1] = Left0(h[1], n+l-l0)
			return h
		}
		h[0] = Delete0(h[0], n, l)
	} else if n == l0 {
		h[1] = Left0(h[1], l)
		return h
	} else if n+l == l0+l1 {
		h[1] = Right0(h[1], n-l0)
		return h
	} else {
		h[1] = Delete0(h[1], n-l0, l)
	}

	return h
}

func At1(h [2][2][]T, off int) T {
	d := [2]int{off, off - Len0(h[0])}
	p := posi(d[1])
	return At0(h[p], d[p])
}
