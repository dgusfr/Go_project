package main

import (
	"fmt"
	"reflect"
)

func main() {
	exibeIntroducao()

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir  Logs")
	fmt.Println("3- Sair")

	var comando int
	// fmt.Scanf("%d", &comando) or without mask with:
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi: ", comando)

	// if comando == 1{
	// 	fmt.Println("Monitoramento iniciado")
	// } else if  comando == 2  {
	// 	fmt.Println("Exibindo logs")
	// } else if comando == 3 {
	// 	fmt.Println("Saindo")
	// }else{
	// 	fmt.Println("Comando incorreto")
	// }

	switch comando {
	case 1:
		fmt.Println("Monitoramento iniciado")
	case 2:
		fmt.Println("Exibindo logs")
	case 3:
		fmt.Println("Comando incorreto")
	default:
		fmt.Println("Comando incorreto")
	}
}

func exibeIntroducao(){
	// var nome string = "Diego" it's equal to:
	nome := "Diego"
	idade := 27
	altura := 1.82

  fmt.Println("Olá,", nome) 
	fmt.Println("Sua idade é de :", idade, "e sua altura:", altura)

	fmt.Println("O tipo da iavél nome é:", reflect.TypeOf((nome)))
}

func
