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
// FUncion que crea una lista enlazada
func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
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
