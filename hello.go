package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntro()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			exibeLog()
		case 3:
			fmt.Println("saindo...")
			os.Exit(0)
		default:
			fmt.Println("não conheço essa instrução")
			os.Exit(-1)
		}
	}

}

func exibeIntro() {
	// variaveis podem ou não ter o tipo declarado.
	var nome string = "Thiago"

	// maneira mais simples de criar uma variavel
	versao := 1.1
	fmt.Println("Olá senhor,", nome)
	fmt.Println("Você esta na versão,", versao)
}
func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido

}
func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair")
}
func iniciaMonitoramento() {
	sites := leSitesdoArquivo()

	fmt.Println("Monitorando...")

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site:", i, "-", site)
			testaSite(site)
		}
		fmt.Println("")
		fmt.Println("Novo Teste:")
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")

}
func testaSite(site string) {

	// go permite mais de um retorno nas funções  | "_" foi usado para ignorar um segundo parametro obrigatorio
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	// status 200 significa que o site está funcionanto bem
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		registralog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas statuscode:", resp.StatusCode)
		registralog(site, false)
	}
}
func leSitesdoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		// aspas simples representa bytes
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}
	arquivo.Close()
	return sites
}

// Função de registro (obs: os.OpenFile foi usado pois possui mais recursos, como os de criar um arquivo caso ele não exista)
func registralog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + " - online:" + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

// Função que exibe logs (obs: "ioutil.ReadFile("log.txt")", não é necessario usar o close(), pois o pacote já faz isso automaticamente )
func exibeLog() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
