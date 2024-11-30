package main

import (
	"fmt"
	"reflect"
)

func main() {
	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitoramento iniciado")
	case 2:
		fmt.Println("Exibindo logs")
	case 3:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Comando incorreto")
	}
}

func exibeIntroducao() {
	nome := "Diego"
	idade := 27
	altura := 1.82

	fmt.Println("Olá,", nome)
	fmt.Println("Sua idade é:", idade, "e sua altura é:", altura)
	fmt.Println("O tipo da variável 'nome' é:", reflect.TypeOf(nome))
}

func exibeMenu() {
	fmt.Println("\nSelecione uma das opções abaixo:")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair")
}

func leComando() int {
	var comandoLido int
	fmt.Print("\nDigite o comando desejado: ")
	fmt.Scan(&comandoLido)
	fmt.Println("O comando lido foi:", comandoLido)
	return comandoLido
}
