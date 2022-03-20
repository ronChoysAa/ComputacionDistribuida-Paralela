package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type contador struct {
	nombre string
	tope   int
}

func (c *contador) run() {
	for i := 1; i <= c.tope; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println(c.nombre, " ", strconv.Itoa(i))
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())

	persona1 := contador{
		nombre: "Aaron",
		tope:   5,
	}
	persona2 := contador{
		nombre: "Juan",
		tope:   7,
	}

	go persona1.run()
	go persona2.run()

	var s string
	fmt.Scan(&s)
}
