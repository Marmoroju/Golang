package main

import (
	"aletmanager/email"
)

func main() {

	email.SendEmail([]string{"golanggoalerta@gmail.com"},
		"Alerta de Servidor Down", "Google",
		"Erro ao conectar no servidor", "06/09/2024 15:14", "./email/template.html")
}
