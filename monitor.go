package main // Este pacote contém um programa simples que imprime uma mensagem de boas-vindas e a versão do programa.

import (
	"bufio"    // O pacote `bufio` fornece funções para trabalhar com fluxos de entrada e saída.
	"fmt"      // O pacote `fmt` fornece funções para formatação e impressão de texto.
	"io"       // O pacote `io` fornece funções para trabalhar com fluxos de entrada e saída.
	"net/http" // O pacote `net/http` fornece funções para enviar e receber solicitações HTTP.
	"os"       // O pacote `os` fornece funções para interagir com o sistema operacional.
	"strconv"  // O pacote strconv fornece funções para converter valores entre formatos de string e formatos de máquina.
	"strings"  // O pacote strings fornece funções para trabalhar com strings.
	"time"     // O pacote `time` fornece funções para trabalhar com datas, horas e intervalos de tempo. Ele também fornece funções para manipular o relógio do sistema.
)

const monitoramentos = 3
const delay = 6

func main() { // Este programa monitora um conjunto de sites e registra seu status em um arquivo de log.
	exibeIntrodução()

	for { // Usar a estrutura for sem parâmetro funciona como loop infinito.
		exibeMenu()

		comando := leComando() // O retorno de leComando é enviado para a variável comando e após para o switch.

		switch comando { //  Usar a estrutura `switch` quando esperar mais de um resultado e usar if quando esperar apenas um resultado lógico.
		case 0:
			fmt.Println("Iniciando Monitoramento...")
			fmt.Println("---------------------------------")
			iniciarMonitoramento()
		case 1:
			fmt.Println("Preparando logs...")
			time.Sleep(delay * time.Second) // Cria uma pausa pé definida, para a exibição de cada loop.
			ExibirLogs()
		case 2:
			fmt.Println("Programa encerrado")
			os.Exit(0) // A função `os.Exit(0)` indica que o programa foi encerrado com sucesso.
		default:
			fmt.Println("Comando não encontrado")
			os.Exit(-1) // A função `os.Exit(-1)` indica que o programa foi encerrado com um erro.
		}
	}
}

func exibeIntrodução() { // Função criada para exibir uma mensagem de boas-vindas ao usuário e a versão do programa.
	nome := "Gilnei" // A variável `nome` é o nome do usuário.
	versao := 1.1    // A variável `versao` é a versão do programa.
	fmt.Println("---------------------------------")
	fmt.Println("Olá, sr.", nome) // A função `fmt.Println()` imprime uma mensagem na tela.
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() { // Função criada para a exibição do menu do programa.
	fmt.Println("=================================")
	fmt.Println("=   [0] Iniciar Monitoramento   =")
	fmt.Println("=   [1] Exibir Logs             =")
	fmt.Println("=   [2] Encerrar Programa       =")
	fmt.Println("=================================")
}

func leComando() int { // Função criada para ler o comando do usuário e retornar um inteiro.
	var comandoLido int // A variável `comando` armazena o comando escolhido pelo usuário.

	fmt.Scan(&comandoLido) // Lê o comando do usúario.
	fmt.Println("O comando escolhido foi: [", comandoLido, "]")
	return comandoLido // O resultado do comando retorna para o a função leComando para chegar no switch pela váriavel comando.
}

func iniciarMonitoramento() { // Função criada para enviar uma solicitação HTTP para a URL especificada e verifica o status da resposta.
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites { // Executando a função testaSite() para cada uma das strings no slice sites.
			testaSite(site) // Recebe o site já com a resposta.
		}

		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) { // Recebe uma string como parâmetro e inicia o teste.
	resp, err := http.Get(site) // `resp` é a resposta HTTP da solicitação.

	if err != nil { // Realiza o tratamento de erros usando uma condição.
		fmt.Println("Ocorreu um erro.\n", err)
	}

	if resp.StatusCode == 200 { // Realiza a condição de acordo com o status code. // `StatusCode` é o código de status da resposta.
		fmt.Println(site, ": Site em funcionamento.")
		registraLog(site, true) // Se o site estiver no ar(status 200), sera enviado `true` para a função de registro de logs.
	} else {
		fmt.Println(site, ": Fora do ar no momento.", "Status Code:", resp.StatusCode)
		registraLog(site, false) // Se o site não estiver no ar por qualquer tipo de erro, sera enviado `false` para a função de registro de logs.
	}
}

func leSitesDoArquivo() []string { // Função criada para ler o arquivo especificado.
	var sites []string

	arquivo, err := os.Open("sites.txt") // lê a posição do arquivo.

	if err != nil {
		fmt.Println("Ocorreu um erro.\nNão foram encontardos sites para monitorar.\n", err)
	}

	leitor := bufio.NewReader(arquivo) // Lê linha por linha do arquivo.

	for {
		linha, err := leitor.ReadString('\n') // Transforma a linha em uma string.
		linha = strings.TrimSpace(linha)      // Função que tira espaços e quebras do arquivo.

		sites = append(sites, linha) // Adiciona cada string do arquivo em um slice chamado `sites`.

		if err == io.EOF {
			break // Fecha o arquivo após um erro de leitura.
		}
	}

	arquivo.Close() // Fechando o arquivo após usar.

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("[ERRO]:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 - 15:4:5 - ") + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func ExibirLogs() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Erro: ", err)
	}

	fmt.Println(string(arquivo))
}
