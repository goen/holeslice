package slice

import "testing"

func Checksum0(h [2][]T) (o uint64) {
	o = 2166136261
	for i := 0; i < Len0(h); i++ {
		o = (o * 16777619) ^ uint64(At0(h, i))
	}
	return o
}

func TestDeletes0(t *testing.T) {
	var h [2][]T
	var q []T = []T("..A!A,,oXXo..B!B,,")

	h = Append0(h, q...)

	if Checksum0(h) != 5070384084921349961 {
		t.Error("Initial Append")
	}
	h = Delete0(h, 8, 2)
	if Checksum0(h) != 17321018617259080033 {
		t.Error("Split In Half delete")
	}
	h = Delete0(h, 7, 2)
	if Checksum0(h) != 10505201662031302483 {
		t.Error("Cross delete")
	}
	h = Delete0(h, 0, 2)
	if Checksum0(h) != 12187123277419432363 {
		t.Error("First start delete")
	}
	h = Delete0(h, 3, 2)
	if Checksum0(h) != 5187593657894948051 {
		t.Error("First end delete")
	}
	h = Delete0(h, 3, 2)
	if Checksum0(h) != 924793118297961255 {
		t.Error("Second start delete")
	}
	h = Delete0(h, 6, 2)
	if Checksum0(h) != 1150925254849619335 {
		t.Error("Second end delete")
	}
	h = Delete0(h, 4, 1)
	if Checksum0(h) != 13472835742592135716 {
		t.Error("Second internal delete")
	}
	h = Delete0(h, 1, 1)
	if Checksum0(h) != 13875767474068099931 {
		t.Error("First internal delete")
	}
}
