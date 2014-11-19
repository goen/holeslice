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

type seed uint64

func (o *seed) gen() (out my) {
	for j := 0; j < 3; j++ {
		for i := range(out.species) {
			(*o) *= 16777619
			out.species[i] ^= byte(*o)
		}
		for i := range(out.name) {
			(*o) *= 16777619
			out.name[i] ^= byte(*o)
		}
		for i := range(out.surname) {
			(*o) *= 16777619
			out.surname[i] ^= byte(*o)
		}
		(*o) *= 16777619
		out.age ^= int(*o)
	}
	return
}


func main() {
	var o seed = 2166136261
	var programmers []my

	for i := 0; i < 100000; i++ {
		programmers = append(programmers, (&o).gen())
	}

	fmt.Println(programmers[100])
}
