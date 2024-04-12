package main

type Item[T any] struct {
	valor   T
	proximo *Item[T]
}

type Fila[T any] struct {
	Primeiro *Item[T]
}

func (f *Fila[T]) Retirar() *Item[T] {
	if f.Primeiro != nil {	
		retirado := Item[T]{
			valor:   f.Primeiro.valor,
			proximo: nil,
		}
		
		f.Primeiro = f.Primeiro.proximo
		
		return &retirado
	}
	return nil
}

func (f *Fila[T]) Adicionar(item Item[T]) {
	cursor := &f.Primeiro
	for {
		if (*cursor == nil) {
			*cursor = &item
			return
		}
		cursor = &(*cursor).proximo
	}
}

func (f Fila[T]) Len() int {
	i := 0
	cursor := f.Primeiro
	for {
		if (cursor == nil) {
			return i
		}
		cursor = cursor.proximo
		i++
	}
}



