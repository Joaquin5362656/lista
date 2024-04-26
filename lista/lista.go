package lista

type Lista[T any] interface {
	EstaVacia() bool
	InsertarPrimero(T)
	InsertarUltimo(T)
	BorrarPrimero() T
	VerPrimero() T
	VerUltimo() T
	Largo() int
	Iterar(visitar func(T) bool)

	// Iterador crea el iterador que permite recorrer los elementos de la lista, usando
	// las primitivas de este
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento en la instancia de la iteracion en que se encuentra el iterador.
	// Si se iteraron todos los elementos entra en panico con el mensaje "El iterador termino de iterar"
	VerActual() T

	// HaySiguiente devuelve true si quedan elementos en la lista para iterar y devuelve false si ya
	// se iteraron todos los elementos y no queda ningun otro por iterar.
	HaySiguiente() bool

	// Siguiente hace que avanze a otra instancia del iterador. Haciendo que pase a otro elemento
	// de la lista listo para ser iterado.
	Siguiente()

	// Insertar agrega un nuevo elemento en la lista en la posicion actual en la que se encuentra el
	// iterador, pasando a una instancia donde este nuevo elemento sea el nuevo actual y este listo
	// para ser iterado y el elemento que anteriormente estaba en la posicion actual pasa a ser
	// el siguiente elemento a iterar.
	// Al insertar se aumenta el largo de la lista y en caso de insertarse al principio o al final
	// de la lista modifica el primer elemento o el ultimo de la lista
	Insertar(T)

	// Borrar elimina el elemento en la posicion actual en la que se encuentra el iterador,
	// disminuyendo el largo de la lista y pasando a la siguiente instancia de la iteracion
	// con el siguiente elemento de la lista.
	// En caso de borrar el primer elemento de la lista, el primero pasa a ser el siguiente
	// elemento de la lista y en caso de ser el elemento al final de la lista, el ultimo
	// pasa a ser el anterior elemento a este
	Borrar() T
}
