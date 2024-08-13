package main

import "fmt"

func main() {
	t := Conjunto[string]{}
	t.Adicionar("a")
	t.Adicionar("b")
	v := Conjunto[string]{}
	v = v.Uniao(t)
	v.Adicionar("S")
	p := Conjunto[Producao]{}
	p.Adicionar(Producao{inicio: "S", fim: "abA"})
	p.Adicionar(Producao{inicio: "S", fim: "a"})
	p.Adicionar(Producao{inicio: "S", fim: "b"})
	p.Adicionar(Producao{inicio: "S", fim: "bbA"})
	p.Adicionar(Producao{inicio: "S", fim: "baA"})
	p.Adicionar(Producao{inicio: "A", fim: ""})
	p.Adicionar(Producao{inicio: "A", fim: "bA"})
	p.Adicionar(Producao{inicio: "A", fim: "aA"})

	a := Gramatica{
		Raiz:        "S",
		Terminais:   t,
		Vocabulario: v,
		Producoes:   p,
	}

	// linguagem := a.Linguagem(50)
	// for _, prod := range linguagem {
	// 	fmt.Println(prod)
	// }
	cad := "babaaaaaa"
	fmt.Printf("Reconhece '%s': %t", cad, a.ReconheceCadeia(cad))
}
