[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-named-return-functions/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-named-return-functions/blob/main/README.md)  
go version 1.22.1  

## Panic e Recover
Muito utilizadas com o ```defer```, vamos usar o exemplo de que calcula a média de notas, mas a média nunca pode ser 6, se ela for 6 nosso programa ira entrar em "pânico" usando a clausula ```panic```, o que o ```panic``` faz é interomper o fluxo normal do seu programa, parando tudo, para de executar tudo e "entra em pânico" termo usado no go.  
Basicamente ela vai chamar todas as funções que tem  o ```defer```, caso você não recupere a função com ```recover```, que veremos a diante, seu programa morre.  

### Iniciando o projeto
Crie um deretório ```go-panic-recover``` com um arquivo ```main.go``` contendo:  
```go
package main

func studentApproved(n1, n2 float64) bool {
	average := (n1 + n2) / 2
	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}
  //em caso de ser igual a 6 seu app entra em "pânico"
	panic("A média é exatamente 6")
}

func main() {
fmt.Println(studentApproved(6, 7)) //outuput true
	fmt.Println(studentApproved(3, 7)) //outuput false
	// caso a media seja 6 e panic seja chamada o último print, após ela, não será executado
	fmt.Println(studentApproved(6, 6)) //outuput panic: A média é exatamente 6 ..... errors
	fmt.Println("Pós execução")
}
```
Aqui vemos que o programa entrou em pânico, pois não esperamos o número 6 como média, e o programa não sabe o que fazer, diferente de ```erro```, que você pode continuar e tratar, o ```panic``` mata a execução do programa, então caso você tenha a clausula ```panic``` e não possui o ```recover```, seu programa morre.  
Mas há uma maneira de recuperar seu programa caso ele entre em ```panic```, que é com a clausula ```recover```, exemplo:  
```go
.....
func studentApproved(n1, n2 float64) bool {
	defer recoverExec() //<<
	average := (n1 + n2) / 2
	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	//em caso de ser igual a 6 seu app entra em "pânico"
	panic("A média é exatamente 6")
}
// função para recuperar a execução do projeto
func recoverExec() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func main() {
	fmt.Println(studentApproved(6, 7)) //outuput true
	fmt.Println(studentApproved(3, 7)) //outuput false
	// caso a media seja 6 e panic seja chamada o último print, após ela, não será executado
	fmt.Println(studentApproved(6, 6)) //outuput panic é ignorada e retomamos com recover
	fmt.Println("Pós execução")
}
```
Vimos então que a função ```recover()```, pode recuperar a execução do nosso programa.  
Entendendo melhor ao adicionar a clausula ```defer``` na chamada da função ```studentApproved(...)``` Quando ocorre o ```panic```, chamamos sempre as funções defer que no nosso caso é ```recoverExec()```, nela verificamos se ```recover()``` conseguiu recuperar algo, portante é != nil nesse caso, caso retorne nil não conseguimos recuperar.  

