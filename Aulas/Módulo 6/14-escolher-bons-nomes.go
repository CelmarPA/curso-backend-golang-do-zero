//go:build ignore

// Escolher bons nomes para variáveis, funções e constantes é essencial para você ter um código limpo e de fácil leitura.

// Exemplo:

package main


import "fmt"

func main() {
	x := []string{"banana", "maça", "uva", "kiwi"}
	frutas := []string{"banana", "maça", "uva", "kiwi"}

	fmt.Println(x, frutas)
}

// Qual dos dois nomes (x ou frutas) deixa mais claro o conteúdo daquela variável? Claro que o nome frutas. Por isso, quando for escolher o nome de uma variável, constante ou função, pare e reflita: o que essa var guarda? o que essa função faz? E use essas respostas para dar um nome descritivo.

// Evitar repetir o nome do pacote na função/variável
// Se o nome do seu pacote é log, por exemplo, você não precisa criar uma func que também tenha esse termo log. Pois quando você chamar essa função, variável ou constante fora do pacote, acaba ficando repetitivo.

// Exemplo:

// log.Info() // ótimo
// log.LogInfo() // ruim

// Índices

// Em índices, procure ser curto. Utilize apenas uma letra:

// Exemplo:

// for i:= 0; i < 10; i++ {}

// Nomear pacotes

// Prefira nomes curtos. O nome deve ser claro e definir a função daquele pacote.
