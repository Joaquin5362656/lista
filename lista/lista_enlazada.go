package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func crearNodoLista[T any](dato T, siguiente *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato, siguiente}
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iteradorLista[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) Iterador() iteradorLista[T] {
	return iteradorLista[T]{lista, lista.primero, nil}
}

func (iterador *iteradorLista[T]) VerActual() T {

	if iterador.actual == nil {
		panic("El iterador termino de iterar")
	}

	return iterador.actual.dato
}

func (iterador *iteradorLista[T]) HaySiguiente() bool {

	if iterador.actual == nil {
		return false
	} else {
		return true
	}
}

func (iterador *iteradorLista[T]) Siguiente() {

	if iterador.actual == nil {
		panic("El iterador termino de iterar")
	}

	anteriorLista := iterador.actual
	iterador.actual = iterador.actual.siguiente
	iterador.anterior = anteriorLista
}

func (iterador *iteradorLista[T]) Insertar(dato T) {

	anteriorLista := iterador.anterior
	esFinalLista := iterador.actual == nil
	esInicioLista := anteriorLista == nil

	iterador.actual = crearNodoLista(dato, iterador.actual)

	if esInicioLista {
		iterador.lista.primero = iterador.actual
	} else {
		anteriorLista.siguiente = iterador.actual
	}

	if esFinalLista {
		iterador.lista.ultimo = iterador.actual
	}

	iterador.lista.largo++
}

func (iterador *iteradorLista[T]) Borrar() T {

	if iterador.actual == nil {
		panic("El iterador termino de iterar")
	}

	elementoABorrar := iterador.actual.dato

	esFinalLista := iterador.actual == nil
	esInicioLista := iterador.anterior == nil

	iterador.actual = iterador.actual.siguiente

	if esInicioLista {
		iterador.lista.primero = iterador.actual
	} else {
		anteriorLista := iterador.anterior
		anteriorLista.siguiente = iterador.actual
	}

	if esFinalLista {
		iterador.lista.ultimo = iterador.anterior
	}

	iterador.lista.largo--

	return elementoABorrar

}
