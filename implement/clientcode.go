package main

import (
	"fmt"
)

type my struct {
	species [6]byte
	name [16]byte
	surname [32]byte
	age int
}

func (seed *uint64) generate() (out my) {
	for j := 0; j < 3; j++ {
		for i := range(out.species) {
			o *= 16777619
			out.species[i] ^= o
		}
		for i := range(out.name) {
			o *= 16777619
			out.name[i] ^= o
		}
		for i := range(out.surname) {
			o *= 16777619
			out.surname[i] ^= o
		}
		o *= 16777619
		out.age[i] ^= o
	}
}


func () {
	programmers []my


}
