package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	apresentacao()
}
func apresentacao() {
	fmt.Println("Bem vindo escolha uma opção")
	menu()
}
func menu() int {
	opcoes := []string{
		"1- Listar sites",
		"2- Iniciar Monitoramento",
		"3- Exibir logs",
		"4- Sair do programa",
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
	escolha(opcao)
	return opcao
}

func escolha(opcao int) {
	switch opcao {
	case 1:
		ListarSites()
	case 2:
		monitoramento()
	case 3:
		println("Exibidindologs")
	case 4:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	}

}
func ListarSites() {
	arquivo, _ := os.Open("sites.txt")
	var sites []string
	fmt.Println("---Lista de Sites Cadastrados---")
	fmt.Println("==================================")
	for {
		lerSites := bufio.NewReader(arquivo)
		site, erro := lerSites.ReadString('\n')
		site = strings.TrimSpace(site)
		sites = append(sites, site)
		for index, site := range sites {
			fmt.Println(index, "-", site)
		}
		if erro == io.EOF {
			break
		}
	}
	fmt.Println("==================================")
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

/* Criações pendentes
-criar o arquivo txt
-fazer um menu interativo que cadastre,edite,ou esclua sites
-programar consultas
-enviar via emails erros
*/
