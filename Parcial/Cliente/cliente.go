package main

import (
	"fmt"
	"net"
	"strings"
)

const (
	//SERVER_HOST = "localhost"
	SERVER_HOST = "172.17.0.2"
	SERVER_PORT = "5000"
	SERVER_TYPE = "tcp"
)

func main() {
	var funci, inic, fin, esperado string
	//establish connection
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	} ///send some data
	fmt.Println("Funcion:")
	fmt.Scanln(&funci)

	fmt.Println("Desde:")
	fmt.Scanln(&inic)

	fmt.Println("Hasta:")
	fmt.Scanln(&fin)

	fmt.Println("Resultado esperado(2 decimales):")
	fmt.Scanln(&esperado)
	esperado = strings.ReplaceAll(string(esperado), ".", "")
	fmt.Println("Esperado: ", esperado)

	_, err = connection.Write([]byte(funci))
	_, err = connection.Write([]byte(inic))
	_, err = connection.Write([]byte(fin))
	_, err = connection.Write([]byte(esperado))

	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)

	if err != nil {
		fmt.Println("Error en la lectura:", err.Error())
	}
	fmt.Println("Resultado: ", string(buffer[:mLen]))

	connection.Close()
}
