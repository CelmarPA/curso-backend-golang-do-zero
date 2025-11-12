//go:build ignore

package main

import "fmt"

type Animal interface {
	Barulho() string
	NumeroDePatas() int
}

type Cachorro struct {
	Raca  string
	Cor   string
	Patas int
}

type Gato struct {
	Cor   string
	Patas int
}

func (c Cachorro) Barulho() string {
	return "Au au!"
}

func (c Cachorro) NumeroDePatas() int {
	return c.Patas
}

func (g Gato) Barulho() string {
	return "Miau!"
}

func (g Gato) NumeroDePatas() int {
	return g.Patas
}

func FazBarulho(animal Animal) {
	fmt.Println(animal.Barulho())
}

type Pessoa struct {
	Nome             string
	Idade            int
	Prof             string
	TempoDeProfissao int
}

type Crianca struct {
	Nome  string
	Idade int
}

func (c Crianca) FalaBomDia() string {
	return c.Nome + " deseja bom dia para você!"
}

func (p Pessoa) FalaBomDia() string {
	return fmt.Sprintf("%s deseja bom dia para você!", p.Nome)
}

func (p Pessoa) Profissao() string {
	return fmt.Sprintf("%s tem %d anos como %s", p.Nome, p.TempoDeProfissao, p.Prof)
}

type Adulto interface {
	FalaBomDia() string
	Profissao() string
}

func BomDia(adulto Adulto) string {
	return adulto.FalaBomDia()
}

func main () {

	adulto := Pessoa {
		Nome: "Celmar",
		Idade: 35,
		Prof: "dev",
		TempoDeProfissao: 0,
	}

	criança := Crianca {
		Nome: "Bento",
		Idade: 4,
	}

	cachorro := Cachorro {
		Raca: "spitz alemão",
		Cor: "preto",
		Patas: 4,
	}

	gato := Gato{
		Cor: "branco",
		Patas: 4,
	}
	
	fmt.Println(criança.FalaBomDia())
	fmt.Println(adulto.FalaBomDia())
	fmt.Println(BomDia(adulto))
	
	fmt.Println(cachorro.Barulho())
	fmt.Println(cachorro.NumeroDePatas())
	fmt.Println(gato.Barulho())
	fmt.Println(gato.NumeroDePatas())

	FazBarulho(cachorro)
	FazBarulho(gato)
}