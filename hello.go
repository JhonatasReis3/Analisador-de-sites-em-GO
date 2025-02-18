package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	opcao := menu()
	escolha(opcao)
}
func menu() int {
	opcoes := []string{
		"1- Iniciar Monitoramento",
		"2- Exibir logs",
		"3- Sair do programa",
	}
	var opcao int
	nome := "douglas"

	for {
		fmt.Println("Olá Sr.", nome, "Escolha um numero do menu:")
		for i := 0; i < len(opcoes); i++ {
			fmt.Println(opcoes[i])
		}
		fmt.Scan(&opcao)
		if opcao != 0 {
			break
		}
	}
	return opcao
}

func escolha(opcao int) {
	switch opcao {
	case 1:
		monitoramento()
	case 2:
		fmt.Println("Exibindo logs:")
	case 3:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	}
}
func monitoramento() {
	fmt.Println("Digite o site")
	corpo := "https://"
	var caminho string
	fmt.Scan(&caminho)
	site := corpo + caminho
	resposta, _ := http.Get(site)
	fmt.Println(resposta)
}
