### Instalação do Go - Windows
- https://go.dev/dl/
    - Instalação Padrão (next, next, next)

- No VS Code, instalar a extensão Go
- Instalar extensões auxiliares
    - CTRL + SHIFT + P
    - Go: Install/Update Tools
    - Selecione todas
    - Aguarde o fim da instalação (SUCCEEDED)
    - Fecha a IDE e abra novamente

### Iniciar go.mod 

Criar dirtório do projeto
- abrir o terminal e inciar o módulo do Go `go mod init nome_do_modulo`
  - ex.: go mod init gohello

Executar o Go:

Opção 1:
 - go run main.go

Opção 2:
Com o nome do módulo definido no comando go mod init nome_do_modulo. Esse comando gera o BUILD executável do projeto criado, que nesse caso será o `gohello.exe`.

- go build
- ./gohello

### Autenticador do http server

Github: https://github.com/abbot/go-http-auth

no terminal digitar: 
```bash
go get github.com/abbot/go-http-auth
```

Password Hash Generator
Link: https://unix4lyfe.org/crypt/

### Monitoria URL com CSV

passagem de parâmetros para execução do código
```bash
go run main.go lista-servidores.csv downtime.csv
```

