package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

//func main() {
//	fs := http.FileServer(http.Dir("C:/GoServer/"))
//	//http.ListenAndServe(":9090", fs)
//	//log.Fatal informa o log
//	fmt.Println("Subindo servidor na porta 9090")
//	log.Fatal(http.ListenAndServe(":9090", fs))
//}

//func main() {
//	if len(os.Args) != 3 {
//		fmt.Println("Uso: go run main.go <diretorio> <porta>")
//		os.Exit(1)
//	}
//	httpDir := os.Args[1]
//	porta := os.Args[2]
//	fs := http.FileServer(http.Dir(httpDir))
//	fmt.Printf("Subindo servidor na porta %s ...", porta)
//	log.Fatal(http.ListenAndServe(":"+porta, fs))
//
//}

func Secret(user, realm string) string {
	if user == "marcos" {
		return "$1$7Yz6Nw4V$N0F28UI8tj2otj1RRaKG8/"
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretorio> <porta>")
		os.Exit(1)
	}
	httpDir := os.Args[1]
	porta := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("meuserver.com", Secret)
	//Função que associa uma URL e protege o que vem antes da "/"
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))
	fmt.Printf("Subindo servidor na porta %s ...", porta)
	log.Fatal(http.ListenAndServe(":"+porta, nil))

}
