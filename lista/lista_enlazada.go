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

// Esta funcion va a dejar de tirar error cuando se implementen los iteradores
// FUncion que crea una lista enlazada
func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{nil, nil, 0}
}

// EstaVacia devuelve un booleano en caso de que la pila este o no vacia
func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

// InsertarPrimero agrega un elemento al final de la lista
func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := &nodoLista[T]{dato: elemento}

	if lista.EstaVacia() {
		lista.primero = nuevoNodo
		lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.primero
		lista.primero = nuevoNodo
	}

	lista.largo++
}

// InsertarUltimo agrega un elemento al final de la lista
func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := &nodoLista[T]{dato: elemento}

	if lista.EstaVacia() {
		lista.primero = nuevoNodo
		lista.ultimo = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
		lista.ultimo = nuevoNodo
	}

	lista.largo++
}

// BorrarPrimero borra y retorna el primer elemento de la lista. Tira panic en caso de que la lista este vacia
func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--

	if lista.largo == 0 {
		lista.ultimo = nil
	}

	return dato
}

// VerPrimero retorna el primer elemento de la lista. Tira panic en caso de que la lista este vacia
func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.primero.dato
}

// VerUltimo retorna el ultimo elemento de la lista. Tira panic en caso de que la lista este vacia
func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.ultimo.dato
}

// Largo retorna el nro de elementos en la lista.
func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {

}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLista[T]{lista, lista.primero, nil}
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

	esFinalLista := iterador.actual.siguiente == nil
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
