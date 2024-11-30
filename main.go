package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
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

func iniciarMonitoramento() {
	fmt.Println("Monitoramento iniciado...")
	
	url := "https://www.g1.com.br"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Erro ao acessar o site:", err)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Printf("O site %s está online! Status Code: %d\n", url, resp.StatusCode)
	} else {
		fmt.Printf("O site %s não está acessível. Status Code: %d\n", url, resp.StatusCode)
	}
}
