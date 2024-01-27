package main

import (
	"fmt"
)

// Esta es una función principal que se ejecuta al iniciar el programa
func main() {
	// Imprimir en la consola
	fmt.Println("Hola, mundo!")

	// variables
	var nombre string = "Juan" // Declaración explícita de tipo de dato
	edad := 25                 // Inferencia de tipo, eso quiere decir que el compilador infiere el tipo de dato

	// Condicionales (if, else if, else)
	if edad >= 18 {
		fmt.Println(nombre, "es mayor de edad")
	} else if edad < 0 {
		fmt.Println("La edad no puede ser menor que cero")
	} else if edad > 110 {
		fmt.Println("La edad no puede ser mayor a 110 años")
	} else {
		fmt.Println(nombre, "es menor de edad")
	}

	// For
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// While
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}

	// Arrays
	numeros := [5]int{1, 2, 3, 4, 5}

	// Iterar sobre un array
	for index, valor := range numeros {
		fmt.Println("Índice:", index, "Valor:", valor)
	}

	// Slice (parte) de un array
	slice := numeros[1:3]
	fmt.Println("Slice:", slice)

	// Mapas (diccionarios)
	puntajes := map[string]int{
		"Juan":  95,
		"María": 80,
		"Pedro": 70,
	}

	// Iterar sobre un mapa
	for nombre, puntaje := range puntajes {
		fmt.Println(nombre, ":", puntaje)
	}

	// Funciones
	resultado := sumar(10, 20)
	fmt.Println("Resultado de la suma:", resultado)

	// Funciones anónimas
	func() {
		fmt.Println("Hola desde una función anónima")
	}()
}

// Esta es una función personalizada para sumar dos números enteros
func sumar(a, b int) int {
	return a + b
}

// Estructuras
type Persona struct {
	Nombre string
	Edad   int
	Género string
}

// Interfaces
type Animal interface {
	Comer()
	Dormir()
}
