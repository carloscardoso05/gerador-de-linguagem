package main

import (
	"strings"
)

type Gramatica struct {
	Vocabulario Conjunto[string]
	Terminais   Conjunto[string]
	Producoes   Conjunto[Producao]
	Raiz        string
}

func ehFormalFinal(valor string, terminais Conjunto[string]) bool {
	for _, c := range valor {
		if !terminais.Contem(string(c)) {
			return false
		}
	}
	return true
}

func (gramatica Gramatica) NaoTerminais() Conjunto[string] {
	return gramatica.Vocabulario.Menos(gramatica.Terminais)
}

func (gramatica Gramatica) Linguagem(limite uint) []string {
	fila := Fila[string]{}
	fila.Adicionar(Item[string]{valor: gramatica.Raiz})
	var linguagem []string

	for i := uint(0); i < limite; i++ {
		atual := fila.Retirar()
		if atual == nil {
			return linguagem
		}

		for prod := range gramatica.Producoes {
			if strings.Contains(atual.valor, prod.inicio) {
				res := prod.Aplicar(atual.valor)
				if ehFormalFinal(res, gramatica.Terminais) {
					linguagem = append(linguagem, res)
				} else {
					fila.Adicionar(Item[string]{valor: res})
				}
			}
		}
	}
	return linguagem
}

func (gramatica Gramatica) GerarLinguagem(linguagem chan string) {
	fila := make(chan Item[string])
	fila <- Item[string]{valor: gramatica.Raiz}

	for atual := range fila {
		for prod := range gramatica.Producoes {
			go aplicaProd(prod, atual.valor, gramatica, linguagem, fila)
		}
	}
}

func aplicaProd(prod Producao, valor string, gramatica Gramatica, linguagem chan string, fila chan Item[string]) {
	if strings.Contains(valor, prod.inicio) {
		res := prod.Aplicar(valor)
		if ehFormalFinal(res, gramatica.Terminais) {
			linguagem <- res
		} else {
			fila <- Item[string]{valor: res}
		}
	}
}

func (gramatica Gramatica) ReconheceCadeia(cadeia string) bool {
	// for _, a := range gramatica.Linguagem(limite) {
	// 	println(a)
	// }
	// return slices.Contains(gramatica.Linguagem(limite), cadeia)
	linguagem := make(chan string)

	gramatica.GerarLinguagem(linguagem)

	for s := range linguagem {
		if s == cadeia {
			return true
		}
	}
	return false
}
