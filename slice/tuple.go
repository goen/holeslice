package slice

type T byte

func Append0(h [2][]T, s ...T) [2][]T {
	return Push0(h, s)
}

func Push0(h [2][]T, s []T) [2][]T {
	i := itoi(len(h[1]))
	h[i] = append(h[i], s...)
	return h
}

func Left0(h [2][]T, n int) [2][]T {
	return Delete0(h, 0, n)
}

func Right0(h [2][]T, n int) [2][]T {
	return Delete0(h, n, Len0(h)-n)
}

func Len0(h [2][]T) int {
	return len(h[0]) + len(h[1])
}

func Delete0(h [2][]T, n, l int) [2][]T {
	l1 := len(h[1])
	l0 := len(h[0])

	if l1 == 0 && n > 0 {
		h[1] = h[0][n+l:]
		h[0] = h[0][:n]
		return h
	}

	if n == 0 {
		if l >= l0 {
			h[0] = h[1][l-l0:]
			h[1] = []T{}
		} else {
			h[0] = h[0][l:]
		}
	} else if n < l0 {
		if n+l == l0 {
			h[0] = h[0][:n]
			return h
		} else if n+l > l0 {
			h[0] = h[0][:n]
			h[1] = h[1][n+l-l0:]
			return h
		}
		h[0] = append(h[0][:n], h[0][n+l:]...)
	} else if n == l0 {
		h[1] = h[1][l:]
		return h
	} else if n+l == l0+l1 {
		h[1] = h[1][:n-l0]
		return h
	} else {
		h[1] = append(h[1][:n-l0], h[1][n-l0+l:]...)
	}

	return h
}

func At0(h [2][]T, off int) T {
	d := [2]int{off, off - len(h[0])}
	p := posi(d[1])
	return h[p][d[p]]
}
