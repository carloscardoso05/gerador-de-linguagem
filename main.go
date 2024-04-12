package main

func main() {
	t := Conjunto[string]{}
	t.Adicionar("a")
	t.Adicionar("b")
	v := Conjunto[string]{}
	v = v.Uniao(t)
	v.Adicionar("S")
	p := Conjunto[Producao]{}
	p.Adicionar(Producao{inicio: "S", fim: "bab"})
	p.Adicionar(Producao{inicio: "S", fim: "Sa"})
	a := Gramatica{
		Raiz:        "S",
		Terminais:   t,
		Vocabulario: v,
		Producoes:   p,
	}

	a.Linguagem()
}
