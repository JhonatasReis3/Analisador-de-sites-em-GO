package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	apresentacao()
}
func apresentacao() {
	fmt.Println("---Bem vindo escolha uma opção---")
	opcao := menuBoasvindas()
	fmt.Println(opcao)
}
func menuBoasvindas() int {
	opcoes := []string{
		"1-Listar sites",
		"2-Iniciar Monitoramento",
		"3-Exibir logs",
		"4-Sair do programa",
	}
	var opcao int
	sitesCadastrados := ObterSites()
	for {
		for i := 0; i < len(opcoes); i++ {
			fmt.Println(opcoes[i])
		}
		fmt.Scan(&opcao)
		if opcao != 0 {
			break
		}
	}
	escolha(opcao, sitesCadastrados)
	return opcao
	/*
		o erro está que a variavel opcao está compilando o numero do comando, preciso resolver
	*/
}

func escolha(opcao int, sitesCadastrados []string) {
	switch opcao {
	case 1:
		ListarSites(sitesCadastrados)
	case 2:
		monitoramento(sitesCadastrados)
	case 3:
		println("Exibidindologs")
	case 4:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	}
}
func ObterSites() []string {
	arquivo, _ := os.Open("sites.txt")
	defer arquivo.Close()
	var sitesCadastrados []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		site := strings.TrimSpace(scanner.Text())
		if site != "" {
			sitesCadastrados = append(sitesCadastrados, site)
		}
	}

	if err := scanner.Err(); err != nil {
		return sitesCadastrados
	}

	return sitesCadastrados
}
func ListarSites(sitesCadastrados []string) {
	fmt.Println("---Lista de Sites Cadastrados---")
	for index, site := range sitesCadastrados {
		fmt.Println(index, "-", site)
	}
	fmt.Println("=======================")
	MenuSites(sitesCadastrados)
}
func MenuSites(siteCadastrados []string) {
	println("---Oque desejas fazer agora---")
	var OpcoesMenusite = []string{
		"1- Monitorar Sites",
		"2- Exibir Logs",
		"3- Cadastrar no Site",
		"4- Editar Sites",
		"5- Voltar ao menu inicial",
	}
	var EscolhaMenusite int
	for i := 0; i < len(OpcoesMenusite); i++ {
		fmt.Println(OpcoesMenusite[i])
	}
	fmt.Scan(&EscolhaMenusite)
	switch EscolhaMenusite {
	case 1:
		monitoramento(siteCadastrados)
	}
}
func monitoramento(sitesCadastrados []string) []string {
	sites := sitesCadastrados
	for j := 0; j < len(sites); j++ {
		resposta, _ := http.Get(sites[j])
		if resposta.StatusCode == 200 {
			fmt.Println("o site tá funcionando", sites[j])
		} else {
			fmt.Println("o site não tá funcionado", sites[j])
		}
	}
	return sites
}
