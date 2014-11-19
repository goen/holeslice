package slice

import "testing"

func Checksum1(h [2][2][]T) (o uint64) {
	o = 2166136261
	for i := 0; i < Len1(h); i++ {
		o = (o * 16777619) ^ uint64(At1(h, i))
	}
	return o
}

func TestDeletes1(t *testing.T) {
	var h [2][2][]T
	var q []T = []T("..A!A,,oXXo..B!B,,")

	h = Append1(h, q...)

	if Checksum1(h) != 5070384084921349961 {
		t.Error("Initial Append")
	}
	h = Delete1(h, 8, 2)
	if Checksum1(h) != 17321018617259080033 {
		t.Error("Split In Half delete")
	}
	h = Delete1(h, 7, 2)
	if Checksum1(h) != 10505201662031302483 {
		t.Error("Cross delete")
	}
	h = Delete1(h, 0, 2)
	if Checksum1(h) != 12187123277419432363 {
		t.Error("First start delete")
	}
	h = Delete1(h, 3, 2)
	if Checksum1(h) != 5187593657894948051 {
		t.Error("First end delete")
	}
	h = Delete1(h, 3, 2)
	if Checksum1(h) != 924793118297961255 {
		t.Error("Second start delete")
	}
	h = Delete1(h, 6, 2)
	if Checksum1(h) != 1150925254849619335 {
		t.Error("Second end delete")
	}
	h = Delete1(h, 4, 1)
	if Checksum1(h) != 13472835742592135716 {
		t.Error("Second internal delete")
	}
	h = Delete1(h, 1, 1)
	if Checksum1(h) != 13875767474068099931 {
		t.Error("First internal delete")
	}
}
