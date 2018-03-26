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

const monitoramentos = 5
const delay = 5

func main() {
	exibeIntroducao()
	for {

		exibeMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Lendo Logs")
			imprimirLogs()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Comando errado")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	fmt.Println("Quem você é?")
	var nome string
	fmt.Scan(&nome)
	versao := 1.1
	fmt.Println("Olá ", nome)
	fmt.Println("Este programa esta na versao ", versao)
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func exibeMenu() {
	fmt.Println("#***************************#")
	fmt.Println("# 1 - Iniciar monitoramento #")
	fmt.Println("# 2 - Exibir logs           #")
	fmt.Println("# 0 - Sair                  #")
	fmt.Println("#***************************#")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// sites := []string{
	// 	"https://www.alura.com.br",
	// 	"https://random-status-code.herokuapp.com/",
	// 	"https://www.caelum.com.br",
	// 	"https://freelinha-landing-page.herokuapp.com/"}

	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println("Site: ", sites[i], ", ok")
	// }

	sites := lerSitesDoArquivo()

	for j := 0; j < monitoramentos; j++ {
		for i, site := range sites {
			fmt.Println("Tentado site: ", i, ": ", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, ", ok", " Status: ", resp.StatusCode)
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, ", esta down", " Status: ", resp.StatusCode)
		registraLog(site, false)
	}
}

func lerSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Erro:", err)
		os.Exit(0)
	}
	// arquivo, err := ioutil.ReadFile("sites.txt")

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	hora := time.Now().Format("02/01/2006 15:04:05")

	arquivo.WriteString(hora + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimirLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println(string(arquivo))
}
