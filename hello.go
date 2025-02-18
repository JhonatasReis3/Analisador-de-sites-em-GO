package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	apresentacao()
	opcao := menu()
	escolha(opcao)

}
func apresentacao() {
	var nome string
	fmt.Println("Digite seu nome:")
	fmt.Scan(&nome)
	fmt.Println("Olá SR.", nome, "Seja bem vindo ao seu monitorador de sites")
	fmt.Println("Oque deseja fazer hoje?")
	fmt.Println("")
	fmt.Println("Digite só uma opção")
}
func menu() int {
	opcoes := []string{
		"1- Iniciar Monitoramento",
		"2- Exibir logs",
		"3- Sair do programa",
	}
	var opcao int
	for {
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
func monitoramento() []string {
	fmt.Println("Digite quantos sites voce quer monitora *apenas numeros")
	var numero_sites int
	fmt.Scan(&numero_sites)
	sites := []string{}
	for i := 0; i < numero_sites; i++ {
		fmt.Printf("digite o site")
		var site string
		fmt.Scan(&site)
		sites = append(sites, site)
	}
	testando_sites(sites)
	reiniciar()
	return sites
}
func testando_sites(sites []string) {
	for j := 0; j < len(sites); j++ {
		resposta, _ := http.Get(sites[j])
		if resposta.StatusCode == 200 {
			fmt.Println("o site tá funcionando", sites[j])
		} else {
			fmt.Println("o site não tá funcionado", sites[j])
		}
	}
}
func reiniciar() string {
	var reiniciar string
	fmt.Println("Deseja fazer uma nova consulta?")
	fmt.Println("1- SIM")
	fmt.Println("2- NÃO")
	fmt.Scan(&reiniciar)
	resposta := strings.ToUpper(reiniciar)
	if resposta == "SIM" || resposta == "1" {
		monitoramento()
	} else if resposta == "NÃO" || resposta == "2" {
		menu()
	} else {

	}
	return resposta
}
