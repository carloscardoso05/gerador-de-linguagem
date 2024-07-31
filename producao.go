package main

import "strings"

type Producao struct {
	inicio string
	fim string
}

func (producao Producao) Aplicar(cadeia string) string {
	return strings.Replace(cadeia, producao.inicio, producao.fim, 1)
}

func (producao Producao) AplicarToda(cadeia string) string {
	atual := cadeia
	prev := producao.Aplicar(cadeia)
	for prev != atual{
		prev = atual
		atual = producao.Aplicar(atual)
	}
	return atual
}