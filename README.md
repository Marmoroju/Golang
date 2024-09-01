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

Imports e saídas - Básico

```bash
package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Args = array
	//ARGS [0,1] = números de argumentos/variáveis
	//Posição 0 do array é o próprio arquivo
	// Posição 1, será passada como argumento no terminal e será concatenada
	// ex.: go run main.go Oi
	// Saída: Hello, Oi
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("Hello,", os.Args[1])
}
```


