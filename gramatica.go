package main

import (
	"fmt"
	"strings"
)

type Gramatica struct {
	Vocabulario Conjunto[string]
	Terminais   Conjunto[string]
	Producoes   Conjunto[Producao]
	Raiz        string
}

func (g Gramatica) NaoTerminais() Conjunto[string] {
	return g.Vocabulario.Menos(g.Terminais)
}

func (g Gramatica) Linguagem() {
	fila := Fila[string]{}
	fila.Adicionar(Item[string]{valor: g.Raiz})

	for i := 0; i < 50; i++ {
		atual := fila.Retirar()
		if atual == nil {
			return
		}
		fmt.Println(atual.valor)
		for p := range g.Producoes {
			if strings.Contains(atual.valor, p.inicio) {
				fila.Adicionar(Item[string]{valor: p.Aplicar(atual.valor)})
			}
		}
	}
}
