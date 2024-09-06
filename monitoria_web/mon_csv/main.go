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
	dataFalha     string
}

func criarListaServidores(serverList *os.File) []Server { // função retorna lista de servidores
	csvReader := csv.NewReader(serverList)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var servidores []Server
	for i, line := range data { //line = segundo vetor
		if i > 0 {
			servidor := Server{ // novo objeto da struct Server
				serverName: line[0],
				serverURL:  line[1],
			}
			servidores = append(servidores, servidor) // servidores: adiciona ao final, a primeira iteração do laço for
		}
	}
	return servidores //retorna lista de servidores ao final da iteração

}

// Verifica e retorna o status dos servidores. Status esperado é sempre o 200
func checkServer(servidores []Server) []Server {
	var downServers []Server //slice

	for _, servidor := range servidores {
		agora := time.Now()                      //tempo
		get, err := http.Get(servidor.serverURL) //acesso a URL, através do objeto criado
		if err != nil {
			fmt.Printf("Server %s is down [%s]\n", servidor.serverName, err.Error())
			servidor.status = 0
			servidor.dataFalha = agora.Format("02/01/2006 15:04:05")
			downServers = append(downServers, servidor)
			continue
		}
		servidor.status = get.StatusCode
		if servidor.status != 200 {
			servidor.dataFalha = agora.Format("02/01/2006 15:04:05")
			downServers = append(downServers, servidor)
		}
		servidor.tempoExecucao = time.Since(agora).Seconds()
		fmt.Printf("URL: [%s] Status: [%d] Tempo de carga: [%f]\n", servidor.serverURL, servidor.status, servidor.tempoExecucao)
	}
	return downServers //retorno do slice
}

// Faz a leitura do arquivo csv de servidores
// e cria um outro arquivo csv com o retorno
func openFiles(serverListFile string, downtimeFile string) (*os.File, *os.File) {
	serverList, err := os.OpenFile(serverListFile, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	downtimeList, err := os.OpenFile(downtimeFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return serverList, downtimeList
}

func generateDowntime(downtimeList *os.File, downServers []Server) {
	csvWriter := csv.NewWriter(downtimeList)
	for _, servidor := range downServers {
		line := []string{servidor.serverName, servidor.serverURL, servidor.dataFalha, fmt.Sprintf("%f", servidor.tempoExecucao), fmt.Sprintf("%d", servidor.status)}
		csvWriter.Write(line)
	}
	csvWriter.Flush() //quando a função terminar, se der erro, ele passa todos os dados em memória para o arquivo. Assim não perde o histórico.
}

func main() {
	serverList, downtimeList := openFiles(os.Args[1], os.Args[2])
	defer serverList.Close()
	defer downtimeList.Close()
	servidores := criarListaServidores(serverList)
	downServers := checkServer(servidores)
	generateDowntime(downtimeList, downServers)
}
