package main

import (
	"fmt"
	"strconv"
)

type contador struct {
	nombre string
	tope   int
}

func (c *contador) run() {
	for i := 1; i <= c.tope; i++ {
		fmt.Println(c.nombre, " ", strconv.Itoa(i))
	}
}

func main() {
	persona1 := contador{
		nombre: "Aaron",
		tope:   5,
	}
	persona2 := contador{
		nombre: "Juan",
		tope:   7,
	}
	persona1.run()
	persona2.run()
}
