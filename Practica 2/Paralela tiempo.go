package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

type contador struct {
	nombre string
	tope   int
}

func (c *contador) run() {
	var tiempo1 time.Time
	tiempo1 = time.Now()

	for i := 1; i <= c.tope; i++ {
		/*time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)*/
		fmt.Println(c.nombre, " ", strconv.Itoa(i))
	}

	fmt.Println("Tiempo Total del Hilo", c.nombre, ": ", time.Now().Sub(tiempo1))
}

func main() {
	/*rand.Seed(time.Now().UnixNano())*/

	var lista [3]contador

	lista[0] = contador{
		nombre: "Aaron",
		tope:   5,
	}
	lista[1] = contador{
		nombre: "Juan",
		tope:   7,
	}
	lista[2] = contador{
		nombre: "Miguel",
		tope:   6,
	}

	var inicio time.Time
	inicio = time.Now()

	for i := 0; i < len(lista); i++ {
		go lista[i].run()
	}

	for true {
		if runtime.NumGoroutine() == 1 {
			fmt.Println("Tiempo Total: ", time.Now().Sub(inicio))
			break
		}
	}

	var s string
	fmt.Scan(&s)
}
