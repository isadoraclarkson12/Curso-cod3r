package main

import (
	f "fmt"
	"math"
	"reflect"
)

func main() {
	i1 := math.MaxInt64
	var b byte = 255
	f.Println("O tipo é", reflect.TypeOf(b))
	f.Println("Maximo valor", i1)
}
