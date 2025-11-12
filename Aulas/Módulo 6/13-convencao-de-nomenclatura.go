//go:build ignore

// Camel case deve começar com a primeira letra minúscula e a primeira letra de cada nova palavra subsequente maiúscula: celmarPereira / golangDoZero

// Pascal case conhecido também como “upper camel case” ou “capital case”, pascal case combina palavras colocando todas com a primeira letra maiúscula: CelmarPereira / GolangDoZero

package escola

type Aluno struct {
	Nome string
	Email string
	Notas []int
}

type professor struct {
	nome string
	email string
	materias []string
}

func Notas() []int {
	return []int{10, 9, 8, 10}
}

// Nesse exemplo, a struct Aluno (e também os campos da struct) e a função Notas() poderá ser usada fora do pacote escola, pois começam com letra maiúscula. 
// Já a struct professor é exclusiva para ser usada e acessada no pacote escola, já que começa com a letra minúscula.
