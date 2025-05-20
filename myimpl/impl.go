package myimpl

import (
	"fmt"
	"math"
	"runtime"
)

type Vertex struct {
	X int
	Y int
}

// Uses the Newton-Raphson method to implement a square root function
func MySqrt(x float64) float64 {
	const epsilon = 1e-10

	z := float64(x)
	count := 0

	for i := 0; i < 100; i, count = i+1, count+1 {
		prevValue := z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Prev: %f, z: %f\n", prevValue, z)
		fmt.Println(prevValue == z)
		if math.Abs(prevValue-z) < epsilon {
			break
		}
	}

	return z

}

func CheckOs() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOs")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

}

func UseStruct() {
	var ptr *Vertex
	vertex := Vertex{X: 56, Y: 5}

	ptr = &vertex

	ptr.Y = 1e9

	fmt.Println(vertex, *ptr)
}

func UseArrays() {
	intSlice := make([]int, 10)

	// This has a capacity of 7
	slice1 := intSlice[3:]

	fmt.Println(cap(slice1))
}

func AppendBytes(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)

	if n > cap(slice) { // If necessary, reallocate
		// allocate what's necessary for future growth
		fmt.Println("Using this")
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func (vertexPointer *Vertex) UseVertexMethod() int {
	vertex := *vertexPointer

	vertex.X = vertex.X << 2
	vertex.Y = vertex.Y << 2

	*vertexPointer = vertex

	return vertexPointer.X*2 + vertexPointer.Y*2
}
