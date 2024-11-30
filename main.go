package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var nome string = "Diego" it's equal to:
	nome := "Diego"
	idade := 27
	altura := 1.82

  fmt.Println("Olá,", nome) 
	fmt.Println("Sua idade é de :", idade, "e sua altura:", altura)

	fmt.Println("O tipo da iavél nome é:", reflect.TypeOf((nome)))
}
