package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	serverName    string
	serverURL     string
	tempoExecucao float64
	status        int
}

func criarListaServidores(data [][]string) []Server { // função retorna lista de servidores
	var servidores []Server
	for i, line := range data { //line = segundo vetor
		if i > 0 {
			servidor := Server{ // novo objeto da struct Server
				serverName: line[0], //indica a linha 0 do arquivo csv, isto é, primeira coluna, Server.
				serverURL:  line[1], //indica a linha 0 do arquivo csv, isto é, segunda coluna, URL.
			}
			servidores = append(servidores, servidor) // servidores, adiciona a primeira iteração do laço for
		}
	}
	return servidores //retorna lista de servidores ao final da iteração

}

func checkServer(servidores []Server) {
	for _, servidor := range servidores {
		agora := time.Now()                      //tempo
		get, err := http.Get(servidor.serverURL) //acesso a URL, através do objeto criado
		if err != nil {
			fmt.Println(err)
		}
		servidor.status = get.StatusCode
		servidor.tempoExecucao = time.Since(agora).Seconds()
		fmt.Printf("URL: [%s] Status: [%d] Tempo de carga: [%f]\n", servidor.serverURL, servidor.status, servidor.tempoExecucao)
		//fmt.Println(servidor)
	}
}

// verifica se o serviço está responsivo e tempo de resposta
// utilizando uma lista csv
func main() {

	f, err := os.Open(os.Args[1]) //primeiro precisa fazer a leitura do arquivo
	if err != nil {               //err se der erro será impresso o erro
		fmt.Println(err)
		os.Exit(1) //Sai do programa Exit(1)
	}
	defer f.Close() //fechar o arquivo quando a função for concluida

	csvReader := csv.NewReader(f)    // será passada para a variável o arquivo lido
	data, err := csvReader.ReadAll() // para ler todo o arquivo, não só a primeira linha
	if err != nil {                  //err se der erro será impresso o erro
		fmt.Println(err) //Sai do programa Exit(1)
		os.Exit(1)
	}
	servidores := criarListaServidores(data)
	checkServer(servidores)

}
