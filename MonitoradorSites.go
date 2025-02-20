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
	fmt.Println("===Bem vindo escolha uma opção===")
	menuBoasvindas()

}
func menuBoasvindas() {
	CabecalhoBoasVindas := []string{
		"1-Listar sites",
		"2-Iniciar Monitoramento",
		"3-Exibir logs",
		"4-Sair do programa",
	}
	var EscolhaCabecalho int
	sitesCadastrados := ObterSites()
	for {
		for i := 0; i < len(CabecalhoBoasVindas); i++ {
			fmt.Println(CabecalhoBoasVindas[i])
		}
		fmt.Println("================================")
		fmt.Scan(&EscolhaCabecalho)
		if EscolhaCabecalho >= 1 && EscolhaCabecalho <= len(CabecalhoBoasVindas) {
			break
		}
	}
	escolha(EscolhaCabecalho, sitesCadastrados)
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
	fmt.Println("===Lista de Sites Cadastrados===")
	for index, site := range sitesCadastrados {
		fmt.Println(index, "-", site)
	}
	fmt.Println("================================")
	MenuSites(sitesCadastrados)
}
func MenuSites(siteCadastrados []string) {
	println("===Oque desejas fazer agora===")
	var OpcoesMenusite = []string{
		"1- Monitorar Sites",
		"2- Exibir Logs",
		"3- Cadastrar novo Site",
		"4- Editar Sites",
		"5- Voltar ao menu inicial",
	}
	var EscolhaMenusite int
	for i := 0; i < len(OpcoesMenusite); i++ {
		fmt.Println(OpcoesMenusite[i])
	}
	fmt.Println("================================")
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

/*
- usuar a função os openfile para criar e escever em um arquivos, vamos também por um conometrô para enviar sempre os logs para um arquivo especifico
-tratar os erros
-cadastrar novo sites
-escluir ou editar sites
-função write string para escrever em arquivos
- codigo vai verificar os status do codigo sempre a cada uma hora, e ao final do dia sempre vai limpar o txt.log pra não acumular memoria
*/
