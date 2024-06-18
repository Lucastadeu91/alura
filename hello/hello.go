package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const check = 3
const delay = 5

func main() {

	showIntro()
	for {
		showMenu()

		actionResp := readComand()
		switch actionResp {
		case 1:
			starCheck()
		case 2:
			fmt.Println("Exibindo Logs...")
			readLog()
		case 3:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Lucas"
	version := 1.1
	fmt.Println("Olá senhor", name)
	fmt.Println("Este programa esta ma versão", version)
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(name))
}

func showMenu() {
	fmt.Println("1- Inicial monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do programa")
}

func readComand() int {
	var action int
	fmt.Scan(&action)
	fmt.Println("O Comando escolhido foi", action)
	fmt.Println("")
	return action
}

func starCheck() {
	fmt.Println("Monitoramento...")
	//sites := []string{"https://httpbin.org/status/200",	"https://www.alura.com.br", "https://www.caelum.com.br"}
	sites := readFile()
	for i := 0; i < check; i++ {
		for i, site := range sites {
			fmt.Println("testando site ", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func testSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O Site", site, "foi carregado com sucesso")
		registerLog(site, true)
	} else {
		fmt.Println("O Site", site, "esta com problemas. Status Code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func readFile() []string {

	var sites []string

	file, err := os.Open("sites.txt")
	//file, err := os.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	read := bufio.NewReader(file)
	for {
		line, err := read.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}
	file.Close()

	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro de registro ", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func readLog() {
	logFile, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao ler o arquivo ", err)
	}

	fmt.Println(string(logFile))
}
