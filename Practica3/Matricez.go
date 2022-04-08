package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var matriz1 [][]int
var matriz2 [][]int
var matriz3 [][]int

func crearMatriz(a int) [][]int {
	var matr = make([][]int, a)
	for i := 0; i < a; i++ {
		matr[i] = make([]int, a)
	}

	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			matr[i][j] = rand.Intn(50)
		}
	}

	return matr
}

func AsigMatriz(numero int) {
	matriz1 = crearMatriz(numero)
	print("Matriz 1:\n")
	ImpMatriz(matriz1)

	matriz2 = crearMatriz(numero)
	print("Matriz 2:\n")
	ImpMatriz(matriz2)

	matriz3 = make([][]int, numero)
	for i := 0; i < numero; i++ {
		matriz3[i] = make([]int, numero)
	}
}

func ImpMatriz(Mat [][]int) {

	for i := 0; i < len(Mat); i++ {
		fmt.Printf("%v \n", Mat[i])
	}
	fmt.Print("\n")
}

func VeriMatriz() bool {
	return len(matriz1) == len(matriz2[0])
}

func suma(i int, j int, fin chan bool) {
	<-fin

	for k := 0; k < len(matriz1[0]); k++ {
		matriz3[i][j] += matriz1[i][k] * matriz2[k][j]
	}
	fin <- true
}

func MultMatriz() {
	finalizar := make(chan bool, len(matriz3)*len(matriz3[0]))
	if VeriMatriz() {
		for i := 0; i < len(matriz1); i++ {
			for j := 0; j < len(matriz2[0]); j++ {
				go suma(i, j, finalizar)
			}
		}

		var aux float64 = float64(len(matriz3))
		finalizar <- true
		for i := 1; i < 2*int(math.Round(math.Logb(aux))); i++ {
			finalizar <- true
		}

	} else {
		fmt.Print("No se puede multiplicar las matrices.")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var numero int
	fmt.Print("Ingrese el tamaño de la matriz: ")
	fmt.Scanln(&numero)

	AsigMatriz(numero)

	var inicio time.Time
	inicio = time.Now()
	MultMatriz()

	for true {
		if runtime.NumGoroutine() == 1 {
			fmt.Println("Tiempo Total: ", time.Now().Sub(inicio))
			/*print("Matriz Resultante:\n")
			ImpMatriz(matriz3)*/

			break
		}
	}
}
