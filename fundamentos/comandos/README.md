# Comandos

Nesta aula estudamos alguns dos principais comandos da linguagem usados durante o desenvolvimento de aplicações com Golang.

 - ```go version```: Exibe a versão do Golang.
 - ```go help```: Exibe uma ajuda sobre os comandos do Go, passando o nome de um comando como argumento é possível detalhar a ajuda de um comando específico.
 - ```go env```: Exibe as variáveis de ambiente do Go, um argumento pode ser passado para exibir o valor de uma variável específica.
 - ```go run```: Compila e executa o código de um programa.
 - ```go build```: Trata-se de um comando bastante complexo com  várias flags opcionais, mas basicamente ele é responśavel pela compilação de um programa.
 - ```go mod```: Provê acesso aos comandos de operações de **módulos** como inicialização, download, verificação, etc. Usando o comando ```go help mod``` é possível ter mais detalhes. O comando ```go mod tidy``` por exemplo, realiza a adequação das importações de módulos de um projeto removendo e adcionando módulos não usados ou faltantes, respectivamente.
 - ```go vet```: Examina o código e reporta construções suspeitas podendo até encontrar erros que podem passar pelos compiladores.
