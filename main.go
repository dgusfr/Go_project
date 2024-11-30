package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Diego"
	var idade int = 27
	var altura float32 = 1.82

  fmt.Println("Olá,", nome) 
	fmt.Println("Sua idade é de :", idade, "e sua altura:", altura)

	fmt.Println("O tipo da variavél nome é:", reflect.TypeOf((nome)))
}
