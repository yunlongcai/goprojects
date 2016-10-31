package main

import (
	"fmt"
)

// Main main struct
type Main struct {
	id   int64
	name string
}

func main() {
	fmt.Print("hello world\n")
	m := &Main{id: 1, name: "cyl"}
	fmt.Printf("%#v\n", m)
}
