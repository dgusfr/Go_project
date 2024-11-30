package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
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

	// Lista de sites para monitorar
	sites := []string{
		"https://www.g1.com.br",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.uol.com.br",
	}

	// Verificar os sites continuamente com pausa de 5 segundos
	for {
		for i := 0; i < len(sites); i++ {
			verificaSite(sites[i])
		}
		fmt.Println("Aguardando 5 segundos para a próxima verificação...")
		time.Sleep(5 * time.Second)
	}
}

func verificaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Printf("Erro ao acessar o site %s: %s\n", site, err)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Printf("O site %s está online! Status Code: %d\n", site, resp.StatusCode)
	} else {
		fmt.Printf("O site %s não está acessível. Status Code: %d\n", site, resp.StatusCode)
	}
}
