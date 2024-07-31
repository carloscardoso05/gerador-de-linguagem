package main

type Item[T any] struct {
	valor   T
	proximo *Item[T]
}

type Fila[T any] struct {
	Primeiro *Item[T]
}

func (fila *Fila[T]) Retirar() *Item[T] {
	if fila.Primeiro != nil {	
		retirado := Item[T]{
			valor:   fila.Primeiro.valor,
			proximo: nil,
		}
		
		fila.Primeiro = fila.Primeiro.proximo
		
		return &retirado
	}
	return nil
}

func (fila *Fila[T]) Adicionar(item Item[T]) {
	cursor := &fila.Primeiro
	for {
		if (*cursor == nil) {
			*cursor = &item
			return
		}
		cursor = &(*cursor).proximo
	}
}

func (fila Fila[T]) Len() int {
	i := 0
	cursor := fila.Primeiro
	for {
		if (cursor == nil) {
			return i
		}
		cursor = cursor.proximo
		i++
	}
}



