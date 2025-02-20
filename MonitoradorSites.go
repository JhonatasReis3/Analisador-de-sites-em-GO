package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	HorasMonitoramento()
	//apresentacao()
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
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		println(err)
	}
	defer arquivo.Close()
	var sitesCadastrados []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		site := strings.TrimSpace(scanner.Text())
		if site != "" {
			sitesCadastrados = append(sitesCadastrados, site)
		}
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
	NumeroPordia := 24
	for i := 0; i < NumeroPordia; i++ {
		for j := 0; j < len(sites); j++ {
			resposta, _ := http.Get(sites[j])
			if resposta.StatusCode == 200 {
				fmt.Println("o site tá funcionando", sites[j])
				ExibirConsultas(sites, true)
			} else {
				fmt.Println("o site não tá funcionado", sites[j])
				ExibirConsultas(sites, false)
			}
		}
		time.Sleep(5 * time.Second)
	}
	return sites
}
func HorasMonitoramento() {
	VezesPorDia := 0
	ultimoReset := time.Now().Format("2006-01-02") // Guarda o dia atual no formato YYYY-MM-DD

	for {
		hoje := time.Now().Format("2006-01-02") // Atualiza a data de hoje

		// Se o dia mudou, zera o contador e atualiza ultimoReset
		if hoje != ultimoReset {
			VezesPorDia = 0
			ultimoReset = hoje
			fmt.Println("Novo dia detectado! Zerando contador.")
		}

		// Simula o incremento da variável durante o dia
		if VezesPorDia < 2 {
			VezesPorDia++
			fmt.Println("Executando ação. VezesPorDia:", VezesPorDia)
		}

		time.Sleep(10 * time.Second) // Simula espera antes da próxima verificação (ajuste conforme necessário)
	}
}
func ExibirConsultas(siteCadastrados []string, status bool) {
	ArquivosLogs, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err) // Tratar erro caso ocorra
	}
	agora := time.Now()
	for i := 0; i < len(siteCadastrados); i++ {
		ArquivosLogs.WriteString(agora.Format("15:04:05 02-01-2006") + " o site " + siteCadastrados[i] + " está online: " + strconv.FormatBool(status) + "\n")
	}
	defer ArquivosLogs.Close()
}

/*
- usuar a função os openfile para criar e escever em um arquivos, vamos também por um conometrô para enviar sempre os logs para um arquivo especifico
-tratar os erros
-cadastrar novo sites
-escluir ou editar sites
-função write string para escrever em arquivos
- codigo vai verificar os status do codigo sempre a cada uma hora, e ao final do dia sempre vai limpar o txt.log pra não acumular memoria

*/
