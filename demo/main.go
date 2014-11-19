package main

import (
	. "example.com/repo.git/holeslice/slice"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	var h [2][2][]T

	fmt.Println("Appending input to slice. X Y deletes\n" +
		"y characters on X-th place. Type collect to gather to one slice.")

	for {
		n := -1
		o := -1
		s := ""
		fmt.Scanf("%s", &s)
		if s == "collect" {
			h[0][0] = append(h[0][0], h[0][0]...)
			h[0][0] = append(h[0][0], h[0][1]...)
			h[0][0] = append(h[0][0], h[1][0]...)
			h[0][0] = append(h[0][0], h[1][1]...)
			h[0][1] = nil
			h[1][0] = nil
			h[1][1] = nil
			continue
		}
		fmt.Sscanf(s, "%d", &n)
		if n >= 0 && n < Len1(h) {
			fmt.Scanf("%s", &s)
			fmt.Sscanf(s, "%d", &o)
			if o >= 0 && n+o < Len1(h) {
				o = 1
			}
			h = Delete1(h, n, o)

		} else {

			h = Append1(h, []T(s)...)
		}

		spew.Dump(h)
	}
}
