package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		exibeLogs()
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

	numTestes := 5 // Número máximo de ondas de testes
	for i := 1; i <= numTestes; i++ {
		fmt.Printf("\nIniciando teste #%d...\n", i)
		for _, site := range sites {
			status := verificaSite(site)
			registraLog(site, status)
		}
		fmt.Println("Aguardando 5 segundos para o próximo teste...")
		time.Sleep(5 * time.Second)
	}
	fmt.Printf("\nFinalizamos %d ondas de testes.\n", numTestes)
}

func verificaSite(site string) string {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Printf("Erro ao acessar o site %s: %s\n", site, err)
		return "offline"
	}

	if resp.StatusCode == 200 {
		fmt.Printf("O site %s está online! Status Code: %d\n", site, resp.StatusCode)
		return "online"
	}

	fmt.Printf("O site %s não está acessível. Status Code: %d\n", site, resp.StatusCode)
	return "offline"
}

func registraLog(site string, status string) {
	arquivo, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de logs:", err)
		return
	}
	defer arquivo.Close()

	log := fmt.Sprintf("%s - Site: %s | Status: %s\n", time.Now().Format("2006-01-02 15:04:05"), site, status)
	arquivo.WriteString(log)
}

func exibeLogs() {
	dados, err := os.ReadFile("logs.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo de logs:", err)
		return
	}

	fmt.Println("\n=========== Logs de Monitoramento ===========")
	fmt.Println(string(dados))
	fmt.Println("=============================================")
}
