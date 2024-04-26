package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iteradorLista[T any] struct {
	//
}

// Esta funcion va a dejar de tirar error cuando se implementen los iteradores
func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

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

// BorrarPrimero borra y retorna el primer elemento de la lista.
// Panic: Si la lista está vacía.
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

// VerPrimero retorna el primer elemento de la lista.
// Panic: Si la lista está vacía.
func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.primero.dato
}

// VerUltimo retorna el último elemento de la lista.
// Panic: Si la lista está vacía.
func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.ultimo.dato
}

// Largo retorna el número de elementos en la lista.
func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}
