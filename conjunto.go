package main

type Conjunto[T comparable] map[T]bool

func (c *Conjunto[T]) Adicionar(elem T) {
	(*c)[elem] = true
}

func (c *Conjunto[T]) Remover(elem T) {
	delete(*c, elem)
}

func (c Conjunto[T]) Contem(elem T) bool {
	return c[elem]
}

func (c Conjunto[T]) ContemConj(conj Conjunto[T]) bool {
	for k := range conj {
		if !c.Contem(k) {
			return false
		}
	}
	return true
}

func (c Conjunto[T]) Igual(conj Conjunto[T]) bool {
	return c.ContemConj(conj) && conj.ContemConj(c)
}

func (c Conjunto[T]) Menos(conj Conjunto[T]) Conjunto[T] {
	for k := range conj {
		c.Remover(k)
	}
	return c
}

func (c Conjunto[T]) Uniao(conj Conjunto[T]) Conjunto[T] {
	for k := range conj {
		c.Adicionar(k)
	}
	return c
}

func (c Conjunto[T]) Intersec(conj Conjunto[T]) Conjunto[T] {
	for k := range c {
		c[k] = conj.Contem(k)
	}
	return c
}
