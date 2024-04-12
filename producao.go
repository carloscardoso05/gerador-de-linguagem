package main

import "strings"

type Producao struct {
	inicio string
	fim string
}

func (p Producao) Aplicar(cadeia string) string {
	return strings.Replace(cadeia, p.inicio, p.fim, 1)
}

func (p Producao) AplicarLoop(cadeia string) string {
	atual := cadeia
	prev := p.Aplicar(cadeia)
	for prev != atual{
		prev = atual
		atual = p.Aplicar(atual)
	}
	return atual
}