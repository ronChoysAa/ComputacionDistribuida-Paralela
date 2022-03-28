package main

import (
	"fmt"
	"time"
)

type contador struct {
	nombre string
	tope   int
}

func (c *contador) run() {
	for true {
		time.Sleep(time.Duration(1) * time.Millisecond)
		fmt.Println("a")
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

	var inicio time.Time
	inicio = time.Now()

	persona1.run()
	persona2.run()

	fmt.Println("Tiempo Total: ", time.Now().Sub(inicio))
}
