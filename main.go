package main

import (
	"github.com/dolnikov/Algorithms/sorting"
)

type IAlgorithms interface {
	Run()
}

func main() {
	algorithms := sorting.NewSorting()
	run(algorithms)
}

func run(a IAlgorithms) {
	a.Run()
}
