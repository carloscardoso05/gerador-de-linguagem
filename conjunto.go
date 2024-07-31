package main

type Conjunto[T comparable] map[T]bool

func (conjunto *Conjunto[T]) Adicionar(elemento T) {
	(*conjunto)[elemento] = true
}

func (conjunto *Conjunto[T]) Remover(elemento T) {
	delete(*conjunto, elemento)
}

func (conjunto Conjunto[T]) Contem(elemento T) bool {
	return conjunto[elemento]
}

func (conjunto1 Conjunto[T]) ContemConj(conjunto2 Conjunto[T]) bool {
	for elemento := range conjunto2 {
		if !conjunto1.Contem(elemento) {
			return false
		}
	}
	return true
}

func (conjunto1 Conjunto[T]) Igual(conjunto2 Conjunto[T]) bool {
	return conjunto1.ContemConj(conjunto2) && conjunto2.ContemConj(conjunto1)
}

func (conjunto1 Conjunto[T]) Menos(conjunto2 Conjunto[T]) Conjunto[T] {
	for k := range conjunto2 {
		conjunto1.Remover(k)
	}
	return conjunto1
}

func (conjunto1 Conjunto[T]) Uniao(conjunto2 Conjunto[T]) Conjunto[T] {
	for elemento := range conjunto2 {
		conjunto1.Adicionar(elemento)
	}
	return conjunto1
}

func (conjunto1 Conjunto[T]) Intersec(conjunto2 Conjunto[T]) Conjunto[T] {
	for elemento := range conjunto1 {
		conjunto1[elemento] = conjunto2.Contem(elemento)
	}
	return conjunto1
}
