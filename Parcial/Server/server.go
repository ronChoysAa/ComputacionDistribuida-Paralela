package main

import (
	"fmt"
	"math"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

const (
	SERVER_HOST = "0.0.0.0"
	//SERVER_HOST = "localhost"
	SERVER_PORT = "5000"
	SERVER_TYPE = "tcp"
)

var final float64
var hilos int = 20

func valor(funcion string, eli string) float64 {
	var ind string = (strings.TrimRight(funcion, eli))
	var val float64

	if len(ind) == 0 {
		return 1
	} else {
		val, _ = strconv.ParseFloat(ind, 64)
		return val
	}
}

func segmento(funcion string, a float64, b float64, finalizar chan bool, wg *sync.WaitGroup) {
	<-finalizar

	var altx float64
	if strings.Index(funcion, "x^2") != -1 {
		altx = a * a
		altx = valor(funcion, "x^2") * altx
	} else if strings.Index(funcion, "x") != -1 {
		altx = a
		altx = valor(funcion, "x") * altx
	} else {
		altx, _ = strconv.ParseFloat(funcion, 64)
	}

	var resul float64 = float64((b - a) * altx)
	final += resul
	wg.Done()
	finalizar <- true
}

func area(funcion string, a float64, b float64, incremento float64) float64 {

	final = 0
	var cant = int((b - a) / incremento)

	finalizar := make(chan bool, cant)
	wg := sync.WaitGroup{}

	var ejem = 0

	for i := a; ejem < cant; i = i + incremento {
		go segmento(funcion, i, i+incremento, finalizar, &wg)
		ejem = ejem + 1
	}

	wg.Add(cant)

	if hilos < cant {
		for i := 1; i < hilos; i++ {
			finalizar <- true
		}
	} else {
		for i := 1; i < cant; i++ {
			finalizar <- true
		}
	}

	wg.Wait()

	fmt.Printf("Ecuacion %s: %.2f\n", funcion, final)
	return float64(math.Round(final*100) / 100)

}
func Ejercicio(funcion string, a float64, b float64, resultado float64) string {
	var i = 1.0
	var gap = 0.5
	//Corre primera funcion
	var res float64 = area(funcion, a, b, i)

	//DEFINICION DE VALORES
	valorres, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", res), 64)

	//Definicion de rango
	rangomenor, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", resultado-gap), 64)
	rangomayor, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", resultado+gap), 64)

	//fmt.Printf("Rango: %.2f y %.2f", rangomayor, rangomenor)
	//fmt.Println()
	for j := 1; j < 5; j++ {

		//Se llego al resultado
		if valorres > rangomenor && valorres < rangomayor {
			return fmt.Sprintf("Resultado Final:%.2f \n", final)

		} else {
			//No se llego al resultado
			i = i / 10
			res = area(funcion, a, b, i)
			valorres, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", res), 64)

		}
	}
	return fmt.Sprintln("Resultado no obtenido")
}

func main() {

	//Ejercicio("x", 3, 10, 45.5)
	//Ejercicio("x^2", 3, 10, 324.33)
	fmt.Println("Servidor iniciado...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Cliente conectado")

		go cliente(connection)
	}
}

func cliente(connection net.Conn) {
	var funcion string
	var inicio, fin int
	var esperado float64

	//Funcion
	//buffer := make([]byte, 1024)
	buffer := make([]byte, 32)
	mLen, err := connection.Read(buffer)
	//fmt.Println("E: ", mLen)
	if err != nil {
		fmt.Println("Error Funcion:", err.Error())
	}
	funcion = string(buffer[:mLen])

	fmt.Println("Funcion: ", funcion)

	//Inicio
	buffer = make([]byte, 2)
	mLen, err = connection.Read(buffer)
	//fmt.Println("E: ", mLen)
	if err != nil {
		fmt.Println("Error Funcion:", err.Error())
	}
	fmt.Println("Buff: ", string(buffer[:mLen]))
	inicio, _ = strconv.Atoi(string(buffer))
	fmt.Println("Inicio: ", inicio, " ", reflect.TypeOf(inicio))

	//Fin
	mLen, err = connection.Read(buffer)
	if err != nil {
		fmt.Println("Error ingresar valor:", err.Error())
	}
	fin, _ = strconv.Atoi(string(buffer))
	fmt.Println("Fin: ", fin, " ", reflect.TypeOf(fin))

	//Esperado
	buffer = make([]byte, 32)
	fmt.Println("decimal: ", buffer)

	mLen, err = connection.Read(buffer)
	if err != nil {
		fmt.Println("Error ingresar valor:", err.Error())
	}

	var aux1 int
	aux1, _ = strconv.Atoi(string(buffer[:mLen]))
	//esperado = float64(aux1) / 100
	esperado, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(aux1)/float64(100)), 64)

	println("Real esperado: ", esperado)

	var aux = Ejercicio(funcion, float64(inicio), float64(fin), esperado)

	_, err = connection.Write([]byte(aux))

	connection.Close()
}
