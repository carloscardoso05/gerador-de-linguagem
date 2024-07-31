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

func (gramatica Gramatica) NaoTerminais() Conjunto[string] {
	return gramatica.Vocabulario.Menos(gramatica.Terminais)
}

func (gramatica Gramatica) Linguagem(limite uint) []string {
	fila := Fila[string]{}
	fila.Adicionar(Item[string]{valor: gramatica.Raiz})
	res := make([]string, limite)

	for i := 0; i < int(limite); i++ {
		atual := fila.Retirar()
		if atual == nil {
			return res
		}
		res[i] = atual.valor
		for p := range gramatica.Producoes {
			if strings.Contains(atual.valor, p.inicio) {
				fila.Adicionar(Item[string]{valor: p.Aplicar(atual.valor)})
			}
		}
	}
	return res
}
